package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"

	"github.com/raducrisan1/microservice-stockinfo/stockinfo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	lis, err := net.Listen("tcp", ":3001")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	stockinfo.RegisterStockInfoServiceServer(s, &server{})
	//this is used to allow API inspection via grpc_cli tool
	reflection.Register(s)
	stop := ossignal()
	go func() {
		<-stop
		s.Stop()
	}()
	fmt.Println("Server started on port 3001")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	fmt.Println("\nThe server has stopped")
}

func (s *server) StockInfo(context.Context, *stockinfo.StockInfoRequest) (*stockinfo.StockInfoResponse, error) {
	rand.Seed(time.Now().UTC().UnixNano())

	rsp := stockinfo.StockInfoResponse{
		Stockname:  "NVDA",
		CciData:    randInt(5, -100, 100),
		RsiData:    randInt(5, 0, 100),
		MacdData:   randInt(5, -30, 30),
		VolumeData: randInt(5, 5000, 50000)}

	return &rsp, nil
}

func randInt(size int, min int, max int) []*stockinfo.Indicator {
	result := make([]*stockinfo.Indicator, size)
	timeRef, _ := time.Parse(time.RFC3339, "2018-11-11 09:30Z")
	for i := range result {
		res := new(stockinfo.Indicator)
		res.Value = float32(min + rand.Intn(max-min))
		d, _ := time.ParseDuration(strconv.Itoa(i+1) + "m")
		res.Timestamp = timeRef.Add(d).Unix()
		result[i] = res
	}
	return result
}
