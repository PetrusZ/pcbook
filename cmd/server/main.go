package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/PetrusZ/pcbook/pb"
	"github.com/PetrusZ/pcbook/service"
	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("port", 0, "the server port")
	flag.Parse()
	log.Printf("start server port %d", *port)

	laptopServer := service.NewLaptopServer(service.NewInMemoryLaptopStore())
	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	address := fmt.Sprintf("127.0.0.1:%d", *port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("can't start server: ", err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("can't start server: ", err)
	}
}
