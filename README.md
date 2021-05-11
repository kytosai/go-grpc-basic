# Go gRPC basic

Video: https://www.youtube.com/watch?v=UScWGktlIng&list=PLC4c48H3oDRwlqUfUYfjdWH-d2uEanRpr&index=2 
Source: https://github.com/nkchuong1607/grpc_course

## Section 02 - Setup môi trường, generate code - The Funzy Dev 

Slider: https://docs.google.com/presentation/d/19zCPlujW2NIvUzG5NZFKqCIJmfyINoTasNtpsFSiQKo/edit#slide=id.p 

Install protobuf on MacOS 

```sh
brew install protobuf
```

**LƯU Ý** version pkg go được sử dụng trong video đã deprecated

Install pkg go
- https://github.com/grpc/grpc-go
- https://github.com/protocolbuffers/protobuf-go
  - https://pkg.go.dev/google.golang.org/protobuf/proto
    - Package proto provides functions operating on protobuf messages such as cloning, merging, and checking equality, as well as binary serialization and text serialization.
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
  - https://pkg.go.dev/google.golang.org/protobuf/cmd/protoc-gen-go
    - The protoc-gen-go binary is a protoc plugin to generate Go code for both proto2 and proto3 versions of the protocol buffer language.

```sh
go get -u google.golang.org/grpc

# Doc: https://grpc.io/docs/languages/go/quickstart/#prerequisites 
go get google.golang.org/protobuf/cmd/protoc-gen-go
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

Câu lệnh generate đổi thành

```sh
protoc calculator/calculatorpb/calculator.proto --go-grpc_out=.
```