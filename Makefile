gen-cal:
  protoc ./calculator/calculatorpb/calculator.proto --go-grpc_out=. 