package executor

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	executorpb "cloud.google.com/go/spanner/executor/proto"
	"google.golang.org/grpc"
)

var (
	port = flag.String("port", "", "server port")
)

func main() {
	// If we're running in a test, write logs to the outputs dir
	// so they will be collected and associated with this test.
	if d := os.Getenv("TEST_UNDECLARED_OUTPUTS_DIR"); d != "" {
		os.Args = append(os.Args, "--log_dir="+d)
	}

	flag.Parse()
	// Print "port:<number>" to STDOUT for the systest worker.
	if *port == "" {
		log.Fatalf("usage: %s --port=8081", os.Args[0])
		// portpicker not available, should we instead return a fatal
		// log.Fatalf("usage: %s --port=8081", os.Args[0])
	}
	fmt.Printf("port:%s\n", *port)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", *port))
	if err != nil {
		log.Fatal(err)
	}

	i, err := NewCloudProxyServer(context.Background())
	if err != nil {
		log.Fatalf("NewCloudProxyServer failed: %v", err)
	}

	s := grpc.NewServer()
	executorpb.RegisterSpannerExecutorProxyServer(s, i)

	log.Fatal(s.Serve(lis))
}
