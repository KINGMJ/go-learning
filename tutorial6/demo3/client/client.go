package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/KINGMJ/go-learning/tutorial6/demo3/proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	demo4()
}

func demo1() {
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

func demo2() {
	flag.Parse()
	// 1. 连接服务端
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// 2. 实例化 grpc 客户端
	client := pb.NewLogStreamServiceClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// 3. 发起流式 rpc 请求
	stream, err := client.GetStreamLogs(ctx, &pb.LogStreamReq{LogLevel: 1})
	if err != nil {
		log.Fatalf("could not get stream logs: %v", err)
	}
	// 4. 循环接收服务端推送的日志消息
	for {
		resp, err := stream.Recv()
		// 判断消息流是否已经结束
		if err == io.EOF {
			fmt.Println("接收完毕，退出...")
			break
		}
		if err != nil {
			log.Fatalf("Failed to receive log message: %v", err)
		}
		fmt.Println("Received log message:", resp.LogMessage)
	}
}

func demo3() {
	flag.Parse()
	// 1. 连接服务端
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// 2. 实例化 grpc 客户端
	client := pb.NewFileServiceClient(conn)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 3. open file
	filePath := "../movie.mkv"
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	// 4. 发起流式 rpc 请求
	stream, err := client.UploadFile(ctx)
	if err != nil {
		log.Fatalf("Failed to upload file: %v", err)
	}

	// 5. 开始上传文件
	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Failed to read file: %v", err)
		}
		if err := stream.Send(&pb.FileChunk{Data: buffer[:n]}); err != nil {
			log.Fatalf("Failed to send chunk: %v", err)
		}
	}
	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Failed to receive response: %v", err)
	}
	if resp.Success {
		log.Println("File uploaded successfully")
	}
}

func demo4() {
	flag.Parse()
	// 1. 连接服务端
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// 2. 实例化 grpc 客户端
	client := pb.NewChatServiceClient(conn)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 3. 调用远程服务器 Chat 函数
	stream, err := client.Chat(ctx)
	if err != nil {
		log.Fatalf("Failed to create chat stream: %v", err)
	}

	// 4. 开一个协程处理服务器返回的消息
	go func() {
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Fatalf("Failed to received message: %v", err)
			}
			log.Printf("[%s]: %s", msg.Sender, msg.Text)
		}
	}()

	// 5. 开始发送消息
	log.Print("Enter your message: ")
	for {
		var text string
		if _, err := fmt.Scanln(&text); err != nil {
			log.Fatalf("Failed to read input: %v", err)
		}
		if err := stream.Send(&pb.Message{Sender: "Client", Text: text}); err != nil {
			log.Fatalf("Failed to send message : %v", err)
		}
	}
}
