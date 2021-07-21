package main

import (
	pb "GRPC/pb"
	"context"
	"fmt"
	"strings"

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
func (s *server) SearchOrder(searchQuery *wrapper.StringValue, stream pb.OrderManagement_SearchOrderServer) error {
	if s.orderMap == nil {
		return fmt.Errorf("Order is Null")
	}
	for orderid, order := range s.orderMap {
		for _, itemStr := range order.Items {
			if strings.Contains(itemStr, searchQuery.Value) { //查找匹配订单：判断字符串s是否包含子串substr。
				err := stream.Send(order) //在流中发送匹配订单
				if err != nil {
					return fmt.Errorf("send error:", err)
				}
				fmt.Println("OK orderId:", orderid)
				break
			}
		}
	}
	return nil
}
