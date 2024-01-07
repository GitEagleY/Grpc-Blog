package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "go-blog/proto"
)

const (
	address = "localhost:50051"
)

func main() {
	// Connect to addr with insecure credentials
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error connecting to gRPC server: %v", err)
		return
	}
	defer conn.Close()

	client := pb.NewBlogServiceClient(conn)

	id := createBlog(client)
	readBlog(client, id)
	updateBlog(client, id)
	listBlog(client)
	deleteBlog(client, id)
	listBlog(client)
}
