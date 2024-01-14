/*
Copyright 2024 Google LLC

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

package test

import (
	"context"
	"errors"
	"testing"
	"time"

	"cloud.google.com/go/spanner"
	"cloud.google.com/go/spanner/apiv1/spannerpb"
	"cloud.google.com/go/spanner/internal"
	stestutil "cloud.google.com/go/spanner/internal/testutil"
	structpb "github.com/golang/protobuf/ptypes/struct"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	"go.opentelemetry.io/otel/sdk/metric/metricdata/metricdatatest"
	"google.golang.org/api/iterator"
)

func TestOTMetrics_InstrumentationScope(t *testing.T) {
	ctx := context.Background()
	te := newOpenTelemetryTestExporter(false, false)
	t.Cleanup(func() {
		te.Unregister(ctx)
	})
	spanner.EnableOpenTelemetryMetrics()
	_, c, teardown := setupMockedTestServerWithConfig(t, spanner.ClientConfig{OpenTelemetryMeterProvider: te.mp})
	defer teardown()

	c.Single().ReadRow(context.Background(), "Users", spanner.Key{"alice"}, []string{"email"})
	rm, err := te.metrics(ctx)
	if err != nil {
		t.Error(err)
	}
	if len(rm.ScopeMetrics) != 1 {
		t.Fatalf("Error in number of instrumentation scope, got: %d, want: %d", len(rm.ScopeMetrics), 1)
	}
	if rm.ScopeMetrics[0].Scope.Name != spanner.OtInstrumentationScope {
		t.Fatalf("Error in instrumentation scope name, got: %s, want: %s", rm.ScopeMetrics[0].Scope.Name, spanner.OtInstrumentationScope)
	}
	if rm.ScopeMetrics[0].Scope.Version != internal.Version {
		t.Fatalf("Error in instrumentation scope version, got: %s, want: %s", rm.ScopeMetrics[0].Scope.Version, internal.Version)
	}
	/*t.Logf("Length of scope metrics %d", len(rm.ScopeMetrics))
	t.Logf("Length of scope metrics %q", rm.ScopeMetrics)
	t.Logf("Length of scope metrics %q", rm.ScopeMetrics[0].Metrics)
	t.Logf("Length of scope metrics %q", rm.ScopeMetrics[0].Scope)
	t.Logf("Length of scope metrics %q", rm.ScopeMetrics[1].Metrics)
	t.Logf("Length of scope metrics %d", len(rm.ScopeMetrics[1].Metrics))
	t.Logf("Length of scope metrics %q", rm.ScopeMetrics[1].Scope)*/

	/*for _, m := range rm.ScopeMetrics {
		t.Log(m.Scope)
		for _, metric := range m.Metrics {
			t.Log(metric.Data)
			switch metric.Data.(type) {
			case metricdata.Gauge[int64]:
				a := metric.Data.(metricdata.Gauge[int64]).DataPoints
				t.Log(a)
			case metricdata.Gauge[float64]:
				a := metric.Data.(metricdata.Gauge[float64]).DataPoints
				t.Log(a)
			case metricdata.Sum[int64]:
				a := metric.Data.(metricdata.Sum[int64]).DataPoints
				t.Log(a)
			case metricdata.Sum[float64]:
				a := metric.Data.(metricdata.Sum[float64]).DataPoints
				t.Log(a)
			}
		}
	}*/
	//metricdatatest.AssertAggregationsEqual()
}

func TestOTMetrics_SessionPool(t *testing.T) {
	ctx := context.Background()
	te := newOpenTelemetryTestExporter(false, false)
	t.Cleanup(func() {
		te.Unregister(ctx)
	})
	spanner.EnableOpenTelemetryMetrics()

	_, client, teardown := setupMockedTestServerWithConfig(t, spanner.ClientConfig{OpenTelemetryMeterProvider: te.mp})
	defer teardown()
	client.Single().ReadRow(context.Background(), "Users", spanner.Key{"alice"}, []string{"email"})

	for _, test := range []struct {
		name           string
		expectedMetric metricdata.Metrics
	}{
		{
			"OpenSessionCount",
			metricdata.Metrics{
				Name:        "open_session_count_test_ot_local",
				Description: "Number of sessions currently opened",
				Unit:        "1",
				Data: metricdata.Gauge[int64]{
					DataPoints: []metricdata.DataPoint[int64]{
						{
							Attributes: attribute.NewSet(getAttributes(client.ClientID())...),
							Value:      25,
						},
					},
				},
			},
		},
		{
			"MaxAllowedSessionsCount",
			metricdata.Metrics{
				Name:        "max_allowed_sessions_test_ot_local",
				Description: "The maximum number of sessions allowed. Configurable by the user.",
				Unit:        "1",
				Data: metricdata.Gauge[int64]{
					DataPoints: []metricdata.DataPoint[int64]{
						{
							Attributes: attribute.NewSet(getAttributes(client.ClientID())...),
							Value:      400,
						},
					},
				},
			},
		},
		{
			"MaxInUseSessionsCount",
			metricdata.Metrics{
				Name:        "max_in_use_sessions_test_ot_local",
				Description: "The maximum number of sessions in use during the last 10 minute interval.",
				Unit:        "1",
				Data: metricdata.Gauge[int64]{
					DataPoints: []metricdata.DataPoint[int64]{
						{
							Attributes: attribute.NewSet(getAttributes(client.ClientID())...),
							Value:      1,
						},
					},
				},
			},
		},
		{
			"AcquiredSessionsCount",
			metricdata.Metrics{
				Name:        "num_acquired_sessions_test_ot_ctr_local",
				Description: "The number of sessions acquired from the session pool.",
				Unit:        "1",
				Data: metricdata.Sum[int64]{
					DataPoints: []metricdata.DataPoint[int64]{
						{
							Attributes: attribute.NewSet(getAttributes(client.ClientID())...),
							Value:      1,
						},
					},
					Temporality: metricdata.CumulativeTemporality,
					IsMonotonic: true,
				},
			},
		},
		{
			"ReleasedSessionsCount",
			metricdata.Metrics{
				Name:        "num_released_sessions_test_ot_ctr_local",
				Description: "The number of sessions released by the user and pool maintainer.",
				Unit:        "1",
				Data: metricdata.Sum[int64]{
					DataPoints: []metricdata.DataPoint[int64]{
						{
							Attributes: attribute.NewSet(getAttributes(client.ClientID())...),
							Value:      1,
						},
					},
					Temporality: metricdata.CumulativeTemporality,
					IsMonotonic: true,
				},
			},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			metricName := test.expectedMetric.Name
			expectedMetric := test.expectedMetric
			validateOTMetric(t, ctx, te, metricName, expectedMetric)
		})
	}
}

func TestOTStats_SessionPool_SessionsCount(t *testing.T) {
	ctx := context.Background()
	te := newOpenTelemetryTestExporter(false, false)
	t.Cleanup(func() {
		te.Unregister(ctx)
	})

	spanner.EnableOpenTelemetryMetrics()
	server, client, teardown := setupMockedTestServerWithConfig(t, spanner.ClientConfig{SessionPoolConfig: spanner.DefaultSessionPoolConfig, OpenTelemetryMeterProvider: te.mp})
	client.DatabaseName()
	defer teardown()
	// Wait for the session pool initialization to finish.
	expectedReads := spanner.DefaultSessionPoolConfig.MinOpened
	waitFor(t, func() error {
		if uint64(server.TestSpanner.TotalSessionsCreated()) == expectedReads {
			return nil
		}
		return errors.New("Not yet initialized")
	})

	client.Single().ReadRow(context.Background(), "Users", spanner.Key{"alice"}, []string{"email"})

	attributesNumInUseSessions := append(getAttributes(client.ClientID()), attribute.Key("type").String("num_in_use_sessions"))
	attributesNumSessions := append(getAttributes(client.ClientID()), attribute.Key("type").String("num_sessions"))

	expectedMetricData := metricdata.Metrics{
		Name:        "num_sessions_in_pool_test_ot_local",
		Description: "The number of sessions currently in use.",
		Unit:        "1",
		Data: metricdata.Gauge[int64]{
			DataPoints: []metricdata.DataPoint[int64]{
				{
					Attributes: attribute.NewSet(attributesNumInUseSessions...),
					Value:      0,
				},
				{
					Attributes: attribute.NewSet(attributesNumSessions...),
					Value:      100,
				},
			},
		},
	}

	validateOTMetric(t, ctx, te, expectedMetricData.Name, expectedMetricData)
}

func TestOTStats_SessionPool_GetSessionTimeoutsCount(t *testing.T) {
	ctx1 := context.Background()
	te := newOpenTelemetryTestExporter(false, false)
	t.Cleanup(func() {
		te.Unregister(ctx1)
	})
	spanner.EnableOpenTelemetryMetrics()
	server, client, teardown := setupMockedTestServerWithConfig(t, spanner.ClientConfig{OpenTelemetryMeterProvider: te.mp})
	defer teardown()

	server.TestSpanner.PutExecutionTime(stestutil.MethodBatchCreateSession,
		stestutil.SimulatedExecutionTime{
			MinimumExecutionTime: 2 * time.Millisecond,
		})

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()
	client.Single().ReadRow(ctx, "Users", spanner.Key{"alice"}, []string{"email"})

	expectedMetricData := metricdata.Metrics{
		Name:        "get_session_timeouts_int64counter",
		Description: "The number of get sessions timeouts due to pool exhaustion.",
		Unit:        "1",
		Data: metricdata.Sum[int64]{
			DataPoints: []metricdata.DataPoint[int64]{
				{
					Attributes: attribute.NewSet(getAttributes(client.ClientID())...),
					Value:      1,
				},
			},
			Temporality: metricdata.CumulativeTemporality,
			IsMonotonic: true,
		},
	}
	validateOTMetric(t, ctx1, te, expectedMetricData.Name, expectedMetricData)
}

func TestOTStats_GFE_Latency(t *testing.T) {
	ctx := context.Background()
	te := newOpenTelemetryTestExporter(false, false)
	t.Cleanup(func() {
		te.Unregister(ctx)
	})
	spanner.EnableOpenTelemetryMetrics()
	server, client, teardown := setupMockedTestServerWithConfig(t, spanner.ClientConfig{OpenTelemetryMeterProvider: te.mp})
	defer teardown()

	if err := server.TestSpanner.PutStatementResult("SELECT email FROM Users", &stestutil.StatementResult{
		Type: stestutil.StatementResultResultSet,
		ResultSet: &spannerpb.ResultSet{
			Metadata: &spannerpb.ResultSetMetadata{
				RowType: &spannerpb.StructType{
					Fields: []*spannerpb.StructType_Field{
						{
							Name: "email",
							Type: &spannerpb.Type{Code: spannerpb.TypeCode_STRING},
						},
					},
				},
			},
			Rows: []*structpb.ListValue{
				{Values: []*structpb.Value{{
					Kind: &structpb.Value_StringValue{StringValue: "test@test.com"},
				}}},
			},
		},
	}); err != nil {
		t.Fatalf("could not add result: %v", err)
	}
	iter := client.Single().Read(context.Background(), "Users", spanner.AllKeys(), []string{"email"})
	for {
		_, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			t.Fatal(err.Error())
		}
	}

	attributeGFELatency := append(getAttributes(client.ClientID()), attribute.Key("grpc_client_method").String("executeBatchCreateSessions"))

	resourceMetrics, err := te.metrics(context.Background())
	if err != nil {
		t.Error(err)
	}
	if resourceMetrics == nil {
		t.Fatal("Resource Metrics is nil")
	}
	if got, want := len(resourceMetrics.ScopeMetrics), 1; got != want {
		t.Fatalf("ScopeMetrics length mismatch, got %v, want %v", got, want)
	}

	idx := getMetricIndex(resourceMetrics.ScopeMetrics[0].Metrics, "gfe_latency_test_ot_local")
	if idx == -1 {
		t.Fatalf("Metric Name %s not found", "gfe_latency_test_ot_local")
	}
	/*
		expectedMetricData := metricdata.Metrics{
			Name:        "gfe_latency_test_ot_local",
			Description: "Latency between Google's network receiving an RPC and reading back the first byte of the response",
			Unit:        "ms",
			Data: metricdata.Histogram[int64]{
				DataPoints: []metricdata.HistogramDataPoint[int64]{
					{
						Attributes: attr_gfe_latency,
						Bounds: []float64{0.0, 0.01, 0.05, 0.1, 0.3, 0.6, 0.8, 1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 8.0, 10.0, 13.0,
							16.0, 20.0, 25.0, 30.0, 40.0, 50.0, 65.0, 80.0, 100.0, 130.0, 160.0, 200.0, 250.0,
							300.0, 400.0, 500.0, 650.0, 800.0, 1000.0, 2000.0, 5000.0, 10000.0, 20000.0, 50000.0,
							100000.0},
					},
				},
				Temporality: metricdata.CumulativeTemporality,
			},
		}
		metricdatatest.AssertEqual(t, expectedMetricData, resourceMetrics.ScopeMetrics[0].Metrics[idx], metricdatatest.IgnoreTimestamp(), metricdatatest.IgnoreExemplars())
	*/
	gfeLatencyRecordedMetric := resourceMetrics.ScopeMetrics[0].Metrics[idx]
	if gfeLatencyRecordedMetric.Name != "gfe_latency_test_ot_local" {
		t.Fatalf("Got metric name: %s, want: %s", gfeLatencyRecordedMetric.Name, "gfe_latency_test_ot_local")
	}
	if _, ok := gfeLatencyRecordedMetric.Data.(metricdata.Histogram[int64]); !ok {
		t.Fatal("gfe latency metric data not of type metricdata.Histogram[int64]")
	}
	gfeLatencyRecordedMetricData := gfeLatencyRecordedMetric.Data.(metricdata.Histogram[int64])
	count := gfeLatencyRecordedMetricData.DataPoints[0].Count
	t.Logf("Gfe latenct count %d", count)
	if got, want := count, uint64(0); got <= want {
		t.Fatalf("Incorrect data: got %d, wanted more than %d for metric %v", got, want, gfeLatencyRecordedMetric.Name)
	}
	metricdatatest.AssertHasAttributes[metricdata.HistogramDataPoint[int64]](t, gfeLatencyRecordedMetricData.DataPoints[0], attributeGFELatency...)

	idx1 := getMetricIndex(resourceMetrics.ScopeMetrics[0].Metrics, "gfe_header_missing_count_local")
	if idx1 == -1 {
		t.Fatalf("Metric Name %s not found", "gfe_header_missing_count_local")
	}

	expectedMetricData := metricdata.Metrics{
		Name:        "gfe_header_missing_count_local",
		Description: "Number of RPC responses received without the server-timing header, most likely means that the RPC never reached Google's network",
		Unit:        "1",
		Data: metricdata.Sum[int64]{
			DataPoints: []metricdata.DataPoint[int64]{
				{
					Attributes: attribute.NewSet(getAttributes(client.ClientID())...),
					Value:      1,
				},
			},
			Temporality: metricdata.CumulativeTemporality,
			IsMonotonic: true,
		},
	}
	metricdatatest.AssertEqual(t, expectedMetricData, resourceMetrics.ScopeMetrics[0].Metrics[idx1], metricdatatest.IgnoreTimestamp(), metricdatatest.IgnoreExemplars())
}

func getMetricIndex(metrics []metricdata.Metrics, metricName string) int64 {
	for i, metric := range metrics {
		if metric.Name == metricName {
			return int64(i)
		}
	}
	return -1
}

func getAttributes(clientID string) []attribute.KeyValue {
	return []attribute.KeyValue{
		attribute.Key("client_id").String(clientID),
		attribute.Key("database").String("[DATABASE]"),
		attribute.Key("instance_id").String("[INSTANCE]"),
		attribute.Key("library_version").String(internal.Version),
	}
}

func validateOTMetric(t *testing.T, ctx context.Context, te *openTelemetryTestExporter, metricName string, expectedMetric metricdata.Metrics) {
	resourceMetrics, err := te.metrics(ctx)
	if err != nil {
		t.Error(err)
	}
	if resourceMetrics == nil {
		t.Fatal("Resource Metrics is nil")
	}
	if got, want := len(resourceMetrics.ScopeMetrics), 1; got != want {
		t.Fatalf("ScopeMetrics length mismatch, got %v, want %v", got, want)
	}

	idx := getMetricIndex(resourceMetrics.ScopeMetrics[0].Metrics, metricName)
	if idx == -1 {
		t.Fatalf("Metric Name %s not found", metricName)
	}
	metricdatatest.AssertEqual(t, expectedMetric, resourceMetrics.ScopeMetrics[0].Metrics[idx], metricdatatest.IgnoreTimestamp(), metricdatatest.IgnoreExemplars())
}
