package service

import (
	pb "GRPC/pb"
	"context"
	"fmt"
	empty "google.golang.org/protobuf/types/known/emptypb"
	wrapper "google.golang.org/protobuf/types/known/wrapperspb"
	"io"
	"strings"
)

const OrderSize = 3 //一次最多处理订单数
type OrderService struct {
	orderMap map[string]*pb.Order
}

func (s *OrderService) InitOrder(ctx context.Context, null *empty.Empty) (*wrapper.StringValue, error) {
	if s.orderMap != nil {
		return &wrapper.StringValue{Value: "order already init"}, nil
	}
	s.orderMap = make(map[string]*pb.Order)
	return &wrapper.StringValue{Value: "order init ok"}, nil

}
func (s *OrderService) AddOrder(ctx context.Context, orderReq *pb.Order) (*wrapper.StringValue, error) {
	if s.orderMap == nil {
		//return nil, fmt.Errorf("order is NULL,please initOrder")
		s.orderMap = make(map[string]*pb.Order)
	}
	fmt.Println("Order Added Id: ", orderReq.Id)
	s.orderMap[orderReq.Id] = orderReq
	return &wrapper.StringValue{Value: "Order Added:" + orderReq.Id}, nil
}
func (s *OrderService) GetOrder(ctx context.Context, orderId *wrapper.StringValue) (*pb.Order, error) {
	if s.orderMap == nil {
		return nil, fmt.Errorf("Order is Null,please initOrder")
	}
	ord, ok := s.orderMap[orderId.Value]
	if ok {
		return ord, nil
	}
	return nil, fmt.Errorf("Order does not exist", orderId)
}
func (s *OrderService) SearchOrder(searchQuery *wrapper.StringValue, stream pb.OrderManagement_SearchOrderServer) error {
	if s.orderMap == nil {
		return fmt.Errorf("order is NULL,please initOrder")
	}
	if searchQuery.Value == " " { //若查询的是空返回全部订单
		for _, order := range s.orderMap {

			err := stream.Send(order) //在流中发送匹配订单
			if err != nil {
				return fmt.Errorf("send error:", err)
			}
		}
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
func (s *OrderService) UpdateOrder(stream pb.OrderManagement_UpdateOrderServer) error {
	if s.orderMap == nil {
		return fmt.Errorf("order is NULL,please initOrder")
	}
	for {
		order, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&wrapper.StringValue{Value: "is all ok"})
		}
		_, ok := s.orderMap[order.Id]
		if ok {
			fmt.Println("Order ID", order.Id, "Updated OK")
		} else {
			s.orderMap[order.Id] = order
			fmt.Println("Order ID", order.Id, "ADD OK")
		}

	}
}
func (s *OrderService) ProcessOrder(stream pb.OrderManagement_ProcessOrderServer) error {
	if s.orderMap == nil {
		return fmt.Errorf("order is NULL,please initOrder")
	}
	var combinedShipmentMap = make(map[string]pb.CombinedShipment)
	var batchMark = 1 //目前批次中订单处理个数
	for {
		orderId, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				for _, comb := range combinedShipmentMap {
					stream.Send(&comb)
				}
				return nil
			}
			fmt.Println("ERROR: ", err)
			return err
		}
		fmt.Println("Doing: ", orderId.GetValue())
		//防空指针
		if _, ok := s.orderMap[orderId.GetValue()]; !ok {
			return fmt.Errorf(orderId.GetValue(), ": this order is null")
		}

		//订单成批处理
		//将目地的相同的订单进行组合，发送给客户端
		//例如3个订单一批，从客户端受到3个订单地址分别为x、x、y 即x、x一组发，y一组发

		//防空指针
		if s.orderMap[orderId.GetValue()].Destination == "" {
			return fmt.Errorf(orderId.GetValue(), ": this order destination is Null")
		}
		destination := s.orderMap[orderId.GetValue()].Destination
		tempDestination, ok := combinedShipmentMap[destination]
		if ok { //若该批次中这个地址已经出现过一次，对那个进行累加
			ord := s.orderMap[orderId.GetValue()]
			tempDestination.OrderList = append(tempDestination.OrderList, ord) //tempDestination是已有的地址
			combinedShipmentMap[destination] = tempDestination
		} else { //该批次中第一次出现的地址，将他加入map
			first := pb.CombinedShipment{Id: s.orderMap[orderId.GetValue()].Destination, Status: "processed!"}
			ord := s.orderMap[orderId.GetValue()]
			first.OrderList = append(tempDestination.OrderList, ord) //tempDestination是空的
			combinedShipmentMap[destination] = first
		}
		if batchMark == OrderSize {
			for _, comb := range combinedShipmentMap { //将本批次的订单组合全部发个客户端
				fmt.Println("Shipping : ", comb.Id, len(comb.OrderList))
				err := stream.Send(&comb)
				if err != nil {
					return err
				}
				combinedShipmentMap = make(map[string]pb.CombinedShipment) //重置组合Map
				batchMark = 0
			}

		} else {
			batchMark++
		}

	}
}
