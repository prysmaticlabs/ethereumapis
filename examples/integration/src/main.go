package main

import (
	"context"
	"log"
	"os"
	"strings"

	pb "github.com/prysmaticlabs/ethereumapis/eth/v1alpha1"
	grpc "google.golang.org/grpc"
)

func main() {
	// try and connect
	url := getRPCProvider()
	conn, err := grpc.Dial(url, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect to remote wallet: %v", err)
	} else {
		log.Printf("connected to remote wallet at: %v", url)
	}

	rpc := pb.NewBeaconChainClient(conn)

	req := pb.ListBlocksRequest{
		QueryFilter: &pb.ListBlocksRequest_Slot{Slot: 123},
	}
	res, error := rpc.ListBlocks(context.Background(), &req)
	if error != nil {
		log.Fatalf("could not fetch blocks: %v", error)
	} else {
		log.Printf("recieved blocks: %v", res.BlockContainers)
	}
}

func getRPCProvider() string {
	argsWithoutProg := os.Args[1:]

	return strings.Split(argsWithoutProg[0], "=")[1]
}
