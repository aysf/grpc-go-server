package main

import (
	"log"
	"net"

	"grpc-go-server/cmd/services"
	"grpc-go-server/pb/productpb"

	"google.golang.org/grpc"
)

const (
	port = ":50011"
)

func main() {

	netListen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("cannot listen port:", err)
	}

	grpcServer := grpc.NewServer()
	productService := services.ProductService{}
	productpb.RegisterProductServiceServer(grpcServer, productService)

	log.Printf("server started at %v\n", netListen.Addr())
	if err := grpcServer.Serve(netListen); err != nil {
		log.Fatal("failed to server:", err)
	}
}
