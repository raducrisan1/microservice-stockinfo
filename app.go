package main

import (
	"context"
	"log"
	"net"

	"github.com/raducrisan1/microservice-stockinfo/stockinfo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	stockinfo.RegisterStockInfoServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *server) StockInfo(context.Context, *stockinfo.StockInfoRequest) (*stockinfo.StockInfoResponse, error) {
	rsp := stockinfo.StockInfoResponse{
		Stockname: "NVDA",
	}
	return &rsp, nil
}
