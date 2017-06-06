package main

import (
	"context"

	pb "github.com/iochti/user-service/proto"
)

// UserSvc is the implementation of gRPC User service (see proto/user.proto)
type UserSvc struct {
	db DataLayerInterface
}

// GetUser is a RPC method that fetches customer from database
func (u *UserSvc) GetUser(ctx context.Context, in *pb.UserRequest) (*pb.UserMessage, error) {
	return &pb.UserMessage{}, nil
}

// CreateUser is a RPC method that creates a user in the database
func (u *UserSvc) CreateUser(ctx context.Context, in *pb.UserMessage) (*pb.UserCreated, error) {
	return &pb.UserCreated{}, nil
}
