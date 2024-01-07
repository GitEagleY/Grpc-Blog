package main

import (
	"context"
	pb "go-blog/proto"
	"log"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("readBlog was invoked")

	req := &pb.BlogId{Id: id}

	res, err := c.ReadBlog(context.Background(), req)

	if err != nil {
		log.Printf("error happened while reading: %v\n", err)
		return nil
	}

	log.Printf("Blog was read with: %v\n", res)
	return res
}
