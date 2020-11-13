package main

import (
	"context"
	"log"
	"net"

	pb "github.com/Asuha-a/URLShortener/api/pb/user"
	"github.com/Asuha-a/URLShortener/api/services/user/db"
	jwt "github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
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
	log.Printf("Received Password: %v", in.GetPassword())
	return &pb.LoginReply{Token: in.GetEmail() + in.GetPassword()}, nil
}

func (s *server) Signup(ctx context.Context, in *pb.SignupRequest) (*pb.SignupReply, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(in.GetPassword()), bcrypt.MinCost)
	if err != nil {
		panic(err)
	}

	user := db.User{UUID: uuid.NewV4(), EMAIL: string(in.GetEmail()), PASSWORD: string(hash), PERMISSION: "normal"}
	result := db.DB.Create(&user)
	if result.Error != nil {
		panic(result.Error)
	}

	mySingningKey := []byte("AllYourBase")
	type UserClaims struct {
		UUID       uuid.UUID
		PERMISSION string
		jwt.StandardClaims
	}
	claims := UserClaims{
		user.UUID,
		user.PERMISSION,
		jwt.StandardClaims{
			ExpiresAt: 15000,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySingningKey)
	if err != nil {
		panic(err)
	}

	return &pb.SignupReply{Token: ss}, nil
}

func main() {
	db.Init()
	defer db.Close()

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
