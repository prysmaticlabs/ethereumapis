module github.com/prysmaticlabs/ethereumapis

go 1.13

require (
	github.com/ferranbt/fastssz v0.0.0-20210120143747-11b9eff30ea9
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.5.1
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.0.1
	github.com/mitchellh/mapstructure v1.4.0 // indirect
	github.com/prysmaticlabs/eth2-types v0.0.0-20210127031309-22cbe426eba6
	google.golang.org/genproto v0.0.0-20210325141258-5636347f2b14
	google.golang.org/grpc v1.36.0
	google.golang.org/protobuf v1.26.0
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace github.com/grpc-ecosystem/grpc-gateway/v2 => github.com/prysmaticlabs/grpc-gateway/v2 v2.3.1-0.20210330070718-b57834d108b1
