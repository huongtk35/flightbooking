package main

import (
	"flag"
	"flightbooking/db"
	"flightbooking/grpc/flight/handlers"
	"flightbooking/grpc/flight/repositories"
	"flightbooking/intercepter"
	"flightbooking/proto"
	"fmt"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

	listen, err := net.Listen("tcp", fmt.Sprintf(":%v", *port))
	if err != nil {
		panic(err)
	}

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			intercepter.UnaryServerLoggingIntercepter(logger),
		)),
	)

	flightService := repositories.NewFlightService()
	if err != nil {
		panic(err)
	}
	h := &handlers.Server{FlightService: &flightService}

	if err != nil {
		panic(err)
	}
	reflection.Register(s)
	proto.RegisterFlightServiceServer(s, h)

	fmt.Printf("Listen at port: %v\n", *port)

	s.Serve(listen)
}
