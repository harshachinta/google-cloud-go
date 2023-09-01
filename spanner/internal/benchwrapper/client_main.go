package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "cloud.google.com/go/spanner/internal/benchwrapper/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/local"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("addr", "localhost:8082", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.example.com", "The server name used to verify the hostname returned by the TLS handshake")
)

// Read gets the feature for the given point.
func Read(client pb.SpannerBenchWrapperClient, point *pb.ReadQuery) {
	log.Printf("Getting feature for point (%v)", point.Query)
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	feature, err := client.Read(ctx, point)
	if err != nil {
		log.Fatalf("client.GetFeature failed: %v", err)
	}
	log.Println(feature)
}

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	// /usr/local/google/home/sriharshach/Downloads/
	opts = append(opts, grpc.WithTransportCredentials(local.NewCredentials()))
	// grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewSpannerBenchWrapperClient(conn)

	// Looking for a valid feature
	Read(client, &pb.ReadQuery{Query: "SELECT id, value, counter, br_id FROM T;"})
}
