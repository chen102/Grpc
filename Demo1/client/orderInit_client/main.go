package main

import (
	pb "GRPC/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	empty "google.golang.org/protobuf/types/known/emptypb"
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
	res, err := c.InitOrder(ctx, &empty.Empty{})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
