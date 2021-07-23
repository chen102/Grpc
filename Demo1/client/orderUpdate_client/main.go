package main

import (
	pb "GRPC/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
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
	updatesteam, err := c.UpdateOrder(ctx)
	orders := InitOrder()
	for _, order := range orders {
		err := updatesteam.Send(order)
		if err != nil {
			panic(err)
		}
	}
	updateRes, err := updatesteam.CloseAndRecv() //关闭流并接收响应
	if err != nil {
		panic(err)
	}
	fmt.Println(updateRes)

}

//加一些测试数据
func InitOrder() []*pb.Order {
	orders := make([]*pb.Order, 2)
	orders[0] = &pb.Order{Id: "114", Items: []string{"apple", "orange"}, Description: "fruit", Price: 12.52}
	orders[1] = &pb.Order{Id: "115", Items: []string{"water", "orange"}, Description: "fruit", Price: 15.21}
	return orders
}
