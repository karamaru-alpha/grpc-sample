# protoのコード生成
.PHONY: gen_proto
gen_proto:
	protoc --go-grpc_out=require_unimplemented_servers=false:./chatserver/ --go_out=./chatserver/ chat.proto

# 8080番Porｔでサーバを起動する
.PHONY: server
server:
	PORT=8080 go run server/main.go


# 8080番Porｔサーバに接続するクライアントを起動する
.PHONY: client
client:
	PORT=8080 go run client/main.go
