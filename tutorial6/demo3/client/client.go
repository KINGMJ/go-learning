package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/KINGMJ/go-learning/tutorial6/demo3/proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()
	// 1. 连接服务端
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// 2. 实例化 grpc 客户端
	client := pb.NewUserInfoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 3. 组装请求参数
	req := new(pb.UserRequest)
	req.Name = "zs"
	// 4. 调用接口
	res, err := client.GetUserInfo(ctx, req)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("响应结果： %v\n", res)
}
