# Run `make gen-cal` don't work
gen-cal:
  protoc calculator/calculatorpb/calculator.proto --go-grpc_out=.
run-server:
  go run calculator/server/main.go
