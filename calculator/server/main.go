package main

import (
	"fmt"
	"log"
	"net"

	"gogrpcbasic/calculator/calculatorpb"

	"google.golang.org/grpc"
)

type server struct {
	// ? khúc này đang xử lý khác video:
	// ? vì probuf code gen ra đã có impl thêm func
	calculatorpb.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50069")
	if err != nil {
		log.Fatalf("err while create listen %v", err)
	}

	s := grpc.NewServer()

	calculatorpb.RegisterCalculatorServiceServer(
		s,
		&server{},
	)

	fmt.Println("calculator is running...")

	err = s.Serve(lis)

	if err != nil {
		log.Fatalf("err while serve %v", err)
	}
}
