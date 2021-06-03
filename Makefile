generate:
	protoc -I . --go-cast_out=. --go-cast_opt paths=source_relative --experimental_allow_proto3_optional \
		--go-grpc_out=. --go-grpc_opt paths=source_relative \
		$(find ./eth -name "*.proto")
	protoc -I . --grpc-gateway_out=. --grpc-gateway_opt=logtostderr=true,paths=source_relative \
		--experimental_allow_proto3_optional \
		$(find ./eth -name "*.proto")
install:
	go install \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc \
		github.com/prysmaticlabs/grpc-gateway/v2/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
		github.com/prysmaticlabs/protoc-gen-go-cast