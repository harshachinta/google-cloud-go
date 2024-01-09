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
	"go.opentelemetry.io/otel/sdk/metric"
	"google.golang.org/api/option"
	"log"
	"math/rand"
	"sort"
	"sync"
	"testing"
	"time"

	"google.golang.org/api/iterator"
)

var muElapsedTimes sync.Mutex
var elapsedTimes []time.Duration
var (
	selectQuery         = "SELECT ID FROM BENCHMARK WHERE ID = @id"
	update_query        = "UPDATE BENCHMARK SET BAR=1 WHERE ID = @id"
	idColumnName        = "id"
	randomSearchSpace   = 99999
	totalReadsPerThread = 100
	parallelThreads     = 1
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

func BenchmarkClientBurstReadIncStep25RealServer(b *testing.B) {
	b.Logf("Running Benchmark")
	elapsedTimes = []time.Duration{}
	meterProvider := setupAndEnableOT()
	burstRead(b, 25, "projects/span-cloud-testing/instances/harsha-test-gcloud/databases/database1", meterProvider)
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
	b.Logf("%q", elapsedTimes)
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
