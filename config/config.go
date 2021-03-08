package config

import "os"

// Port はサーバーのポート番号を取得します
func Port() string {
	return os.Getenv("PORT")
}
