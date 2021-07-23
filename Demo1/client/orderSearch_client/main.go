package main

import (
	pb "GRPC/pb"
	"context"
	"fmt"
	"io"
	"time"

	"google.golang.org/grpc"
	wrapper "google.golang.org/protobuf/types/known/wrapperspb"
)

func main() {
	//实现订单模糊匹配
	//输入相关字符，匹配对应的订单
	conn, err := grpc.Dial("127.0.0.1:9003", grpc.WithInsecure()) //跳过证书验证
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := pb.NewOrderManagementClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
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
	//服务端往流中写入了多个响应，客户端也需处理多个响应
	searchStream, err := c.SearchOrder(ctx, &wrapper.StringValue{Value: "strawberry"}) //SearchOrder返回OrderManagement_SearchOrderClient的客户端流，他有一个recv的方法
	if err != nil {
		panic(err)
	}
	for {
		searchorder, err := searchStream.Recv()
		if err == io.EOF { //发现流结束时，Recv返回io.EOF
			break
		}
		fmt.Println(searchorder)
	}

}

//加一些测试数据
func InitOrder() []*pb.Order {
	orders := make([]*pb.Order, 10)
	orders[0] = &pb.Order{Id: "116", Items: []string{"apple", "strawberry"}, Description: "fruit", Price: 29.21}
	orders[1] = &pb.Order{Id: "112", Items: []string{"watermelon", "orange"}, Description: "fruit", Price: 26.62}
	orders[2] = &pb.Order{Id: "113", Items: []string{"apple", "strawberry"}, Description: "fruit", Price: 22.61}
	orders[3] = &pb.Order{Id: "114", Items: []string{"strawberry", "orange"}, Description: "fruit", Price: 27.25}
	orders[4] = &pb.Order{Id: "115", Items: []string{"pear", "orange", "apple"}, Description: "fruit", Price: 35.21}
	return orders
}
