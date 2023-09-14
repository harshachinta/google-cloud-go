package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"

	"cloud.google.com/go/spanner/executor/executor"
	executorpb "cloud.google.com/go/spanner/executor/proto"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

var (
	// port = flag.String("port", "", "server port")
	port                       = flag.String("proxy_port", "", "Proxy port to start worker proxy on.")
	spanner_port               = flag.String("spanner_port", "", "Port of Spanner Frontend to which to send requests.")
	cert                       = flag.String("cert", "", "Certificate used to connect to Spanner GFE.")
	service_key_file           = flag.String("service_key_file", "", "Service key file used to set authentication.")
	use_plain_text_channel     = flag.String("use_plain_text_channel", "", "Use a plain text gRPC channel (intended for the Cloud Spanner Emulator).")
	enable_grpc_fault_injector = flag.String("enable_grpc_fault_injector", "", "Enable grpc fault injector in cloud client executor")
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
	fmt.Printf("Server started on proxyPort:%s\n", *port)
	log.Printf("Server started on proxyPort:%s\n", *port)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", *port))
	if err != nil {
		log.Fatal(err)
	}

	clientOptions := getClientOptions()
	i, err := executor.NewCloudProxyServer(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("NewCloudProxyServer failed: %v", err)
	}

	s := grpc.NewServer()
	executorpb.RegisterSpannerExecutorProxyServer(s, i)

	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(s, healthServer)

	log.Fatal(s.Serve(lis))
}

func getClientOptions() []option.ClientOption {
	var options []option.ClientOption

	endpoint := "https://localhost:" + *spanner_port
	endpoint2 := "localhost:" + *spanner_port
	log.Printf("endpoint for grpc dial :  %s", endpoint2)

	options = append(options, option.WithEndpoint(endpoint))
	options = append(options, option.WithCredentialsFile(*service_key_file))

	// Create TLS credentials from the certificate and key files.
	//creds, err := credentials.NewClientTLSFromFile(*cert, "test-cert-2")
	//if err != nil {
	//	log.Fatalf("Failed to load TLS credentials: %v", err)
	//}
	/*creds, err := loadTLSCredentials()
	if err != nil {
		log.Fatalf("Failed to load TLS credentials: %v", err)
	}*/

	creds, err := credentials.NewClientTLSFromFile(*cert, "")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial(endpoint2, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}

	options = append(options, option.WithGRPCConn(conn))
	return options
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed server's certificate
	pemServerCA, err := ioutil.ReadFile(*cert)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	// Create the credentials and return it
	config := &tls.Config{
		RootCAs: certPool,
	}

	return credentials.NewTLS(config), nil
}

func newChannelProviderHelper() (grpc.DialOption, error) {
	// Load the certificate file.
	/*certData, err := os.ReadFile(*cert)
	if err != nil {
		return nil, err
	}

	// Create a pool of trusted certificates.
	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM(certData)
	if !ok {
		return nil, err
	}*/

	creds, err := credentials.NewClientTLSFromFile(*cert, "")
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}

	return grpc.WithTransportCredentials(creds), nil
}
