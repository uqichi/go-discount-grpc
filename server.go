package main

import (
	"net"

	"log"

	"github.com/uqichi/go-discount-grpc/proto"
	"github.com/uqichi/go-discount-grpc/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const port string = ":50061"

func main() {
	m := service.NewDiscountService()

	listner, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	pb.RegisterDiscountServiceServer(server, m)
	reflection.Register(server)

	// run server
	log.Printf("gRPC server started on %s", port)
	err = server.Serve(listner)
	if err != nil {
		log.Fatal(err)
	}
}
