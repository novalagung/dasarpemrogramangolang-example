package main

import (
	"chapter-c29/model"
	"fmt"
	"google.golang.org/protobuf/encoding/protojson"
	"os"
)

var user1 = &model.User{
	Id:       "u001",
	Name:     "Sylvana Windrunner",
	Password: "f0r Th3 H0rD3",
	Gender:   model.UserGender_FEMALE,
}

var userList = &model.UserList{
	List: []*model.User{
		user1,
	},
}

var garage1 = &model.Garage{
	Id:   "g001",
	Name: "Kalimdor",
	Coordinate: &model.GarageCoordinate{
		Latitude:  23.2212847,
		Longitude: 53.22033123,
	},
}

var garageList = &model.GarageList{
	List: []*model.Garage{
		garage1,
	},
}

var garageListByUser = &model.GarageListByUser{
	List: map[string]*model.GarageList{
		user1.Id: garageList,
	},
}

func main() {
	// =========== original
	fmt.Printf("# ==== Original\n       %#v \n", user1)

	// =========== as string
	fmt.Printf("# ==== As String\n       %s \n", user1.String())

	// =========== as json string

	jsonb, err1 := protojson.Marshal(garageList)
	if err1 != nil {
		fmt.Println(err1.Error())
		os.Exit(0)
	}
	fmt.Printf("# ==== As JSON String\n       %s \n", string(jsonb))

	// =========== from json string to protobuf object
	protoObject := new(model.GarageList)
	err2 := protojson.Unmarshal(jsonb, protoObject)
	if err2 != nil {
		fmt.Println(err2.Error())
		os.Exit(0)
	}

	fmt.Printf("# ==== As String\n       %s \n", protoObject.String())
}
