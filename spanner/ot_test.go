package spanner

import (
	"context"
	"testing"

	"cloud.google.com/go/internal/testutil"
)

func TestOTStats(t *testing.T) {
	t.Logf("restarting the test with different parameter")
	ctx := context.Background()
	te := testutil.NewOpenTelemetryTestExporter()
	t.Cleanup(func() {
		te.Unregister(ctx)
	})

	_, c, teardown := setupMockedTestServer(t)
	defer teardown()

	c.Single().ReadRow(context.Background(), "Users", Key{"alice"}, []string{"email"})
	rm, err := te.Metrics(ctx)
	if err != nil {
		t.Error(err)
	}
	t.Logf("Length of scope metrics %d", len(rm.ScopeMetrics))
	t.Logf("Length of scope metrics %q", rm.ScopeMetrics)
	t.Logf("Length of scope metrics %q", rm.ScopeMetrics[0].Metrics)
	t.Logf("Length of scope metrics %q", rm.ScopeMetrics[0].Scope)
	t.Logf("Length of scope metrics %q", rm.ScopeMetrics[1].Metrics)
	t.Logf("Length of scope metrics %d", len(rm.ScopeMetrics[1].Metrics))
	t.Logf("Length of scope metrics %q", rm.ScopeMetrics[1].Scope)
}
