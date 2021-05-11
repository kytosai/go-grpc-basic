# Lưu ý: `Makefile` phải sử dụng `tab` không được sử dụng `space`
# để tạo các tab thụt lùi code
gen-cal:
	protoc calculator/calculatorpb/calculator.proto --go-grpc_out=.
run-server:
	go run calculator/server/server.go
run-client:
	go run calculator/client/client.go
