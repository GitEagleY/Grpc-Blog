package main

import (
	"context"
	"log"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"

	pb "go-blog/proto"
)

var collection *mongo.Collection

const address = "0.0.0.0:50051"

type Server struct {
	pb.BlogServiceServer
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic: %v", r)
		}
	}()

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("blogdb").Collection("blog")

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	defer lis.Close()
	log.Printf("Listening at %s", address)

	server := grpc.NewServer()

	pb.RegisterBlogServiceServer(server, &Server{})
	log.Println("Service server registered")

	log.Println("Starting gRPC server...")
	err = server.Serve(lis)

	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
