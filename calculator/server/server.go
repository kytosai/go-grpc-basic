package main

import (
	"context"
	"log"
	"net"
	"time"

	"gogrpcbasic/calculator/calculatorpb"

	"google.golang.org/grpc"
)

type server struct {
	// ! khúc này đang xử lý khác video:
	// ! vì probuf code gen ra đã có impl thêm func
	// ! nên cần đưa các thông tin của interface vào đầy đủ
	calculatorpb.CalculatorServiceServer
}

func (*server) Sum(
	ctx context.Context,
	req *calculatorpb.SumRequest,
) (*calculatorpb.SumResponse, error) {
	resp := calculatorpb.SumResponse{
		Result: req.GetNum1() + req.GetNum2(),
	}

	return &resp, nil
}

func (*server) PrimeNumberDecomposition(
	req *calculatorpb.PNDRequest,
	stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer,
) error {
	log.Println("PrimeNumberDecomposition called...")

	k := int32(2)
	N := req.GetNumber()

	for N > 1 {
		if N%k == 0 {
			N = N / k

			// Send data to client
			stream.Send(&calculatorpb.PNDResponse{
				Result: k,
			})

			time.Sleep(time.Second)
		} else {
			k++
			log.Printf("k increase to %v", k)
		}
	}

	return nil
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

	log.Println("calculator server is running...")

	err = s.Serve(lis)

	if err != nil {
		log.Fatalf("err while serve %v", err)
	}
}
