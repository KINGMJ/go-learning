package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/KINGMJ/go-learning/tutorial6/demo3/proto/pb"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// 定义空接口
type UserInfoService struct {
	pb.UnimplementedUserInfoServiceServer
}

func (s *UserInfoService) GetUserInfo(ctx context.Context, req *pb.UserRequest) (res *pb.UserResponse, err error) {
	// 通过用户名查询用户信息
	name := req.Name
	// 模拟数据库里查询用户信息
	if name == "zs" {
		res = &pb.UserResponse{
			Id:    1,
			Name:  name,
			Age:   22,
			Hobby: []string{"吃饭", "睡觉", "打豆豆"},
		}
	}
	return
}

func main() {
	flag.Parse()
	// 1. 监听 tcp 端口
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// 2. 实例化 grpc
	s := grpc.NewServer()
	// 3. 在 grpc 上注册微服务
	pb.RegisterUserInfoServiceServer(s, &UserInfoService{})

	// 4. 启动服务端
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
