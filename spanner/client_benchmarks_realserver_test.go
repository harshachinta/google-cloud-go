/*
Copyright 2020 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package spanner

import (
	"context"
	"fmt"
	"go.opencensus.io/trace"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"google.golang.org/api/option"
	"log"
	"testing"
	"time"

	. "cloud.google.com/go/spanner/internal/testutil"
	"contrib.go.opencensus.io/exporter/stackdriver"
	metricExporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/metric"
	"github.com/loov/hrtime"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"google.golang.org/api/iterator"
)

func createBenchmarkRealServer(incStep uint64, mp *metric.MeterProvider) (client *Client) {
	t := &testing.T{}
	endpoint := "staging-wrenchworks.sandbox.googleapis.com:443"
	options := []option.ClientOption{option.WithEndpoint(endpoint)}
	clientConfig := ClientConfig{
		SessionPoolConfig: SessionPoolConfig{
			MinOpened: 100,
			MaxOpened: 400,
			incStep:   incStep,
		},
	}
	if mp != nil {
		clientConfig.OpenTelemetryMeterProvider = mp
	}

	client, err := NewClientWithConfig(context.Background(), "projects/span-cloud-testing/instances/harsha-test-gcloud/databases/database1",
		clientConfig, options...)

	fmt.Printf("Client Id: %s", client.sc.id)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Waiting for session pool")
	// Wait until the session pool has been initialized.
	/*waitFor(t, func() error {
		if uint64(client.idleSessions.idleList.Len()) == client.idleSessions.MinOpened {
			fmt.Printf("Idle Sessions length : %d ", client.idleSessions.idleList.Len())
			return nil
		}
		return fmt.Errorf("not yet initialized")
	})*/
	return
}

func readWorkerReal(client *Client, b *testing.B, jobs <-chan int, results chan<- int) {
	for range jobs {
		mu.Lock()
		d := time.Millisecond * time.Duration(rnd.Int63n(rndWaitTimeBetweenRequests))
		mu.Unlock()
		time.Sleep(d)
		iter := client.Single().Query(context.Background(), NewStatement(SelectSingerIDAlbumIDAlbumTitleFromAlbums))
		row := 0
		for {
			_, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				b.Fatal(err)
			}
			row++
			if row == 1 {
				mu.Lock()
				d := time.Millisecond * time.Duration(rnd.Int63n(holdSessionTime))
				mu.Unlock()
				time.Sleep(d)
			}
		}
		iter.Stop()
		results <- row
	}
}

func writeWorkerReal(goroutine int, client *Client, b *testing.B, jobs <-chan int, results chan<- int64) {
	for range jobs {
		mu.Lock()
		d := time.Millisecond * time.Duration(rnd.Int63n(rndWaitTimeBetweenRequests))
		mu.Unlock()
		time.Sleep(d)
		var updateCount int64
		var err error
		if _, err = client.ReadWriteTransaction(context.Background(), func(ctx context.Context, transaction *ReadWriteTransaction) error {
			if updateCount, err = transaction.Update(ctx, NewStatement("UPDATE Foo SET Age=11 WHERE name='test'")); err != nil {
				return err
			}
			//b.Logf("Inside RW Transaction routine %d, value %d, Update successful: %d", goroutine, i, updateCount)
			return nil
		}); err != nil {
			b.Fatal(err)
		}
		//b.Logf("sending Update channel successful: %d", updateCount)
		results <- updateCount
	}
}

func Benchmark_Client_BurstReadAndWrite25Real(b *testing.B) {
	sd := setupAndEnableOC()
	defer sd.Flush()
	defer sd.StopMetricsExporter()
	benchmarkClientBurstReadAndWriteRealServer(b, 25, nil)
	log.Println("Main Completed")
}

func Benchmark_Client_BurstReadAndWrite25RealOpenT(b *testing.B) {
	meterProvider := setupAndEnableOT()
	benchmarkClientBurstReadAndWriteRealServer(b, 25, meterProvider)
	log.Println("Main Completed")
}

func Benchmark_Client_OC(b *testing.B) {
	bench := hrtime.NewBenchmark(1)
	for bench.Next() {
		sd := setupAndEnableOC()
		defer sd.Flush()
		defer sd.StopMetricsExporter()
		benchmarkClientBurstReadAndWriteRealServer(b, 25, nil)
		log.Println("Main Completed")
	}
	fmt.Println(bench.Histogram(10))
}

func benchmarkClientBurstReadAndWriteRealServer(b *testing.B, incStep uint64, mp *metric.MeterProvider) {
	for n := 0; n < b.N; n++ {
		b.Logf("Running benchmark code for %d th time", n)
		client := createBenchmarkRealServer(incStep, mp)
		sp := client.idleSessions

		totalUpdates := int(sp.MaxOpened * 4)
		writeJobs := make(chan int, totalUpdates)
		writeResults := make(chan int64, totalUpdates)
		parallelWrites := int(sp.MaxOpened)

		totalQueries := int(sp.MaxOpened * 4)
		readJobs := make(chan int, totalQueries)
		readResults := make(chan int, totalQueries)
		parallelReads := int(sp.MaxOpened)

		for w := 0; w < parallelWrites; w++ {
			go writeWorkerReal(w, client, b, writeJobs, writeResults)
		}
		for j := 0; j < totalUpdates; j++ {
			writeJobs <- j
		}
		for w := 0; w < parallelReads; w++ {
			go readWorkerReal(client, b, readJobs, readResults)
		}

		for j := 0; j < totalQueries; j++ {
			readJobs <- j
		}

		close(writeJobs)
		close(readJobs)

		totalUpdatedRows := int64(0)
		for a := 0; a < totalUpdates; a++ {
			totalUpdatedRows = totalUpdatedRows + <-writeResults
		}
		totalReadRows := 0
		for a := 0; a < totalQueries; a++ {
			totalReadRows = totalReadRows + <-readResults
		}
	}
}

func setupAndEnableOC() *stackdriver.Exporter {
	if err := EnableStatViews(); err != nil {
		log.Fatalf("Failed: %v", err)
	}
	if err := EnableGfeLatencyView(); err != nil {
		log.Fatalf("Failed: %v", err)
	}
	// Create OpenCensus Stackdriver exporter.
	sd, err := stackdriver.NewExporter(stackdriver.Options{
		ProjectID:         "span-cloud-testing",
		ReportingInterval: 10 * time.Second,
	})
	if err != nil {
		log.Fatalf("Failed: %v", err)
	}
	// Start the metrics exporter
	sd.StartMetricsExporter()
	defer sd.StopMetricsExporter()

	// Register it as a trace exporter
	trace.RegisterExporter(sd)
	//trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
	return sd
}

func setupAndEnableOT() *metric.MeterProvider {
	res, err := newResource()
	if err != nil {
		log.Fatal(err)
	}

	// Create a meter provider.
	// You can pass this instance directly to your instrumented code if it
	// accepts a MeterProvider instance.
	meterProvider, err := newMeterProvider(res)
	if err != nil {
		panic(err)
	}
	return meterProvider
}

func newMeterProvider(res *resource.Resource) (*metric.MeterProvider, error) {
	exporter, err := metricExporter.New(
		metricExporter.WithProjectID("span-cloud-testing"),
		/*metricExporter.WithMetricDescriptorTypeFormatter(
			func(metrics metricdata.Metrics) string {
				return fmt.Sprintf("custom.googleapis.com/%s", metrics.Name)
			},
		)*/
	)
	if err != nil {
		return nil, err
	}

	meterProvider := metric.NewMeterProvider(
		metric.WithResource(res),
		metric.WithReader(metric.NewPeriodicReader(exporter, metric.WithInterval(10*time.Second))), // Default is 1m. Set to 3s for demonstrative purposes.
		//metric.WithReader(reader),
		//metric.WithInterval(10*time.Second),

		//metric.WithView(),
	)

	return meterProvider, nil
}

func newResource() (*resource.Resource, error) {
	return resource.Merge(resource.Default(),
		resource.NewWithAttributes(semconv.SchemaURL,
			semconv.ServiceName("my-service-spanner"),
			semconv.ServiceVersion("0.1.0"),
		))
}
