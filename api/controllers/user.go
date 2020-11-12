package controllers

import (
	"context"
	"log"
	"time"

	pb "github.com/Asuha-a/URLShortener/api/pb/user"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

const (
	address = "user:50051"
)

// Login app
func Login(c *gin.Context) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewAuthClient(conn)
	log.Println("test logging")
	email := c.Query("email")
	password := c.Query("password")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.Login(ctx, &pb.LoginRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		log.Fatalf("could not login: %v", err)
	}
	log.Printf("Token: %s", r.GetToken())
}

// Signup app
func Signup(c *gin.Context) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewAuthClient(conn)

	email := c.Query("email")
	password := c.Query("password")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.Signup(ctx, &pb.SignupRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		log.Fatalf("could not signup: %v", err)
	}
	log.Printf("Token: %s", r.GetToken())
}
