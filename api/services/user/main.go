package main

import (
	"context"
	"log"
	"net"

	pb "github.com/Asuha-a/URLShortener/api/pb/user"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedAuthServer
}

func (s *server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginReply, error) {
	log.Printf("Received Email: %v", in.GetEmail())
	log.Printf("Received Email: %v", in.GetPassword())
	return &pb.LoginReply{Token: in.GetEmail() + in.GetPassword()}, nil
}

func (s *server) Signup(ctx context.Context, in *pb.SignupRequest) (*pb.SignupReply, error) {
	log.Printf("Received Email: %v", in.GetEmail())
	log.Printf("Received Email: %v", in.GetPassword())
	return &pb.SignupReply{Token: in.GetEmail() + in.GetPassword()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAuthServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
