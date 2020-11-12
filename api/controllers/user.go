package controllers

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/Asuha-a/URLShortener/api/pb/user"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

const (
	address     = "user:50051"
	defaultName = "world"
)

// Hello user microservice
func Hello(c *gin.Context) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
