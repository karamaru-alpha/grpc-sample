package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/karamaru-alpha/grpc-sample/chatserver"
	"github.com/karamaru-alpha/grpc-sample/config"
)

func main() {
	// リスナー登録
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", config.Port()))
	if err != nil {
		log.Fatalf("Could not listen @ %v :: %v", config.Port(), err)
	}
	log.Printf("Listening @ :%s", config.Port())

	// gRPC作成
	grpcserver := grpc.NewServer()
	cs := chatserver.ChatServer{}
	chatserver.RegisterServicesServer(grpcserver, &cs)

	// サーバ起動
	if err = grpcserver.Serve(listen); err != nil {
		log.Fatalf("Failed to start gRPC Server :: %v", err)
	}
}
