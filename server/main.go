package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"

	pb "github.com/karamaru-alpha/grpc-sample/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.ChatServer
}

func (s *server) BiliChat(stream pb.Chat_BiliChatServer) error {
	log.Println("start new server in BiliChat!")
	ctx := stream.Context()

	for {
		// チャネルが終了した場合return。それ以外(値が入っていない場合も)処理を続行する。
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		// Streamから値を受け取る
		req, err := stream.Recv()
		// タイムアウト時などはエラーがEOFで来るらしい
		if err == io.EOF {
			log.Println("exit")
			return nil
		}
		if err != nil {
			log.Printf("receive error %v", err)
			continue
		}

		resp := pb.MessageResponse{Message: req.Message}
		if err := stream.Send(&resp); err != nil {
			log.Printf("send error %v", err)
		}
		log.Printf("send message=%s", req.Message)
	}
}

func main() {
	fmt.Print()

	// リスナー登録
	lis, err := net.Listen("tcp", ":"+os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// gRPC作成
	s := grpc.NewServer()
	pb.RegisterChatServer(s, &server{})

	// サーバ起動
	log.Printf("Server running in port: %s", os.Getenv("PORT"))
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
