module github.com/prysmaticlabs/ethereumapis

go 1.13

require (
	github.com/ferranbt/fastssz v0.0.0-20210526181520-7df50c8568f8
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.0.1
	github.com/klauspost/cpuid/v2 v2.0.6 // indirect
	github.com/minio/sha256-simd v1.0.0 // indirect
	github.com/mitchellh/mapstructure v1.4.1 // indirect
	github.com/prysmaticlabs/eth2-types v0.0.0-20210303084904-c9735a06829d
	github.com/prysmaticlabs/go-bitfield v0.0.0-20210515192923-def021850363
	github.com/prysmaticlabs/protoc-gen-go-cast v0.0.0-20210505221644-3b823fdaca7f // indirect
	google.golang.org/genproto v0.0.0-20210426193834-eac7f76ac494
	google.golang.org/grpc v1.37.0
	google.golang.org/protobuf v1.26.0
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

// See: https://github.com/prysmaticlabs/grpc-gateway/issues/2
replace github.com/grpc-ecosystem/grpc-gateway/v2 => github.com/prysmaticlabs/grpc-gateway/v2 v2.3.1-0.20210524202756-cefc26c3f2bf
