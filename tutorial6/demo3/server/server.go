package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"

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

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------

// 服务端流式 rpc
type LogStreamService struct {
	pb.UnimplementedLogStreamServiceServer
}

func (s *LogStreamService) GetStreamLogs(req *pb.LogStreamReq, srv pb.LogStreamService_GetStreamLogsServer) error {
	// 模拟 10s 发送完所有的数据
	timeout := time.After(5 * time.Second)
	for {
		select {
		case <-timeout:
			fmt.Println("finished...")
			return nil
		default:
			logMessage := generateLogMessage(req.LogLevel)
			if err := srv.Send(&pb.LogStreamRes{
				LogLevel:   req.LogLevel,
				LogMessage: logMessage,
			}); err != nil {
				return err
			}
		}
		// 模拟日志的间隔时间
		time.Sleep(1 * time.Second)
	}
}

func generateLogMessage(logLevel int32) string {
	return fmt.Sprintf("Log message with level %d", logLevel)
}

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------
type FileService struct {
	pb.UnimplementedFileServiceServer
}

func (s *FileService) UploadFile(stream pb.FileService_UploadFileServer) error {
	filePath := "./file/movie.mkv"
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			log.Println("File uploaded successfully")
			return stream.SendAndClose(&pb.UploadStatus{Success: true})
		}
		if err != nil {
			return err
		}
		_, err = file.Write(chunk.Data)
		if err != nil {
			return err
		}
	}

}

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------
type ChatService struct {
	pb.UnimplementedChatServiceServer
}

func (s *ChatService) Chat(stream pb.ChatService_ChatServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("[%s]: %s", msg.Sender, msg.Text)
		stream.Send(&pb.Message{Sender: "Server", Text: msg.Text})
	}
}

func main() {
	demo4()
}

func demo1() {
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

func demo2() {
	flag.Parse()
	// 1. 监听 tcp 端口
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// 2. 实例化 grpc
	s := grpc.NewServer()
	// 3. 在 grpc 上注册微服务
	pb.RegisterLogStreamServiceServer(s, &LogStreamService{})
	// 4. 启动服务端
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func demo3() {
	flag.Parse()
	// 1. 监听 tcp 端口
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// 2. 实例化 grpc
	s := grpc.NewServer()
	// 3. 在 grpc 上注册微服务
	pb.RegisterFileServiceServer(s, &FileService{})
	// 4. 启动服务端
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func demo4() {
	flag.Parse()
	// 1. 监听 tcp 端口
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// 2. 实例化 grpc
	s := grpc.NewServer()
	// 3. 在 grpc 上注册微服务
	pb.RegisterChatServiceServer(s, &ChatService{})
	// 4. 启动服务端
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
