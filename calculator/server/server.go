package main

import (
	"context"
	"io"
	"log"
	"math"
	"net"
	"time"

	"gogrpcbasic/calculator/calculatorpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
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

// Section 09 - deadline context
func (*server) SumWithDeadline(
	ctx context.Context,
	req *calculatorpb.SumRequest,
) (*calculatorpb.SumResponse, error) {
	log.Println("SumWithDeadline() called...")

	// Giả đoạn code là có nhiều tác vụ nào đó
	// xử lý quá nhiều chiếm nhiều thời gian trước
	// khi response về tới client
	for i := 0; i < 3; i++ {
		if ctx.Err() == context.Canceled {
			log.Println("context.Canceled...")
			return nil, status.Errorf(codes.Canceled, "client canceled request")
		}

		time.Sleep(1 * time.Second)
	}

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

func (*server) Average(stream calculatorpb.CalculatorService_AverageServer) error {
	log.Println("Average() called...")

	var total float32
	var count int

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			// Tính trung bình cộng và gửi về client
			// chỉ 1 lần duy nhất
			resp := calculatorpb.AverageResponse{
				Result: total / float32(count),
			}

			return stream.SendAndClose(&resp)
		}

		if err != nil {
			log.Fatalf("err while recv Average %v", err)
			return err
		}

		log.Printf("receive num %v", req.GetNum())
		total += req.GetNum()
		count++
	}
}

func (*server) FindMax(stream calculatorpb.CalculatorService_FindMaxServer) error {
	log.Println("FindMax() called")

	max := int32(0)
	for {
		req, err := stream.Recv()

		// client ngắt kết nối tới thì dừng không
		// cần xử lý đặc biệt gì thêm
		if err == io.EOF {
			log.Println("EOF")
			return nil
		}

		if err != nil {
			log.Fatalf("err while recv %v", err)
			return err
		}

		num := req.GetNum()
		log.Printf("recv num %v", num)
		if num > max {
			max = num
		}

		err = stream.Send(&calculatorpb.FindMaxResponse{
			Max: max,
		})

		if err != nil {
			log.Fatalf("send max err %v", err)
			return err
		}
	}
}

// section 08 - Handle Error
func (*server) Square(
	ctx context.Context,
	req *calculatorpb.SquareRequest,
) (*calculatorpb.SquareResponse, error) {
	log.Println("Square() called...")

	num := req.GetNum()
	if num < 0 {
		log.Printf("req num < 0, num=%v, return InvalidArgument", num)
		return nil, status.Errorf(codes.InvalidArgument, "Expect num > 0, req num war %v", num)
	}

	return &calculatorpb.SquareResponse{
		SquareRoot: math.Sqrt(float64(num)),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50069")
	if err != nil {
		log.Fatalf("err while create listen %v", err)
	}

	// Setup SSL for server
	certFile := "ssl/server.crt"
	keyFile := "ssl/server.pem"

	creds, sslErr := credentials.NewServerTLSFromFile(certFile, keyFile)
	if sslErr != nil {
		log.Fatalf("create cred ssl err %v\n", sslErr)
		return
	}
	opts := grpc.Creds(creds)

	// Create new server
	s := grpc.NewServer(opts)

	calculatorpb.RegisterCalculatorServiceServer(
		s,
		&server{},
	)

	log.Println("calculator server is running...")

	// Start server
	err = s.Serve(lis)

	if err != nil {
		log.Fatalf("err while serve %v", err)
	}
}
