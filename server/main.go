/*
Title : Simple GO GRPC SERVER
Description: GRPC SERVER TO Expose Compute Service
Author: TAMILHCE
version : 1
*/

package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/tamilhce/grpcTest/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	//Fetching the env variable grpcPort
	port, exists := os.LookupEnv("GrpcPort")
	if !exists {
		port = "5000"
	}
	port = ":" + port
	//create a tcp listner with port 5000
	listener, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	fmt.Println("The server listens on", port)
	//create a grpc server
	srv := grpc.NewServer()
	//Add the services to the server
	proto.RegisterComputeServiceServer(srv, &server{})
	/*reflection - assists clients at runtime to construct RPC
	requests and responses without precompiled service information.
	*/
	reflection.Register(srv)
	//Finally - grpc server listner
	if e := srv.Serve(listener); e != nil {
		panic(err)
	}
}

//Add function request inputs a b & operation. returns response
func (s *server) ComputeAdd(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()
	log.Println("Got the Request for sum of ", a, b)
	result := a + b
	return &proto.Response{Result: result}, nil
}

//Mulitpy  function request inputs a b & operation. returns response
func (s *server) ComputeMultiply(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()
	log.Println("Got the Request for Multiply of ", a, b)
	result := a * b
	return &proto.Response{Result: result}, nil
}
