package controllers

import (
	"context"
	"log"
	"time"

	pbUser "github.com/Asuha-a/URLShortener/api/pb/user"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

const (
	userAddress = "user:50051"
)

// Login app
func Login(c *gin.Context) {
	conn, err := grpc.Dial(userAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pbUser.NewAuthClient(conn)

	email := c.Query("email")
	password := c.Query("password")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.Login(ctx, &pbUser.LoginRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		panic(err)
		c.AbortWithStatus(400)
	} else {
		c.JSON(200, r.GetToken())
	}
}

// Signup app
func Signup(c *gin.Context) {
	conn, err := grpc.Dial(userAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pbUser.NewAuthClient(conn)

	email := c.Query("email")
	password := c.Query("password")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.Signup(ctx, &pbUser.SignupRequest{
		Email:    email,
		Password: password,
	})

	if err != nil {
		panic(err)
		c.AbortWithStatus(400)
	} else {
		c.JSON(200, r.GetToken())
	}
}
