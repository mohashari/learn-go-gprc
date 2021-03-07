package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/mohashari/learn-grpc/common/config"
	"github.com/mohashari/learn-grpc/common/model"
	"google.golang.org/grpc"
)

func serverGarage() model.GaragesClient {
	port := config.SERVICE_GARAGE_PORT
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}
	return model.NewGaragesClient(conn)
}

func serverUser() model.UsersClient {
	port := config.SERVICE_USER_PORT
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}
	return model.NewUsersClient(conn)
}

func main() {
	user1 := model.User{
		Id:       "n001",
		Name:     "Moh Ashari Muklis",
		Password: "1234 jkl",
		Gender:   model.UserGender(model.UserGender_value["MALE"]),
	}

	user2 := model.User{
		Id:       "n001",
		Name:     "Moh Ashari Muklis",
		Password: "1234 jkl",
		Gender:   model.UserGender(model.UserGender_value["MALE"]),
	}

	garage1 := model.Garage{
		Id:   "q001",
		Name: "Quel thalas",
		Coordinate: &model.GarageCoordinate{
			Latitude:  45.09987,
			Longitude: 54.00000,
		},
	}

	user := serverUser()

	fmt.Println("\n", "==============> user test")

	//register user1
	user.Register(context.Background(), &user1)

	//register user2
	user.Register(context.Background(), &user2)

	rest1, err := user.List(context.Background(), new(empty.Empty))
	if err != nil {
		log.Fatal(err.Error())
	}
	rest1String, _ := json.Marshal(rest1.List)
	log.Println(string(rest1String))

	garage := serverGarage()

	fmt.Println("\n", "==============> garage test")

	garage.Add(context.Background(), &model.GarageAndUserID{
		UserId: user1.Id,
		Garage: &garage1,
	})
	res2, err := garage.List(context.Background(), &model.GarageUserID{UserId: user1.Id})
	if err != nil {
		log.Fatal(err.Error())
	}
	res2string, _ := json.Marshal(res2.List)
	log.Println(string(res2string))
}
