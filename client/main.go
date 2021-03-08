package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/karamaru-alpha/grpc-sample/chatserver"
	"github.com/karamaru-alpha/grpc-sample/config"

	"google.golang.org/grpc"
)

func main() {

	// grpcサーバーに接続
	conn, err := grpc.Dial(config.Endpoint(), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Faile to conncet to gRPC server :: %v", err)
	}
	defer conn.Close()

	// Client登録
	client := chatserver.NewServicesClient(conn)
	stream, err := client.ChatService(context.Background())
	if err != nil {
		log.Fatalf("Failed to call ChatService :: %v", err)
	}

	ch := clientHandler{stream: stream}
	ch.setname()

	// ストリーミング通信でメッセージを送受信する
	go ch.sendMessage()
	go ch.receiveMessage()

	bl := make(chan bool)
	<-bl
}

//clientHandler
type clientHandler struct {
	stream chatserver.Services_ChatServiceClient
	name   string
}

// クライアントのユーザー名を設定する
func (ch *clientHandler) setname() {

	// ユーザー名を要求する
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Your Name : ")

	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Failed to read from console :: %v", err)
	}

	name = strings.Trim(name, "\r\n")
	if name == "" {
		log.Fatal("Failed to register name")
	}

	ch.name = strings.Trim(name, "\r\n")
}

// メッセージを送信する
func (ch *clientHandler) sendMessage() {
	for {
		// チャット文を要求する
		reader := bufio.NewReader(os.Stdin)
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf(" Failed to read from console :: %v", err)
		}

		// チャット文が空の場合、何もしないで再度チャット文要求
		message = strings.Trim(message, "\r\n")
		if message == "" {
			continue
		}

		// Streamに送信
		messageBox := &chatserver.FromClient{
			Name: ch.name,
			Body: message,
		}
		if err = ch.stream.Send(messageBox); err != nil {
			log.Printf("Error while sending message to server :: %v", err)
		}
	}
}

// メッセージを受け取る
func (ch *clientHandler) receiveMessage() {
	for {
		// Streamから受信
		message, err := ch.stream.Recv()
		if err != nil {
			log.Printf("Error in receiving message from server :: %v", err)
		}

		// 受信したメッセージ(本文+名前)を表示
		fmt.Printf("%s : %s \n", message.Name, message.Body)
	}
}
