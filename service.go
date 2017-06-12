package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"golang.org/x/net/context"

	"github.com/iochti/user-service/models"
	pb "github.com/iochti/user-service/proto"
)

// UserSvc is the implementation of gRPC User service (see proto/user.proto)
type UserSvc struct {
	Db DataLayerInterface
}

// GetUser is a RPC method that fetches customer from database
func (u *UserSvc) GetUser(ctx context.Context, in *pb.UserRequest) (*pb.UserMessage, error) {
	categ := in.GetCateg()
	user := new(models.User)
	var id int
	var err error
	// If the expected value is a number instead of a string
	if categ == "id" {
		id, err = strconv.Atoi(in.GetValue())
		if err != nil {
			return nil, err
		}
	}
	// Switch by search category
	switch categ {
	case "id":
		user, err = u.Db.GetUserByID(id)
		break
	case "email":
		user, err = u.Db.GetUserByEmail(in.GetValue())
		break
	case "login":
		user, err = u.Db.GetUserByLogin(in.GetValue())
		break
	default:
		return nil, fmt.Errorf("Error: unknown search category")
	}

	if err != nil {
		return nil, err
	}

	res := new(pb.UserMessage)
	// Encodes the user as an array of bytes, so it can be read by the client
	enc, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}
	res.User = enc
	return res, nil
}

// CreateUser is a RPC method that creates a user in the database
func (u *UserSvc) CreateUser(ctx context.Context, in *pb.UserMessage) (*pb.UserMessage, error) {
	userBytes := in.GetUser()
	var user models.User
	if err := json.Unmarshal(userBytes, &user); err != nil {
		return nil, err
	}

	if err := u.Db.CreateUser(&user); err != nil {
		return nil, err
	}

	userBytes, err := json.Marshal(&user)
	if err != nil {
		return nil, err
	}

	return &pb.UserMessage{User: userBytes}, nil
}

// DeleteUser is a RPC method that deletes a user in the database
func (u *UserSvc) DeleteUser(ctx context.Context, in *pb.UserID) (*pb.UserDeleted, error) {
	userID := in.GetId()

	if userID == "" {
		return nil, fmt.Errorf("Error, id should be > 0")
	}

	// If the deletion fails
	if err := u.Db.DeleteUser(userID); err != nil {
		return nil, err
	}
	res := new(pb.UserDeleted)
	res.Id = userID
	res.Deleted = true
	return res, nil
}
