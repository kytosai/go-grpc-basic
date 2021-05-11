package main

import (
	"gogrpcbasic/calculator/calculatorpb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	clientConn, err := grpc.Dial("localhost:50069", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("err while dial %v", err)
	}
	defer clientConn.Close()

	client := calculatorpb.NewCalculatorServiceClient(clientConn)

	// "%f"	decimal point but no exponent (số mũ), e.g. 123.456
	log.Printf("service client %f", client)
}
