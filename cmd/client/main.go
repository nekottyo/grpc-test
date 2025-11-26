package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"

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

	args := os.Args
	defer conn.Close()

	client := pb.NewTimeServiceClient(conn)

	switch args[1] {
	case "stream":
		GetCurrentTimeStream(client)
	default:
		GetCurrentTime(client)
	}
}

func GetCurrentTime(client pb.TimeServiceClient) {
	req := &pb.GetCurrentTimeRequest{}
	res, err := client.GetCurrentTime(context.Background(), req)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res.GetDate())
	}

}

func GetCurrentTimeStream(client pb.TimeServiceClient) {
	req := &pb.GetCurrentTimeRequest{}
	stream, err := client.GetCurrentTimeStream(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		res, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res)
	}
}
