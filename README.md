# Ethereum APIs

[![Build status](https://badge.buildkite.com/62be08099e9e228b165c2dba69c637eb9ca7a1ca95efd54b9f.svg?branch=master)](https://buildkite.com/prysmatic-labs/ethereum-apis)
[![Discord](https://user-images.githubusercontent.com/7288322/34471967-1df7808a-efbb-11e7-9088-ed0b04151291.png)](https://discord.gg/KSA7rPr)
[![Gitter](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/prysmaticlabs/geth-sharding?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge)
[![ETH2.0_Spec_Version 0.12.0](https://img.shields.io/badge/ETH2.0%20Spec%20Version-v0.12.0-blue.svg)](https://github.com/ethereum/eth2.0-specs/releases/tag/v0.12.0)
[![PyPI](https://img.shields.io/pypi/v/ethereumapis.svg)](https://pypi.org/project/ethereumapis/)


This repository hosts [Prysm](https://github.com/prysmaticlabs/prysm/)'s service interface definitions for the Ethereum 2.0 API. These [protocol buffer](https://developers.google.com/protocol-buffers/) service definitions support [gRPC](https://grpc.io/) as well as JSON over HTTP.

### Need assistance?
More indepth descriptions of each service are available in [this section](https://prysmaticlabs.gitbook.io/prysm/how-prysm-works/ethereum-2.0-public-api) of our official documentation. For more general information on the functionality of gRPC and protocol buffers, see the [gRPC guide](https://grpc.io/docs/guides/). If you still have questions, feel free to stop by either our [Discord](https://discord.gg/KSA7rPr) or [Gitter](https://gitter.im/prysmaticlabs/geth-sharding?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge) and a member of the team or our community will be happy to assist you.

## Service definitions

| Package | Service | Version | Description |
| :--- | :--- | :--- | :--- |
| eth | [BeaconChain](https://github.com/prysmaticlabs/ethereumapis/blob/master/eth/v1alpha1/beacon_chain.proto#L36) | v1alpha1 | This service is used to retrieve critical data relevant to the Ethereum 2.0 phase 0 beacon chain, including the most recent head block, current pending deposits, the chain state and more. |
| eth | [Node](https://github.com/prysmaticlabs/ethereumapis/blob/master/eth/v1alpha1/node.proto#L33) | v1alpha1 | The Node service returns information about the Ethereum node itself, including versioning and general information as well as network sync status and a list of services currently implemented on the node. |
| eth | [Validator](https://github.com/prysmaticlabs/ethereumapis/blob/master/eth/v1alpha1/validator.proto) | v1alpha1 | This API provides the information a validator needs to retrieve throughout its life cycle, including recieved assignments from the network, its current index in the state as well as the rewards and penalties that have been applied to it. |

## Contributing

Thanks for wanting to contribute to our eth2 API! Go libraries may be generated from the [ethereumapis repository](https://github.com/prysmaticlabs/ethereumapis) using [Bazel](https://bazel.build), making it easy to make changes to the schemas needed and generate Go files from them.

Python libraries can be generated using [`scripts/build-py-package.py`](https://github.com/prysmaticlabs/ethereumapis/blob/master/scripts/build-py-package.py); we regularly publish these libraries to PyPI as [ethereumapis](https://pypi.org/project/ethereumapis/).

### Dependencies

Here's what you need to get started:

- A modern, UNIX operating system
- The latest release of [Bazel](https://docs.bazel.build/versions/master/install.html) installed
- The `cmake` package installed
- The `git` package installed

### Making API Schema Changes

Say you want to add a new endpoint to the `BeaconChain` gRPC service in our API schema to retrieve orphaned blocks. First, make sure the functionality you wish to add is not already covered by one of our endpoints on https://api.prylabs.network. Also, keep in mind making strict changes to the API schema can often times be difficult without a significant reason as this API is used by many different developers building on eth2. If you are confident in your desired changes, you can proceed by modifying the protobuf schema:

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
$ ./scripts/update-go-pbs.sh
```

Then, open a pull request with your changes on https://github.com/prysmaticlabs/ethereumapis. Next, you'll be ready to implement your new changes in Prysm itself.

### Implementing Your Changes in Prysm

Ensure you have read our [contribution guidelines](https://docs.prylabs.network/docs/contribute/contribution-guidelines/) first. Then, once your changes to the API schema are merged into the master branch of ethereumapis, you can update Prysm's dependency on ethereumapis to its latest version with the command:

```bash
$ bazel run //:gazelle -- update-repos github.com/prysmaticlabs/ethereumapis
```

Prysm also utilizes generated mocks for testing gRPC requests/responses, so you will also need to regenerate the required mocks by running:

```bash
$ ./scripts/update-mockgen.sh
```

Now, you will be able to implement your required changes in Prysm.

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

### JSON mapping

The majority of field primitive types for Ethereum are either `bytes` or `uint64`. The canonical JSON mapping for those fields are a Base64 encoded string for `bytes`, or a string representation of `uint64`. Since JavaScript loses precision for values over [MAX\_SAFE\_INTEGER](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Number/MAX_SAFE_INTEGER), uint64 must be a JSON string in order to accommodate those values. If the field value not changed and is still set to protobuf's default, this field will be omitted from the JSON encoding entirely.

For more details on JSON mapping for other types, view the relevant section in the [proto3 language guide](https://developers.google.com/protocol-buffers/docs/proto3#json).

### gRPC tooling and resources

* [Awesome gRPC](https://github.com/grpc-ecosystem/awesome-grpc)
* [Google's API Style Guide](https://cloud.google.com/apis/design/)
* [Language reference for protoc3](https://developers.google.com/protocol-buffers/docs/proto3)
* [Protocol Buffer Basics: Go](https://developers.google.com/protocol-buffers/docs/gotutorial)
* [Transcoding gRPC to JSON/HTTP using Envoy](https://blog.jdriven.com/2018/11/transcoding-grpc-to-http-json-using-envoy/)

## License
[Apache 2.0](https://github.com/prysmaticlabs/ethereumapis/blob/master/LICENSE)
