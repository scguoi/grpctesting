package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/scguoi/grpctesting/example"
	"github.com/scguoi/grpctesting/impl"
)

func main() {
	serviceAddr := ":18000"
	service := grpc.NewServer()

	example.RegisterH3WrapperServer(service, &impl.Handler{})
	reflection.Register(service)

	log.Printf("service starting.")

	lis, e := net.Listen("tcp", serviceAddr)
	if e != nil {
		log.Fatalf("listen socket failed %v.", e)
	}

	go func() {
		e := service.Serve(lis)
		if e != nil {
			log.Fatalf("start its GRPC service failed %v.", e)
		}
	}()

	select {}
}
