package main

import (
	"context"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/mohashari/learn-grpc/common/config"
	"github.com/mohashari/learn-grpc/common/model"
	"google.golang.org/grpc"
)

var localStorage *model.UserList

func init() {
	localStorage = new(model.UserList)
	localStorage.List = make([]*model.User, 0)
}

type UserService struct{}

func (UserService) Register(ctx context.Context, param *model.User) (*empty.Empty, error) {
	localStorage.List = append(localStorage.List, param)
	log.Println("Registering user", param.String())
	return new(empty.Empty), nil
}

func (UserService) List(ctx context.Context, void *empty.Empty) (*model.UserList, error) {
	return localStorage, nil
}

func main() {
	srv := grpc.NewServer()
	var userSrv UserService
	model.RegisterUsersServer(srv, userSrv)
	log.Println("starting RPC server at", config.SERVICE_USER_PORT)

	l, err := net.Listen("tcp", config.SERVICE_USER_PORT)
	if err != nil {
		log.Fatalf("could not lister to %s: %v", config.SERVICE_USER_PORT, err)
	}
	log.Fatal(srv.Serve(l))
}
