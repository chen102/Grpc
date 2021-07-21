package main

import (
	pb "GRPC/pb"
	"context"
	"fmt"

	//wrapper "github.com/golang/protobuf/ptypes/wrappers"
	wrapper "google.golang.org/protobuf/types/known/wrapperspb"
)

type server struct {
	orderMap map[string]*pb.Order
}

func (s *server) AddOrder(ctx context.Context, orderReq *pb.Order) (*wrapper.StringValue, error) {
	fmt.Println("Order Added Id: ", orderReq.Id)
	if s.orderMap == nil {
		s.orderMap = make(map[string]*pb.Order)
	}
	s.orderMap[orderReq.Id] = orderReq
	return &wrapper.StringValue{Value: "Order Added:" + orderReq.Id}, nil
}
func (s *server) GetOrder(ctx context.Context, orderId *wrapper.StringValue) (*pb.Order, error) {
	ord, ok := s.orderMap[orderId.Value]
	if ok {
		return ord, nil
	}
	return nil, fmt.Errorf("Order does not exist", orderId)
}
