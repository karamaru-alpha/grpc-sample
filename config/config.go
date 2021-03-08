package config

import (
	"fmt"
	"os"
)

// Port サーバのport番号を取得
func Port() string {
	port := os.Getenv("Port")
	if port == "" {
		port = "8080"
	}
	return port
}

// Endpoint サーバのエンドポイント
func Endpoint() string {
	return fmt.Sprintf("localhost:%s", Port())
}
