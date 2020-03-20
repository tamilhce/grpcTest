/*
Title : CLI client
Description: CLI client for the Compute Microservice
Author: TAMILHCE
version : 1
*/
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/tamilhce/grpcTest/proto"
	"google.golang.org/grpc"
)

func main() {

	//command line args
	port := flag.String("port", "5000", "listening port")
	address := flag.String("address", "localhost", " grpc  server address")
	operation := flag.String("ops", "sum", "compute operation : sum/multiply")
	fmt.Printf("****Server Address : %s & port %s***\n", *address, *port)
	flag.Parse()
	//Parsing the values
	values := flag.Args()
	if len(values) < 2 {
		fmt.Println("Kindly pass 2 Args for compute eg : cliclient --address localhost --ops sum 10 20")
		os.Exit(1)
	}
	a, err := strconv.ParseInt(values[0], 10, 64)
	b, err := strconv.ParseInt(values[1], 10, 64)
	// To open insecure grpc tcp socket
	tcpsocket := *address + ":" + *port
	conn, err := grpc.Dial(tcpsocket, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	//client stub
	client := proto.NewComputeServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//contacting the sever
	req := &proto.Request{A: a, B: b}
	if *operation == "sum" {
		r, err := client.ComputeAdd(ctx, req)
		if err != nil {
			log.Fatalf("Attempt to reach ComputeAdd failed: %v", err)
		}
		//printing the response
		fmt.Printf("The sum of %v + %v is :%v\n", a, b, r.Result)
	} else if *operation == "multiply" {
		r, err := client.ComputeMultiply(ctx, req)
		if err != nil {
			log.Fatalf("Attempt to reach ComputeMultiply failed: %v", err)
		}
		//printing the response
		fmt.Printf("The  multply of %v * %v :%v\n", a, b, r.Result)

	}
}
