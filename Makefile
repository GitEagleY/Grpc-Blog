blog:
	protoc -I./proto --go_opt=module=go-blog --go_out=. --go-grpc_opt=module=go-blog --go-grpc_out=. ./proto/*.proto
	go build -o bin/server ./server
	go build -o bin/client ./client