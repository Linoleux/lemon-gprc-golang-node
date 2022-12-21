package main

import (
	"flag"
	"fmt"
	"grpc/pb"
	"grpc/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "the server port")
)

func main() {
	flag.Parse()
	log.Printf("server start on port %d", *port)
	address := fmt.Sprintf(":%d", *port)
	listener, err := net.Listen("tcp", address)
	fmt.Println(address)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

	laptopStore := service.NewInMemoryLaptopStore()
	laptopServer := service.NewLaptopServer(laptopStore)
	grpcServer := grpc.NewServer()

	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
