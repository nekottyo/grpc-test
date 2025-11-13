package main

import (
	"context"
	"fmt"

	pb "github.com/nekottyo/grpc-test/pkg/time"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	address := "localhost:50051"
	conn, err := grpc.NewClient(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return
	}
	defer conn.Close()

	client := pb.NewTimeServiceClient(conn)

	req := &pb.GetCurrentTimeRequest{}
	res, err := client.GetCurrentTime(context.Background(), req)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res.GetDate())
	}
}
