package main

import (
	"context"
	"errors"
	"log"
	"net"

	pb "github.com/itosho/go-grpc-sample/idol"
	"google.golang.org/grpc"
)

type IdolService struct {
}

func (s *IdolService) GetIdol(ctx context.Context, message *pb.GetIdolMessage) (*pb.IdolResponse, error) {
	switch message.TargetIdol {
	case "nanase":
		return &pb.IdolResponse{
			Name:     "Nanase NISHINO",
			NickName: "Nachan",
			Height:   158,
		}, nil
	case "sayuri":
		return &pb.IdolResponse{
			Name:     "Sayuru MATSUMURA",
			NickName: "Sayurin",
			Height:   162,
		}, nil
	}

	return nil, errors.New("Not Found Idol")
}

func main() {
	listenPort, err := net.Listen("tcp", ":8046")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()

	pb.RegisterIdolServer(server, &IdolService{})
	server.Serve(listenPort)
}
