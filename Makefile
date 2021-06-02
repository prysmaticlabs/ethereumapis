protos:
	protoc -Iproto --grpc-gateway_out=logtostderr=true,paths=source_relative:proto\
	  proto/faucet/faucet.proto
	protoc -Iproto --go_out=proto --go_opt=paths=source_relative \
    --go-grpc_out=proto --go-grpc_opt=paths=source_relative \
    proto/faucet/faucet.proto