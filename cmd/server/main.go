/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/nekottyo/grpc-test/pkg/time"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedTimeServiceServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) GetCurrentTime(_ context.Context, in *pb.GetCurrentTimeRequest) (*pb.GetCurrentTimeResponse, error) {
	now := time.Now()
	log.Printf("response to %s", now)
	return &pb.GetCurrentTimeResponse{Date: now.String()}, nil
}

func (s *server) GetCurrentTimeStream(in *pb.GetCurrentTimeRequest, stream pb.TimeService_GetCurrentTimeStreamServer) error {
	count := 5
	for i := 0; i < count; i++ {
		if err := stream.Send(&pb.GetCurrentTimeResponse{
			Date: time.Now().String(),
		}); err != nil {
			return err
		}
		time.Sleep(time.Second * 1)
	}
	return nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTimeServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
