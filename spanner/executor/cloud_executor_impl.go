package executor

import (
	"context"

	executorpb "cloud.google.com/go/spanner/executor/proto"
)

// CloudProxyServer holds the cloud go server.
type CloudProxyServer struct{}

// NewCloudProxyServer initializes and returns a new InfraProxyServer instance.
func NewCloudProxyServer(ctx context.Context) (*CloudProxyServer, error) {
	return &CloudProxyServer{}, nil
}

// ExecuteActionAsync is implementation of ExecuteActionAsync in AsyncSpannerExecutorProxy. It's a
// streaming method in which client and server exchange SpannerActions and SpannerActionOutcomes.
func (s *CloudProxyServer) ExecuteActionAsync(stream executorpb.SpannerExecutorProxy_ExecuteActionAsyncServer) error {
	h := &cloudStreamHandler{
		cloudProxyServer: s,
		stream:           stream,
	}
	return h.execute()
}
