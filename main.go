package main

import (
	"context"
	"log"
	"math/rand"
	"net"

	pb "github.com/abdoroot/currency/currency/pb"
	"google.golang.org/grpc"
)

type CurrencyConverte struct {
	pb.UnimplementedCurrencyConverterServer
}

func (c CurrencyConverte) GetRate(ctx context.Context, in *pb.CurrencyRequest) (*pb.CurrencyResponse, error) {
	return &pb.CurrencyResponse{
		Rate: float32(rand.Int()),
	}, nil
}

func main() {
	l, err := net.Listen("tcp", ":2026")
	if err != nil {
		log.Fatal("colud not start listener : " + err.Error())
	}
	gs := grpc.NewServer()
	cc := CurrencyConverte{}
	pb.RegisterCurrencyConverterServer(gs, cc)
	gs.Serve(l)
}
