package main

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"time"

	pb "github.com/Asuha-a/URLShortener/api/pb/url"
	"github.com/Asuha-a/URLShortener/api/services/url/db"
	"github.com/Asuha-a/URLShortener/api/utility"
	"github.com/golang/protobuf/ptypes"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc"
)

const (
	port = ":50052"
)

type server struct {
	pb.UnimplementedURLServer
}

/*
func (s *server) GetAllURL(rect pb.Rectangle, stream pb.URLGetAllURLServer) error {
	var urls []db.URL
	uuid, permission, err := utility.ParseJWT(string(in.GetToken()))
	if err != nil {
		panic(err)
	}
	result := db.DB.Where("uuid <> ?", in.GetUuid()).Find(&urls)
	if result.Error != nil {
		panic(result.Error)
	}
	return &pb.GetAllURLReply{
		Token: ss,
	}, nil
}
*/

func (s *server) PostURL(ctx context.Context, in *pb.PostURLRequest) (*pb.PostURLReply, error) {
	uuidUser, _, err := utility.ParseJWT(string(in.GetToken()))
	if err != nil {
		panic(err)
	}
	redirectTo, err := json.Marshal(in.GetRedirectTo())
	if err != nil {
		panic(err)
	}
	url := db.URL{UUID: uuid.NewV4(), UserID: uuidUser, Title: string(in.GetTitle()), URL: utility.RandStringRunes(8), RedirectTo: redirectTo, CreatedAt: time.Now()}
	result := db.DB.Create(&url)
	if result.Error != nil {
		panic(result.Error)
	}
	CreatedAt, err := ptypes.TimestampProto(url.CreatedAt)
	if err != nil {
		panic(err)
	}
	node := pb.Node{}
	jsonData := []byte(url.RedirectTo)
	err = json.Unmarshal(jsonData, &node)
	if err != nil {
		panic(err)
	}
	return &pb.PostURLReply{
		Uuid:       url.UUID.String(),
		UserId:     url.UserID.String(),
		Title:      url.Title,
		Url:        url.URL,
		RedirectTo: &node,
		CreatedAt:  CreatedAt,
	}, nil
}

func main() {
	db.Init()
	defer db.Close()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterURLServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
