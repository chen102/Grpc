package main

import (
	pb "GRPC/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	wrapper "google.golang.org/protobuf/types/known/wrapperspb"
	"time"
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
	r, err := c.AddOrder(ctx, &pb.Order{Id: "111", Items: []string{"apple", "orange"}, Description: "fruit", Price: 22.21})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Value, " ADD OK")
	order, err := c.GetOrder(ctx, &wrapper.StringValue{Value: "111"})
	if err != nil {
		panic(err)
	}
	fmt.Println(order)
}
