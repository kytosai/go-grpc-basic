# Go gRPC basic

- Video: https://www.youtube.com/watch?v=UScWGktlIng&list=PLC4c48H3oDRwlqUfUYfjdWH-d2uEanRpr&index=2 
- Source: https://github.com/nkchuong1607/grpc_course

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

## Section 03.1 - Setup server

Video: https://www.youtube.com/watch?v=HMTz0-qjpgk&list=PLC4c48H3oDRwlqUfUYfjdWH-d2uEanRpr&index=3 

Code generate từ protobuf đã khác so với video 

## Section 03.2 - Client

gRPC recommend chúng ta khi connect tới nhau nên sử dụng SSL để đảm bảo an toàn -> tuy nhiên trong video demo chưa cần dùng nên khi connect tới server ta dùng `grpc.WithInsecure()` -> không an toàn

**LƯU Ý** khi viết `Makefile` cần sử dụng định dạng `tab` để thụt lùi code thay vì `space` nếu không sẽ gây bug không sử dụng được lệnh `make xxx`

## Section 04 - Unary API  

- Unary api: 1 request từ client lên sẽ có 1 response trả về từ server, tương tự như HTTP request 

**LƯU Ý VỀ PHẦN GENERATE FILE PROTO**
- Theo chuẩn mới của generate code golang cho grpc sẽ có 2 file
  - 1 file chứa type, struct,... (file `xxx.pb.go`)
  - 1 file chứa client and server code (file `xxx_grpc.pb.go`)
- Doc gốc: https://grpc.io/docs/languages/go/quickstart/#regenerate-grpc-code

## Section 05 - Server Streaming API 

- Slider: https://docs.google.com/presentation/d/1QG0hmkzQDRzeNgE0AYRewoaEiL2wbNLAfSLugDqpIzo/edit#slide=id.p 

- Server streaming api: có nghĩa là client gửi lên 1 request và nhận về nhiều response 
  - VD: giống nhu news feed của facebook, cứ đẩy thông tin về và hiển thị,...

- Video: https://www.youtube.com/watch?v=bJA99UJPZnk&list=PLC4c48H3oDRwlqUfUYfjdWH-d2uEanRpr&index=6 

### Đề bài

**Mô tả**
- Làm về bài toán `prime number decomposition` -> nhập vào 1 con số, tìm 1 dãy số nguyên tố nhân lại với nhau tạo ra được số mình nhập
- VD: input là số 120. Kết quả trả ra là danh sách các số `2 * 2 * 2 * 3 * 5` là đúng

**Phân tích:**
- Số nguyên tố (prime number): là số tự nhiên lớn hơn 1 không phải là tích của hai số tự nhiên nhỏ hơn. Nói cách khác, số nguyên tố là những số chỉ có đúng hai ước số là 1 và chính nó. Các số tự nhiên lớn hơn 1 không phải là số nguyên tố được gọi là hợp số
- Ước số: Mô tả rõ hơn thì khi một số tự nhiên A được gọi là ước số của số tự nhiên B nếu B chia hết cho A.
  - Ví dụ: 6 chia hết được cho [1,2,3,6], thì [1,2,3,6] được gọi là ước số của 6.

## Section 06 - Client Streaming API 

**Đề bài** thực hiện tính trung bình cho 1 dãy số
VD: input vào là [5,10,12,3,4.2] => kết quả là 6.48

## section 07 - Bi-Directional Streaming API

**Đề bài** tìm con số lớn nhất, client sẽ gửi lên dãy số và server trả về số lớn nhất trong dãy số đó
  - VD: client bắn liên tục dãy số [5,10,12,3,4] -> thì phía server tính toán sau mỗi lần nhận message từ client sẽ trả ra là [5,10,12,12,12]

## section 08 - Handle Error (Xử lý lỗi)

Video: https://www.youtube.com/watch?v=btsU8EtP2dY&list=PLC4c48H3oDRwlqUfUYfjdWH-d2uEanRpr&index=9 
  - Video này xử lý lỗi chi tiết các lỗi xảy ra ở các video trước khi ta đều chỉ print ra là xong

- Các status code trong grpc: https://github.com/grpc/grpc/blob/master/doc/statuscodes.md 
- Repo sample code handle err: https://github.com/avinassh/grpc-errors 
- Video này ta sẽ làm hàm `Square` (căn bậc hai)
- Sử dụng 2 pkg này để trả lỗi

```txt
"google.golang.org/grpc/codes"
"google.golang.org/grpc/status"
```

## Section 09 - Deadline context

Video: https://www.youtube.com/watch?v=IhNd11EiXUk&list=PLC4c48H3oDRwlqUfUYfjdWH-d2uEanRpr&index=10
  - Cách sử dụng context deadline để xử lý request timeout trong gRPC 
