package main

import (
	"fmt"
	"grpc-demo/server/stockpb"
	"log"
	"net"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) StockTracker(req *stockpb.StockRequest,
	stream stockpb.Stock_StockTrackerServer) error {

	return nil
}

var (
	port int = 8000
)

func main() {
	addr := fmt.Sprintf("0.0.0.0:%d", port)

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Error while listening: %v", err)
	}

	s := grpc.NewServer()
	stockpb.RegisterStockServer(s, &server{})

	log.Printf("Starting server on port: %d\n", port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error while serving: %v", err)
	}

}