package spanner

import (
	"context"
	"log"
	"strconv"
	"strings"

	"go.opencensus.io/tag"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"google.golang.org/grpc/metadata"
)

const otStatsPrefix = "cloud.google.com/go/spanner/"

var (
	attributeKeyClientID   = attribute.Key("client_id")
	attributeKeyDatabase   = attribute.Key("database")
	attributeKeyInstance   = attribute.Key("instance_id")
	attributeKeyLibVersion = attribute.Key("library_version")
	attributeKeyType       = attribute.Key("type")
	attributeKeyMethod     = attribute.Key("grpc_client_method")

	attributeNumInUseSessions = attributeKeyType.String("num_in_use_sessions")
	attributeNumSessions      = attributeKeyType.String("num_sessions")
)

func captureGFELatencyStatsOT(ctx context.Context, md metadata.MD, keyMethod string, otClientConfig *openTelemetryClientConfig, attr []attribute.KeyValue) error {
	if len(md.Get("server-timing")) == 0 {
		otClientConfig.gfeHeaderMissingCount.Add(ctx, 1, metric.WithAttributes(attr...))
		return nil
	}
	serverTiming := md.Get("server-timing")[0]
	gfeLatency, err := strconv.Atoi(strings.TrimPrefix(serverTiming, "gfet4t7; dur="))
	if !strings.HasPrefix(serverTiming, "gfet4t7; dur=") || err != nil {
		return err
	}
	// Record GFE latency with OpenCensus.
	ctx = tag.NewContext(ctx, tag.FromContext(ctx))
	attr = append(attr, attributeKeyMethod.String(keyMethod))
	otClientConfig.gfeLatency.Record(ctx, int64(gfeLatency), metric.WithAttributes(attr...))
	return nil
}

func createContextAndCaptureGFELatencyMetricsOT(ctx context.Context, ct *commonTags, md metadata.MD, keyMethod string, otClientConfig *openTelemetryClientConfig) error {
	attr := []attribute.KeyValue{
		attributeKeyClientID.String(ct.clientID),
		attributeKeyDatabase.String(ct.database),
		attributeKeyInstance.String(ct.instance),
		attributeKeyLibVersion.String(ct.libVersion),
	}
	return captureGFELatencyStatsOT(ctx, md, keyMethod, otClientConfig, attr)
}

func initializeOTMetricInstruments(otClientConfig *openTelemetryClientConfig) {
	meter := otClientConfig.meterProvider.Meter(otStatsPrefix)

	otClientConfig.openSessionCount, _ = meter.Int64ObservableGauge(
		"open_session_count_test_ot_local",
		metric.WithDescription("Number of sessions currently opened"),
		metric.WithUnit("1"),
	)

	otClientConfig.maxAllowedSessionsCount, _ = meter.Int64ObservableGauge(
		"max_allowed_sessions_test_ot_local",
		metric.WithDescription("The maximum number of sessions allowed. Configurable by the user."),
		metric.WithUnit("1"),
	)

	otClientConfig.sessionsCount, _ = meter.Int64ObservableGauge(
		"num_sessions_in_pool_test_ot_local",
		metric.WithDescription("The number of sessions currently in use."),
		metric.WithUnit("1"),
	)

	otClientConfig.maxInUseSessionsCount, _ = meter.Int64ObservableGauge(
		"max_in_use_sessions_test_ot_local",
		metric.WithDescription("The maximum number of sessions in use during the last 10 minute interval."),
		metric.WithUnit("1"),
	)

	otClientConfig.getSessionTimeoutsCount, _ = meter.Int64Counter(
		"get_session_timeouts_int64counter",
		metric.WithDescription("The number of get sessions timeouts due to pool exhaustion."),
		metric.WithUnit("1"),
	)

	otClientConfig.acquiredSessionsCount, _ = meter.Int64Counter(
		"num_acquired_sessions_test_ot_ctr_local",
		metric.WithDescription("The number of sessions acquired from the session pool."),
		metric.WithUnit("1"),
	)

	otClientConfig.releasedSessionsCount, _ = meter.Int64Counter(
		"num_released_sessions_test_ot_ctr_local",
		metric.WithDescription("The number of sessions released by the user and pool maintainer."),
		metric.WithUnit("1"),
	)

	otClientConfig.gfeLatency, _ = meter.Int64Histogram(
		"gfe_latency_test_ot_local",
		metric.WithDescription("Latency between Google's network receiving an RPC and reading back the first byte of the response"),
		metric.WithUnit("ms"),
		metric.WithExplicitBucketBoundaries(0.0, 0.01, 0.05, 0.1, 0.3, 0.6, 0.8, 1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 8.0, 10.0, 13.0,
			16.0, 20.0, 25.0, 30.0, 40.0, 50.0, 65.0, 80.0, 100.0, 130.0, 160.0, 200.0, 250.0,
			300.0, 400.0, 500.0, 650.0, 800.0, 1000.0, 2000.0, 5000.0, 10000.0, 20000.0, 50000.0,
			100000.0),
	)

	otClientConfig.gfeHeaderMissingCount, _ = meter.Int64Counter(
		"gfe_header_missing_count_local",
		metric.WithDescription("Number of RPC responses received without the server-timing header, most likely means that the RPC never reached Google's network"),
		metric.WithUnit("1"),
	)
}

func registerSessionPoolOTMetrics(pool *sessionPool) error {
	otConfig := pool.openTelemetryConfig
	if otConfig == nil {
		return nil
	}

	attributes := otConfig.attributeMap
	attributesInUseSessions := append(attributes, attributeNumInUseSessions)
	attributesAvailableSessions := append(attributes, attributeNumSessions)

	reg, err := otConfig.meterProvider.Meter(otStatsPrefix).RegisterCallback(
		func(ctx context.Context, o metric.Observer) error {
			pool.mu.Lock()
			defer pool.mu.Unlock()
			log.Print("RegisterCallback called for meter")

			o.ObserveInt64(otConfig.openSessionCount, int64(pool.numOpened), metric.WithAttributes(attributes...))
			o.ObserveInt64(otConfig.maxAllowedSessionsCount, int64(pool.SessionPoolConfig.MaxOpened), metric.WithAttributes(attributes...))
			o.ObserveInt64(otConfig.sessionsCount, int64(pool.numInUse), metric.WithAttributes(attributesInUseSessions...))
			o.ObserveInt64(otConfig.sessionsCount, int64(pool.numSessions), metric.WithAttributes(attributesAvailableSessions...))
			o.ObserveInt64(otConfig.maxInUseSessionsCount, int64(pool.maxNumInUse), metric.WithAttributes(attributes...))

			return nil
		},
		otConfig.openSessionCount,
		otConfig.maxAllowedSessionsCount,
		otConfig.sessionsCount,
		otConfig.maxInUseSessionsCount,
	)
	if err != nil {
		return err
	}

	pool.openTelemetryConfig.otMetricRegistration = reg
	return nil
}
