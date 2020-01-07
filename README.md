# Ethereum APIs

[![Build status](https://badge.buildkite.com/62be08099e9e228b165c2dba69c637eb9ca7a1ca95efd54b9f.svg?branch=master)](https://buildkite.com/prysmatic-labs/ethereum-apis)
[![Discord](https://user-images.githubusercontent.com/7288322/34471967-1df7808a-efbb-11e7-9088-ed0b04151291.png)](https://discord.gg/KSA7rPr)
[![Gitter](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/prysmaticlabs/geth-sharding?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge)
[![ETH2.0_Spec_Version 0.8.1](https://img.shields.io/badge/ETH2.0%20Spec%20Version-v0.8.1-blue.svg)](https://github.com/ethereum/eth2.0-specs/commit/452ecf8e27c7852c7854597f2b1bb4a62b80c7ec)


This repository hosts [Prysm](https://github.com/prysmaticlabs/prysm/)'s service interface definitions for the Ethereum 2.0 API. These [protocol buffer](https://developers.google.com/protocol-buffers/) service definitions support [gRPC](https://grpc.io/) as well as JSON over HTTP.

### Need assistance?
More indepth descriptions of each service are available in [this section](https://prysmaticlabs.gitbook.io/prysm/how-prysm-works/ethereum-2.0-public-api) of our official documentation. For more general information on the functionality of gRPC and protocol buffers, see the [gRPC guide](https://grpc.io/docs/guides/). If you still have questions, feel free to stop by either our [Discord](https://discord.gg/KSA7rPr) or [Gitter](https://gitter.im/prysmaticlabs/geth-sharding?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge) and a member of the team or our community will be happy to assist you.

## Services

| Package | Service | Version | Description |
|---------|---------|---------|-------------|
| eth | [BeaconChain](eth/v1alpha1/beacon_chain.proto#L36) | v1alpha1 | This service is used to retrieve critical data relevant to the Ethereum 2.0 phase 0 beacon chain, including the most recent head block, current pending deposits, the chain state and more. |
| eth | [Node](eth/v1alpha1/node.proto#L33) | v1alpha1 | The Node service returns information about the Ethereum node itself, including versioning and general information as well as network sync status and a list of services currently implemented on the node.
| eth | [Validator](eth/v1alpha1/validator.proto) | v1alpha1 | This API provides the information a validator needs to retrieve throughout its lifecycle, including recieved assignments from the network, its current index in the state, as well the rewards and penalties that have been applied to it.


## Generating client libraries
Client libraries may be generated using [protoc](https://github.com/protocolbuffers/protobuf), a language-neutral tool for serializing structured data. Below are instructions for performing an installation and compiling a `.proto` file.

### Dependencies
- The latest release of [protobuf compiler (protoc)](https://github.com/protocolbuffers/protobuf/releases/tag/v3.9.0)

### Compiling with protoc

1. First, ensure you have the most recent version of `protoc` installed by issuing the command:
```
protoc --version
```

2. To install the Go protocol buffers plugin, issue the command:
```
go get -u github.com/golang/protobuf/protoc-gen-go
```
  The compiler plugin `protoc-gen-go` will be installed in `$GOBIN`, defaulting to `$GOPATH/bin`. It must be in your `$PATH` for `protoc` to locate it.

3. The compiler can now be run. Note that this command requires a few parameters; namely a source directly, a destination directory and a path to the `.proto` file itself. Example:
```
protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/node.proto
```
The above command will generate a ``node.pb.go`` file in your specified destination directory. 


## Additional information

### RESTful endpoints (gRPC Transcoding)

All of the gRPC services should define JSON over HTTP endpoints by specifying [HTTPRules](https://github.com/googleapis/googleapis/blob/master/google/api/http.proto) for such. Developers may choose to bundle a REST service of gRPC with their client implementation binaries, or alternatively, they may use a JSON encoding proxy such as [Envoy Proxy](https://www.envoyproxy.io/), [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway), etc.

For more information on gRPC transcoding, see the examples found in [google/api/http.proto](https://github.com/googleapis/googleapis/blob/master/google/api/http.proto#L45).

An example:
```proto
service Messaging {
    rpc GetMessage(GetMessageRequest) returns (Message) {
        option (google.api.http) = {
            get: "/v1/{name=messages/*}"
        };
    }
    message GetMessageRequest {
        string name = 1; // Mapped to URL Path.
    }
    message Message {
        string text = 1; // The resource content.
    }
}
```

This enables an HTTP REST to gRPC mapping, as shown below:

HTTP | gRPC
-----|-----
`GET /v1/messages/123456`  | `GetMessage(name: "messages/123456")`


### JSON mapping

The majority of field primative types for Ethereum are either `bytes` or `uint64`. The canonical JSON mapping for those fields are a Base64 encoded string for `bytes`, or a string representation of `uint64`. Since JavaScript loses precision for values over [MAX_SAFE_INTEGER](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Number/MAX_SAFE_INTEGER), uint64 must be a JSON string in order to accommodate those values. If the field value not changed and is still set to protobuf's default, this field will be omitted from the JSON encoding entirely. 

For more details on JSON mapping for other types, view the relevent section in the [proto3 language guide](https://developers.google.com/protocol-buffers/docs/proto3#json).

### Helpful gRPC tooling and resources

- [Awesome gRPC](https://github.com/grpc-ecosystem/awesome-grpc)
- [Google's API Style Guide](https://cloud.google.com/apis/design/)
- [Language reference for protoc3](https://developers.google.com/protocol-buffers/docs/proto3)
- [Protocol Buffer Basics: Go](https://developers.google.com/protocol-buffers/docs/gotutorial)
- [Transcoding gRPC to JSON/HTTP using Envoy](https://blog.jdriven.com/2018/11/transcoding-grpc-to-http-json-using-envoy/)


## Contributing
We have put all of our contribution guidelines into [CONTRIBUTING.md](https://github.com/prysmaticlabs/prysm/blob/master/CONTRIBUTING.md)! Check it out to get started.

## License
[Apache 2.0](https://github.com/prysmaticlabs/ethereumapis/blob/master/LICENSE)
