package chatserver

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

type message struct {
	id       int
	name     string
	body     string
	clientID int
}

var messageHandler = struct {
	messages []message
	mu       sync.Mutex
}{}

// ChatServer tream(channel)でメッセージの送受信を行うサーバ
type ChatServer struct{}

// ChatService tream(channel)でメッセージの送受信を行う
func (*ChatServer) ChatService(s Services_ChatServiceServer) error {

	// クライアントにユニークIDを振る
	clientID := rand.Intn(1e6)

	// stream(channel)でメッセージの送受信を行う。エラー時に接続を切る。
	errch := make(chan error)
	go receiveFromStream(s, clientID, errch)
	go sendToStream(s, clientID, errch)

	return <-errch

}

func receiveFromStream(s Services_ChatServiceServer, clientID int, errch chan error) {
	for {
		// メッセージを受信
		mssg, err := s.Recv()
		if err != nil {
			log.Printf("Error in receiving message from client :: %v", err)
			errch <- err
		}

		// 処理内でデータの一貫性を持たせるためにLock
		messageHandler.mu.Lock()

		// メッセージリストに新規メッセージを追加
		newMessage := message{
			id:       rand.Intn(1e8),
			name:     mssg.Name,
			body:     mssg.Body,
			clientID: clientID,
		}
		messageHandler.messages = append(messageHandler.messages, newMessage)

		// 新規メッセージをlogに表示
		log.Printf("%+v", newMessage)

		// Lock解除
		messageHandler.mu.Unlock()
	}
}

func sendToStream(s Services_ChatServiceServer, clientID int, errch chan error) {
	for {
		time.Sleep(1 * time.Second)

		// 処理内でデータの一貫性を持たせるためにLock
		messageHandler.mu.Lock()

		if len(messageHandler.messages) == 0 {
			messageHandler.mu.Unlock()
			continue
		}

		// 最終送信メッセージ情報
		senderID := messageHandler.messages[0].clientID
		senderName := messageHandler.messages[0].name
		sendedMessage := messageHandler.messages[0].body

		messageHandler.mu.Unlock()

		// 送信者にはｌlogが表示されないように
		if senderID != clientID {
			err := s.Send(&FromServer{Name: senderName, Body: sendedMessage})

			if err != nil {
				errch <- err
			}

			messageHandler.mu.Lock()

			if len(messageHandler.messages) > 1 {
				messageHandler.messages = messageHandler.messages[1:]
			} else {
				messageHandler.messages = []message{}
			}

			messageHandler.mu.Unlock()
		}
	}
}
