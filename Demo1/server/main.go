package main

import (
	"GRPC/Demo1/server/service"
	pb "GRPC/pb"
	"google.golang.org/grpc"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":9003")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterOrderManagementServer(s, &service.OrderService{})
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
