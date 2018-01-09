package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"log"
	"net/http"
	"testing"

	"golang.org/x/net/http2"

	pb "github.com/kh411d/grpcrest/protobuff/helloworld"
	"google.golang.org/grpc"
)

func BenchmarkGRPC(b *testing.B) {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	for i := 0; i < b.N; i++ {
		c.SayHello(context.Background(), &pb.HelloRequest{Name: "John Doe"})
	}

}

func BenchmarkREST(b *testing.B) {
	/*client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}*/

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	http2.ConfigureTransport(tr)

	client := &http.Client{
		Transport: tr,
	}
	buf := bytes.NewBufferString(`
		{
			"name":"John Doe"
		}
	`)
	// run http posts against it
	for i := 0; i < b.N; i++ {
		client.Post("https://localhost:9090", "application/json", buf)
	}
}
