# Run `make gen-cal` don't work
gen-cal:
  protoc calculator/calculatorpb/calculator.proto --go-grpc_out=.