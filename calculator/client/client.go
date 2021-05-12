package main

import (
	"context"
	"gogrpcbasic/calculator/calculatorpb"
	"io"
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

	// Call api from server
	// callSum(client)
	callPND(client)

	// "%f"	decimal point but no exponent (số mũ), e.g. 123.456
	// log.Printf("service client %f", client)

	log.Println("close connect to server!")
}

func callSum(c calculatorpb.CalculatorServiceClient) {
	resp, _ := c.Sum(context.TODO(), &calculatorpb.SumRequest{
		Num1: 1,
		Num2: 2,
	})

	log.Println(resp.GetResult())
}

func callPND(c calculatorpb.CalculatorServiceClient) {
	stream, err := c.PrimeNumberDecomposition(context.TODO(), &calculatorpb.PNDRequest{
		Number: 120,
	})

	if err != nil {
		log.Fatalf("callPND err %v", err)
	}

	for {
		// Vòng for này không chạy liên tục vì khi chạy dòng code dưới
		// nó sẽ handle cho đến khi có kết quả trả về, kiểu giống như
		// đợi nhận kết quả từ channel vậy
		resp, recvErr := stream.Recv()

		// Kiểm tra xem server đã gửi kết thúc chưa
		if recvErr == io.EOF {
			log.Println("server finish streaming!")
			return
		}

		log.Printf("prime number %v", resp.GetResult())
	}
}
