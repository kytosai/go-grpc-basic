package main

import (
	"context"
	"gogrpcbasic/calculator/calculatorpb"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	callSumWithDeadline(client)
	// callPND(client)
	// callAverage(client)
	// callFindMax(client)
	// callSquare(client)

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

// # section 09 - deadline context
func callSumWithDeadline(c calculatorpb.CalculatorServiceClient) {
	log.Println("callSumWithDeadline() called...")

	// set time max request
	timeDeadline := time.Second * 4

	ctx, cancel := context.WithTimeout(context.TODO(), timeDeadline)
	defer cancel() // phải cancel để tránh tốn resource

	resp, err := c.SumWithDeadline(ctx, &calculatorpb.SumRequest{
		Num1: 1,
		Num2: 2,
	})

	if err != nil {
		if statusErr, ok := status.FromError(err); ok {
			if statusErr.Code() == codes.DeadlineExceeded { // hết thời gian deadline
				log.Println("calling callSumWithDeadline is DeadlineExceeded")
			} else {
				log.Printf("calling callSumWithDeadline err %v", err)
			}
		} else {
			log.Fatalf("calling callSumWithDeadline unknown err %v", err)
		}

		return
	}

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

func callAverage(c calculatorpb.CalculatorServiceClient) {
	log.Println("calling average api...")

	stream, err := c.Average(context.TODO())
	if err != nil {
		log.Fatalf("call average err %v", err)
	}

	listReq := []calculatorpb.AverageRequest{
		{
			Num: 5,
		},
		{
			Num: 10,
		},
		{
			Num: 12,
		},
		{
			Num: 3,
		},
		{
			Num: 4.2,
		},
	}

	// Send request to server
	for _, req := range listReq {
		err := stream.Send(&req)
		if err != nil {
			log.Fatalf("send average request err %v", err)
		}
		time.Sleep(time.Second)
	}

	// Close connect and receive result
	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("receive average response err %v", err)
	}

	log.Printf("average result %v", resp.GetResult())
}

func callFindMax(c calculatorpb.CalculatorServiceClient) {
	log.Println("calling FindMax()...")

	stream, err := c.FindMax(context.TODO())
	if err != nil {
		log.Fatalf("call FindMax() err %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		// Send multi request
		listReq := []calculatorpb.FindMaxRequest{
			{
				Num: 5,
			},
			{
				Num: 10,
			},
			{
				Num: 12,
			},
			{
				Num: 3,
			},
			{
				Num: 4,
			},
		}

		for _, req := range listReq {
			err := stream.Send(&req)
			if err != nil {
				log.Fatalf("send FindMax() request err %v", err)
			}
			time.Sleep(time.Second)
		}

		stream.CloseSend()
	}()

	go func() {
		// Receive multi request
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				log.Println("ending FindMax() api!")
				break
			}

			if err != nil {
				log.Fatalf("recv find max request err %v", err)
				break
			}

			log.Printf("max: %v", resp.GetMax())
		}

		close(waitc)
	}()

	// neo chờ để tránh chương trình gọi 2 goroutine trên
	// xong lại stop ngay lập tức
	<-waitc
}

// section 08 - handle error
func callSquare(c calculatorpb.CalculatorServiceClient) {
	num := int32(-1) // Thử với số <= 0 để thử show lỗi

	resp, err := c.Square(context.TODO(), &calculatorpb.SquareRequest{
		Num: num,
	})

	if err != nil {
		errStatus, ok := status.FromError(err)
		if ok {
			log.Printf("err msg: %v\n", errStatus.Message())
			log.Printf("err code: %v\n", errStatus.Code())

			if errStatus.Code() == codes.InvalidArgument {
				log.Printf("InvalidArgument num %v\n", num)
				return
			}
		}
	}

	log.Printf("square result: %v", resp.GetSquareRoot())
}
