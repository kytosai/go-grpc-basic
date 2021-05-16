# Lưu ý: `Makefile` phải sử dụng `tab` không được sử dụng `space`
# để tạo các tab thụt lùi code
gen-proto:	
	protoc \
		--go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
    --grpc-gateway_out ./ \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
    demopb/demopb.proto
run-server:
	go run main.go
run-proxy:
	go run proxy/proxy.go