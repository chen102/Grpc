package main

import (
	pb "GRPC/pb"
	"context"
	"fmt"
	"io"
	"time"

	"google.golang.org/grpc"
	empty "google.golang.org/protobuf/types/known/emptypb"
	wrapper "google.golang.org/protobuf/types/known/wrapperspb"
)

func main() {

	conn, err := grpc.Dial("127.0.0.1:9003", grpc.WithInsecure()) //跳过证书验证
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := pb.NewOrderManagementClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//初始化
	res, err := c.InitOrder(ctx, &empty.Empty{})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
	//加入测试数据
	orders := InitOrder()
	for _, order := range orders {
		if order == nil {
			break
		}
		r, err := c.AddOrder(ctx, order)
		if err != nil {
			panic(err)
		}
		fmt.Println(r.Value, " ADD OK")
	}
	//模拟处理订单
	steamProcOrder, err := c.ProcessOrder(ctx)
	if err != nil {
		panic(err)
	}
	channel := make(chan struct{})
	go revComb(steamProcOrder, channel)
	for _, orderId := range []string{"10", "12", "13", "14"} {
		err := steamProcOrder.Send(&wrapper.StringValue{Value: orderId})
		if err != nil {
			panic(err)
		}
	}
	//time.Sleep(time.Millisecond * 1000) //若延迟，服务端会认为客户端已经发送完毕，关闭调用
	for _, orderId := range []string{"16", "17", "18"} {
		err := steamProcOrder.Send(&wrapper.StringValue{Value: orderId})
		if err != nil {
			panic(err)
		}
	}
	if err := steamProcOrder.CloseSend(); err != nil {
		panic(err)
	}
	channel <- struct{}{}
}

func revComb(steamProcOrder pb.OrderManagement_ProcessOrderClient, c chan struct{}) {
	for {
		comb, err := steamProcOrder.Recv() //开始接收服务端的响应，直到EOF
		if err == io.EOF {
			break
		}
		if comb == nil {
			break
		}
		fmt.Println("COMB: ", comb)

	}
	<-c
}

func InitOrder() []*pb.Order {
	orders := make([]*pb.Order, 7)
	orders[0] = &pb.Order{Id: "10", Items: []string{"apple", "strawberry"}, Description: "fruit", Price: 29.21, Destination: "Beijing"}
	orders[1] = &pb.Order{Id: "12", Items: []string{"watermelon", "orange"}, Description: "fruit", Price: 26.62, Destination: "Henan"}
	orders[2] = &pb.Order{Id: "13", Items: []string{"apple", "strawberry"}, Description: "fruit", Price: 22.61, Destination: "Beijing"}
	orders[3] = &pb.Order{Id: "14", Items: []string{"strawberry", "orange"}, Description: "fruit", Price: 27.25, Destination: "Henan"}
	orders[4] = &pb.Order{Id: "16", Items: []string{"strawberry", "orange"}, Description: "fruit", Price: 27.25, Destination: "Beijing"}
	orders[5] = &pb.Order{Id: "17", Items: []string{"strawberry", "orange"}, Description: "fruit", Price: 27.25, Destination: "Wuhan"}
	orders[6] = &pb.Order{Id: "18", Items: []string{"strawberry", "orange"}, Description: "fruit", Price: 27.25, Destination: "Wuhan"}
	return orders
}
