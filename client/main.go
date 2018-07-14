package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/itosho/go-grpc-sample/idol"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8046", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()
	client := pb.NewIdolClient(conn)
	message := &pb.GetIdolMessage{TargetIdol: "sayuri"}

	res, err := client.GetIdol(context.TODO(), message)
	fmt.Printf("result:%v \n", res)
	fmt.Printf("error::%v \n", err)
}
