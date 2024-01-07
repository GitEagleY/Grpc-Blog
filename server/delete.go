package main

import (
	"context"
	"fmt"
	pb "go-blog/proto"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func DeleteBlog(ctx context.Context, in *pb.BlogId) (*emptypb.Empty, error) {
	log.Printf("DeleteBlog was invoked with %v\n", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Cannot parse ID",
		)

		res, err := collection.DeleteOne(ctx, bson.M{"_id": oid})

		if err != nil {
			return nil, status.Errorf(
				codes.Internal,
				fmt.Sprintf("cannot delete object in mongodb:%v\n", err),
			)
		}

		if res.DeletedCount == 0 {
			return nil, status.Errorf(
				codes.NotFound,
				fmt.Sprintf("blog was not found:%v\n", err),
			)
		}
	}
	return &emptypb.Empty{}, nil
}
