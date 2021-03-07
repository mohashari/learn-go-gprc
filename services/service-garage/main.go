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

var localStorage *model.GarageListByUser

func init() {
	localStorage = new(model.GarageListByUser)
	localStorage.List = make(map[string]*model.GarageList, 0)
}

type GaragegesServer struct{}

func (GaragegesServer) Add(ctx context.Context, param *model.GarageAndUserID) (*empty.Empty, error) {
	userID := param.UserId
	garage := param.Garage
	if _, ok := localStorage.List[userID]; !ok {
		localStorage.List[userID] = new(model.GarageList)
		localStorage.List[userID].List = make([]*model.Garage, 0)
	}
	localStorage.List[userID].List = append(localStorage.List[userID].List, garage)
	log.Println("Adding garage", garage.String(), "for user", userID)
	return new(empty.Empty), nil
}

func (GaragegesServer) List(ctx context.Context, param *model.GarageUserID) (*model.GarageList, error) {
	userID := param.UserId
	return localStorage.List[userID], nil
}

func main() {
	srv := grpc.NewServer()
	var garageSrv GaragegesServer
	model.RegisterGaragesServer(srv, garageSrv)
	log.Println("Starting RPG server at", config.SERVICE_GARAGE_PORT)

	l, err := net.Listen("tcp", config.SERVICE_GARAGE_PORT)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", config.SERVICE_GARAGE_PORT, err)
	}
	log.Fatal(srv.Serve(l))
}
