package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"log"
	"net"

	"chapter-c30/common/config"
	"chapter-c30/common/model"

	"google.golang.org/grpc"
)

var localStorage *model.GarageListByUser

func init() {
	localStorage = new(model.GarageListByUser)
	localStorage.List = make(map[string]*model.GarageList)
}

type GaragesServer struct {
	model.UnimplementedGaragesServer
}

func (GaragesServer) Add(_ context.Context, param *model.GarageAndUserId) (*empty.Empty, error) {
	userId := param.UserId
	garage := param.Garage

	if _, ok := localStorage.List[userId]; !ok {
		localStorage.List[userId] = new(model.GarageList)
		localStorage.List[userId].List = make([]*model.Garage, 0)
	}
	localStorage.List[userId].List = append(localStorage.List[userId].List, garage)

	log.Println("Adding garage", garage.String(), "for user", userId)

	return new(empty.Empty), nil
}

func (GaragesServer) List(_ context.Context, param *model.GarageUserId) (*model.GarageList, error) {
	userId := param.UserId

	return localStorage.List[userId], nil
}

func main() {
	srv := grpc.NewServer()
	var garageSrv GaragesServer
	model.RegisterGaragesServer(srv, garageSrv)

	log.Println("Starting RPC server at", config.ServiceGaragePort)

	l, err := net.Listen("tcp", config.ServiceGaragePort)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", config.ServiceGaragePort, err)
	}

	log.Fatal(srv.Serve(l))
}
