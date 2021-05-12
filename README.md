# Go gRPC basic

Video: https://www.youtube.com/watch?v=UScWGktlIng&list=PLC4c48H3oDRwlqUfUYfjdWH-d2uEanRpr&index=2 
Source: https://github.com/nkchuong1607/grpc_course

## Section 02 - Setup môi trường, generate code - The Funzy Dev 

### Install

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

### Syntax


## Section 03.1 - Setup server

Video: https://www.youtube.com/watch?v=HMTz0-qjpgk&list=PLC4c48H3oDRwlqUfUYfjdWH-d2uEanRpr&index=3 

Code generate từ protobuf đã khác so với video 

## Section 03.2 - Client

gRPC recommend chúng ta khi connect tới nhau nên sử dụng SSL để đảm bảo an toàn -> tuy nhiên trong video demo chưa cần dùng nên khi connect tới server ta dùng `grpc.WithInsecure()` -> không an toàn

**LƯU Ý** khi viết `Makefile` cần sử dụng định dạng `tab` để thụt lùi code thay vì `space` nếu không sẽ gây bug không sử dụng được lệnh `make xxx`

## Section 04 - Unary API  

**LƯU Ý VỀ PHẦN GENERATE FILE PROTO**
- Theo chuẩn mới của generate code golang cho grpc sẽ có 2 file
  - 1 file chứa type, struct,... (file `xxx.pb.go`)
  - 1 file chứa client and server code (file `xxx_grpc.pb.go`)
- Doc gốc: https://grpc.io/docs/languages/go/quickstart/#regenerate-grpc-code

