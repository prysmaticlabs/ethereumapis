# Ethereum APIs

[![Build status](https://badge.buildkite.com/62be08099e9e228b165c2dba69c637eb9ca7a1ca95efd54b9f.svg?branch=master)](https://buildkite.com/prysmatic-labs/ethereum-apis)
[![Discord](https://user-images.githubusercontent.com/7288322/34471967-1df7808a-efbb-11e7-9088-ed0b04151291.png)](https://discord.gg/prysmaticlabs)
[![PyPI](https://img.shields.io/pypi/v/ethereumapis.svg)](https://pypi.org/project/ethereumapis/)


This repository hosts a **mirror** of [Prysm](https://github.com/prysmaticlabs/prysm/)'s service interface definitions for the Ethereum. These [protocol buffer](https://developers.google.com/protocol-buffers/) service definitions support [gRPC](https://grpc.io/) as well as JSON over HTTP.

## Mirror details

This repository is a **mirror** of [github.com/prysmaticlabs/prysm/proto/eth](https://github.com/prysmaticlabs/prysm/tree/develop/proto/eth). This means changes and contributions to protobufs **ARE NOT** accepted in this repository and must instead be made in the source repository instead. This project is updated on every new release of Prysm.

### Need assistance?

For general information on the functionality of gRPC and protocol buffers, see the [gRPC guide](https://grpc.io/docs/guides/). If you still have questions, feel free to stop by either our [Discord](https://discord.gg/prysmaticlabs) and a member of the team or our community will be happy to assist you.

## Service definitions

| Package | Service | Version | Description |
| :--- | :--- | :--- | :--- |
| eth | [BeaconChain](https://github.com/prysmaticlabs/ethereumapis/blob/master/eth/v1alpha1/beacon_chain.proto#L36) | v1alpha1 | This service is used to retrieve critical data relevant to the Ethereum proof-of-stake beacon chain, including the most recent head block, current pending deposits, the chain state and more. |
| eth | [Node](https://github.com/prysmaticlabs/ethereumapis/blob/master/eth/v1alpha1/node.proto#L33) | v1alpha1 | The Node service returns information about the Ethereum beacon node itself, including versioning and general information as well as network sync status and a list of services currently implemented on the beacon node. |
| eth | [Validator](https://github.com/prysmaticlabs/ethereumapis/blob/master/eth/v1alpha1/validator.proto) | v1alpha1 | This API provides the information a validator needs to retrieve throughout its lifecycle, including assignments from the network, its current index in the state as well as the rewards and penalties that have been applied to it. |

## Contributing


### Dependencies

Here's what you need to get started:

- A modern, UNIX operating system
- Go installed
- The latest release of [protocol buffers](https://grpc.io/docs/protoc-installation/) installed (the protoc compiler)
- The `cmake` package installed

#### Installing and Building Go Stubs

To install the required dependencies for compiling the protocol buffers, run
```text
make install
```

Next, to generate the Go stubs for the protocol buffers, run
```text
make generate
```

#### Building Python stubs

Python libraries can be generated using [`scripts/build-py-package.py`](https://github.com/prysmaticlabs/ethereumapis/blob/master/scripts/build-py-package.py).

### Making API Schema Changes

Say you want to add a new endpoint to the `BeaconChain` gRPC service in our API schema to retrieve orphaned blocks. Keep in mind making strict changes to the API schema can often times be difficult without a significant reason as this API is used by many different developers building on eth2. If you are confident in your desired changes, you can proceed by modifying the protobuf schema:

```go
service BeaconChain {
    // Retrieve orphaned blocks from the eth2 chain.
    rpc GetOrphanedBlocks(OrphanedBlocksRequest) returns (OrphanedBlocksResponse) {
        option (google.api.http) = {
            get: "/eth/v1alpha1/beacon/blocks/orphaned"
        };
    }
    ...
}

message OrphanedBlocksRequest {
    uint64 slot = 1;
}

message OrphanedBlocksResponse {
    repeated BeaconBlock blocks = 1;
}
```

After making your changes, you can regenerate the Go libraries from the schema by running:

```bash
make generate
```

## RESTful endpoints \(gRPC Transcoding\)

All of the gRPC services should define JSON over HTTP endpoints by specifying [HTTPRules](https://github.com/googleapis/googleapis/blob/master/google/api/http.proto). Developers may choose to bundle a REST service of gRPC with their client implementation binaries, or alternatively, they may use a JSON encoding proxy such as [Envoy Proxy](https://www.envoyproxy.io/), [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway), etc.

For more information on gRPC transcoding, see the examples found [here](https://github.com/googleapis/googleapis/blob/master/google/api/http.proto#L45).

Code sample:

```text
service Messaging {
    rpc GetMessage(GetMessageRequest) returns (Message) {
        option (google.api.http) = {
            get: "/v1/{name=messages/*}"
        };
    }
}
message GetMessageRequest {
    string name = 1; // Mapped to URL Path.
}
message Message {
    string text = 1; // The resource content.
}
```

This enables an HTTP REST to gRPC mapping, as shown below:

| HTTP | gRPC |
| :--- | :--- |
| `GET /v1/messages/123456` | `GetMessage(name: "messages/123456")` |

### gRPC tooling and resources

* [Awesome gRPC](https://github.com/grpc-ecosystem/awesome-grpc)
* [Google's API Style Guide](https://cloud.google.com/apis/design/)
* [Language reference for protoc3](https://developers.google.com/protocol-buffers/docs/proto3)
* [Protocol Buffer Basics: Go](https://developers.google.com/protocol-buffers/docs/gotutorial)
* [Transcoding gRPC to JSON/HTTP using Envoy](https://blog.jdriven.com/2018/11/transcoding-grpc-to-http-json-using-envoy/)

## License
[Apache 2.0](https://github.com/prysmaticlabs/ethereumapis/blob/master/LICENSE)
