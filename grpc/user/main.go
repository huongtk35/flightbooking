package main

import (
	"flag"
	"flightbooking/db"
	"flightbooking/grpc/user/handlers"
	repositories "flightbooking/grpc/user/repositpries"
	"flightbooking/proto"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

var (
	configFile = flag.String("config-file", "config.yml", "Location of config file")
	port       = flag.Int("port", 2222, "Port of grpc")
)

func main() {
	err := db.AutoBindConfig(*configFile)
	if err != nil {
		panic(err)
	}
	grpcAddr := fmt.Sprintf("0.0.0.0:%v", db.GetUserGrpcPort())
	// init grpc
	c := make(chan bool)
	userService := repositories.NewUserService()
	go func() {
		list, err := net.Listen("tcp", grpcAddr)
		if err != nil {
			log.Fatalf("Failed to start listener %v", err)
		}
		s := grpc.NewServer()
		server := handlers.NewServer(&userService)
		proto.RegisterUserServiceServer(s, server)
		log.Printf("Listening grpc on %v\n", grpcAddr)
		if err = s.Serve(list); err != nil {
			c <- false
			log.Fatalf("Failed to serve %v\n", err)
		}
	}()
	select {
	case success := <-c:
		if !success {
			panic("Cannot init grpc")
		}
	case _ = <-time.After(3 * time.Second):
		log.Println("Serving grpc for user-service...")
	}
}
