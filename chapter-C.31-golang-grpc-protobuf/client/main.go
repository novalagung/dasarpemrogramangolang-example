package main

import (
	"chapter-c30/common/config"
	"chapter-c30/common/model"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"strings"

	"google.golang.org/grpc"
)

func serviceGarage() model.GaragesClient {
	port := config.ServiceGaragePort
	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return model.NewGaragesClient(conn)
}

func serviceUser() model.UsersClient {
	port := config.ServiceUserPort
	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return model.NewUsersClient(conn)
}

func main() {
	user1 := model.User{
		Id:       "n001",
		Name:     "Noval Agung",
		Password: "kw8d hl12/3m,a",
		Gender:   model.UserGender(model.UserGender_value["MALE"]),
	}
	user2 := model.User{
		Id:       "n002",
		Name:     "Nabila Rozan",
		Password: "PasswordTralala",
		Gender:   model.UserGender(model.UserGender_value["FEMALE"]),
	}

	garage1 := model.Garage{
		Id:   "q001",
		Name: "Quel'thalas",
		Coordinate: &model.GarageCoordinate{
			Latitude:  45.123123123,
			Longitude: 54.1231313123,
		},
	}
	garage2 := model.Garage{
		Id:   "f001",
		Name: "Frostwing",
		Coordinate: &model.GarageCoordinate{
			Latitude:  32.123123123,
			Longitude: 11.1231313123,
		},
	}
	garage3 := model.Garage{
		Id:   "u001",
		Name: "Undercity",
		Coordinate: &model.GarageCoordinate{
			Latitude:  22.123123123,
			Longitude: 123.1231313123,
		},
	}

	user := serviceUser()

	fmt.Printf("\n %s> user test\n", strings.Repeat("=", 10))

	// register user1
	user.Register(context.Background(), &user1)

	// register user2
	user.Register(context.Background(), &user2)

	// show all registered users
	res1, err := user.List(context.Background(), new(emptypb.Empty))
	if err != nil {
		log.Fatal(err.Error())
	}
	res1String, _ := json.Marshal(res1.List)
	log.Println(string(res1String))

	garage := serviceGarage()

	fmt.Printf("\n %s> garage test A\n", strings.Repeat("=", 10))

	// add garage1 to user1
	garage.Add(context.Background(), &model.GarageAndUserId{
		UserId: user1.Id,
		Garage: &garage1,
	})

	// add garage2 to user1
	garage.Add(context.Background(), &model.GarageAndUserId{
		UserId: user1.Id,
		Garage: &garage2,
	})

	// show all garages of user1
	res2, err := garage.List(context.Background(), &model.GarageUserId{UserId: user1.Id})
	if err != nil {
		log.Fatal(err.Error())
	}
	res2String, _ := json.Marshal(res2.List)
	log.Println(string(res2String))

	fmt.Printf("\n %s> garage test B\n", strings.Repeat("=", 10))

	// add garage3 to user2
	garage.Add(context.Background(), &model.GarageAndUserId{
		UserId: user2.Id,
		Garage: &garage3,
	})

	// show all garages of user2
	res3, err := garage.List(context.Background(), &model.GarageUserId{UserId: user2.Id})
	if err != nil {
		log.Fatal(err.Error())
	}
	res3String, _ := json.Marshal(res3.List)
	log.Println(string(res3String))
}
