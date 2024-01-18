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
	"log"
	"time"

	"cloud.google.com/go/internal/trace"
	metricExporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/metric"
	traceExporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

/*
var muElapsedTimes sync.Mutex
var elapsedTimes []time.Duration
var (
	selectQuery           = "SELECT ID FROM BENCHMARK WHERE ID = @id"
	updateQuery           = "UPDATE BENCHMARK SET BAR=1 WHERE ID = @id"
	idColumnName          = "id"
	randomSearchSpace     = 99999
	totalReadsPerThread   = 30000
	totalUpdatesPerThread = 10000
	parallelThreads       = 5
)

func createBenchmarkActualServer(ctx context.Context, incStep uint64, clientConfig ClientConfig, database string, mp *metric.MeterProvider) (client *Client, err error) {
	t := &testing.T{}
	clientConfig.SessionPoolConfig = SessionPoolConfig{
		MinOpened: 100,
		MaxOpened: 400,
		incStep:   incStep,
	}
	if mp != nil {
		clientConfig.OpenTelemetryMeterProvider = mp
	}
	options := []option.ClientOption{option.WithEndpoint("staging-wrenchworks.sandbox.googleapis.com:443")}
	client, err = NewClientWithConfig(ctx, database, clientConfig, options...)
	if err != nil {
		log.Printf("Newclient error : %q", err)
	}
	log.Printf("New client initialized")
	// Wait until the session pool has been initialized.
	waitFor(t, func() error {
		if uint64(client.idleSessions.idleList.Len()) == client.idleSessions.MinOpened {
			return nil
		}
		return fmt.Errorf("not yet initialized")
	})
	return
}

func readWorkerReal1(client *Client, b *testing.B, jobs <-chan int, results chan<- int) {
	for range jobs {
		startTime := time.Now()
		iter := client.Single().Query(context.Background(), getRandomisedReadStatement())
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
		}
		iter.Stop()

		// Calculate the elapsed time
		elapsedTime := time.Since(startTime)
		storeElapsedTime(elapsedTime)

		// return row as 1, so that we know total number of queries executed.
		results <- row
	}
}

func writeWorkerReal1(client *Client, b *testing.B, jobs <-chan int, results chan<- int64) {
	for range jobs {
		startTime := time.Now()
		var updateCount int64
		var err error
		if _, err = client.ReadWriteTransaction(context.Background(), func(ctx context.Context, transaction *ReadWriteTransaction) error {
			if updateCount, err = transaction.Update(ctx, getRandomisedUpdateStatement()); err != nil {
				return err
			}
			return nil
		}); err != nil {
			b.Fatal(err)
		}

		// Calculate the elapsed time
		elapsedTime := time.Since(startTime)
		storeElapsedTime(elapsedTime)

		results <- updateCount
	}
}

func BenchmarkClientBurstReadIncStep25RealServerOpenTelemetry(b *testing.B) {
	b.Logf("Running Burst Read Benchmark With opentelemetry instrumentation")
	elapsedTimes = []time.Duration{}
	meterProvider := setupAndEnableOT()
	burstRead(b, 25, "projects/span-cloud-testing/instances/harsha-test-gcloud/databases/database1", meterProvider)
}

func BenchmarkClientBurstWriteIncStep25RealServerOpenTelemetry(b *testing.B) {
	b.Logf("Running Burst Write Benchmark With opentelemetry instrumentation")
	elapsedTimes = []time.Duration{}
	meterProvider := setupAndEnableOT()
	burstWrite(b, 25, "projects/span-cloud-testing/instances/harsha-test-gcloud/databases/database1", meterProvider)
}

func BenchmarkClientBurstReadWriteIncStep25RealServerOpenTelemetry(b *testing.B) {
	b.Logf("Running Burst Read and Write Benchmark With opentelemetry instrumentation")
	elapsedTimes = []time.Duration{}
	meterProvider := setupAndEnableOT()
	burstReadAndWrite(b, 25, "projects/span-cloud-testing/instances/harsha-test-gcloud/databases/database1", meterProvider)
}

func burstRead(b *testing.B, incStep uint64, database string, mp *metric.MeterProvider) {
	for n := 0; n < b.N; n++ {
		log.Printf("burstRead called once")
		client, err := createBenchmarkActualServer(context.Background(), incStep, ClientConfig{}, database, mp)
		if err != nil {
			b.Fatalf("Failed to initialize the client: error : %q", err)
		}
		sp := client.idleSessions
		log.Printf("Session pool length, %d", sp.idleList.Len())
		if uint64(sp.idleList.Len()) != sp.MinOpened {
			b.Fatalf("session count mismatch\nGot: %d\nWant: %d", sp.idleList.Len(), sp.MinOpened)
		}

		totalQueries := parallelThreads * totalReadsPerThread
		jobs := make(chan int, totalQueries)
		results := make(chan int, totalQueries)
		parallel := parallelThreads

		for w := 0; w < parallel; w++ {
			go readWorkerReal1(client, b, jobs, results)
		}
		for j := 0; j < totalQueries; j++ {
			jobs <- j
		}
		close(jobs)
		totalRows := 0
		for a := 0; a < totalQueries; a++ {
			totalRows = totalRows + <-results
		}
		b.Logf("Total Rows: %d", totalRows)
		reportBenchmarkResults(b, sp)
		client.Close()
	}
}

func burstWrite(b *testing.B, incStep uint64, database string, mp *metric.MeterProvider) {
	for n := 0; n < b.N; n++ {
		log.Printf("burstWrite called once")
		client, err := createBenchmarkActualServer(context.Background(), incStep, ClientConfig{}, database, mp)
		if err != nil {
			b.Fatalf("Failed to initialize the client: error : %q", err)
		}
		sp := client.idleSessions
		log.Printf("Session pool length, %d", sp.idleList.Len())
		if uint64(sp.idleList.Len()) != sp.MinOpened {
			b.Fatalf("session count mismatch\nGot: %d\nWant: %d", sp.idleList.Len(), sp.MinOpened)
		}

		totalUpdates := parallelThreads * totalUpdatesPerThread
		jobs := make(chan int, totalUpdates)
		results := make(chan int64, totalUpdates)
		parallel := parallelThreads

		for w := 0; w < parallel; w++ {
			go writeWorkerReal1(client, b, jobs, results)
		}
		for j := 0; j < totalUpdates; j++ {
			jobs <- j
		}
		close(jobs)
		totalRows := int64(0)
		for a := 0; a < totalUpdates; a++ {
			totalRows = totalRows + <-results
		}
		b.Logf("Total Updates: %d", totalRows)
		reportBenchmarkResults(b, sp)
		client.Close()
	}
}

func burstReadAndWrite(b *testing.B, incStep uint64, database string, mp *metric.MeterProvider) {
	for n := 0; n < b.N; n++ {
		log.Printf("burstReadAndWrite called once")
		client, err := createBenchmarkActualServer(context.Background(), incStep, ClientConfig{}, database, mp)
		if err != nil {
			b.Fatalf("Failed to initialize the client: error : %q", err)
		}
		sp := client.idleSessions
		if uint64(sp.idleList.Len()) != sp.MinOpened {
			b.Fatalf("session count mismatch\nGot: %d\nWant: %d", sp.idleList.Len(), sp.MinOpened)
		}

		totalUpdates := parallelThreads * totalUpdatesPerThread
		writeJobs := make(chan int, totalUpdates)
		writeResults := make(chan int64, totalUpdates)
		parallelWrites := parallelThreads

		totalQueries := parallelThreads * totalReadsPerThread
		readJobs := make(chan int, totalQueries)
		readResults := make(chan int, totalQueries)
		parallelReads := parallelThreads

		for w := 0; w < parallelWrites; w++ {
			go writeWorkerReal1(client, b, writeJobs, writeResults)
		}
		for j := 0; j < totalUpdates; j++ {
			writeJobs <- j
		}
		for w := 0; w < parallelReads; w++ {
			go readWorkerReal1(client, b, readJobs, readResults)
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
		b.Logf("Total Updates: %d", totalUpdatedRows)
		totalReadRows := 0
		for a := 0; a < totalQueries; a++ {
			totalReadRows = totalReadRows + <-readResults
		}
		b.Logf("Total Reads: %d", totalReadRows)
		reportBenchmarkResults(b, sp)
		client.Close()
	}
}


func reportBenchmarkResults(b *testing.B, sp *sessionPool) {
	sp.mu.Lock()
	defer sp.mu.Unlock()
	b.Logf("NumSessions: %d\t", sp.idleList.Len())

	muElapsedTimes.Lock()
	defer muElapsedTimes.Unlock()
	sort.Slice(elapsedTimes, func(i, j int) bool {
		return elapsedTimes[i] < elapsedTimes[j]
	})

	b.Logf("Total number of queries: %d\n", len(elapsedTimes))
	//	b.Logf("%q", elapsedTimes)
	b.Logf("P50: %q\n", percentile(50, elapsedTimes))
	b.Logf("P95: %q\n", percentile(95, elapsedTimes))
	b.Logf("P99: %q\n", percentile(99, elapsedTimes))
	elapsedTimes = nil
}

func percentile(percentile int, orderedResults []time.Duration) time.Duration {
	index := percentile * len(orderedResults) / 100
	value := orderedResults[index]
	return value
}

func storeElapsedTime(elapsedTime time.Duration) {
	muElapsedTimes.Lock()
	defer muElapsedTimes.Unlock()
	elapsedTimes = append(elapsedTimes, elapsedTime)
}

func getRandomisedReadStatement() Statement {
	randomKey := rand.Intn(randomSearchSpace)
	stmt := NewStatement(selectQuery)
	stmt.Params["id"] = randomKey
	return stmt
}

func getRandomisedUpdateStatement() Statement {
	randomKey := rand.Intn(randomSearchSpace)
	stmt := NewStatement(updateQuery)
	stmt.Params["id"] = randomKey
	return stmt
}*/

func setupAndEnableOT() *metric.MeterProvider {
	res, err := newResource()
	if err != nil {
		log.Fatal(err)
	}

	EnableOpenTelemetryMetrics()
	// Create a meter provider.
	// You can pass this instance directly to your instrumented code if it
	// accepts a MeterProvider instance.
	meterProvider, err := newMeterProvider(res)
	if err != nil {
		panic(err)
	}
	trace.OpenTelemetryTracingEnabled = true
	setTracerProvider(res)
	return meterProvider
}

func newMeterProvider(res *resource.Resource) (*metric.MeterProvider, error) {
	exporter, err := metricExporter.New(
		metricExporter.WithProjectID("span-cloud-testing"),
		//metricExporter.WithMetricDescriptorTypeFormatter(
		//	func(metrics metricdata.Metrics) string {
		//		return fmt.Sprintf("custom.googleapis.com/%s", metrics.Name)
		//	},
		//)
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

func setTracerProvider(res *resource.Resource) {
	exporter, err := traceExporter.New(traceExporter.WithProjectID("span-cloud-testing"))
	if err != nil {
		log.Print(err)
	}

	traceProvider := sdktrace.NewTracerProvider(
		sdktrace.WithResource(res),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
	)
	otel.SetTracerProvider(traceProvider)
}
