package main

import (
	pb "/Users/anusharesu/Desktop/Nani_Bitbucket/Go/gRPC"
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dail("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatal("Didn't Connect")
	}
	defer conn.Close()

	c := pb.NewHelloServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

}
