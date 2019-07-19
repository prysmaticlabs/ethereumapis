# Ethereum APIs

[![Build status](https://badge.buildkite.com/62be08099e9e228b165c2dba69c637eb9ca7a1ca95efd54b9f.svg?branch=master)](https://buildkite.com/prysmatic-labs/ethereum-apis)
[![ETH2.0_Spec_Version 0.8.0](https://img.shields.io/badge/ETH2.0%20Spec%20Version-v0.8.0-blue.svg)](https://github.com/ethereum/eth2.0-specs/commit/8d324b7497bcb558e0183a30002d78d18704e3fa)

This repository hosts service interface definitions of Ethereum (2.0) APIs. These [protocol buffer](https://developers.google.com/protocol-buffers/) service definitions support gRPC and JSON over HTTP protocols.

## Motivation

Standardizing common APIs is critical for the success of Ethereum 2.0. Developers may utilize these self-describing schemas to interact with Ethereum projects and having a central repository for service definitions may reduce the disparity between services across the Ethereum ecosystem. 

## Services

| Package | Service | Version | Description |
|---------|---------|---------|-------------|
| eth | [BeaconChain](eth/v1alpha1/beacon_chain.proto#L36) | v1alpha1 | The beacon chain service can be used to access data relevant to the Ethereum 2.0 phase 0 beacon chain. |
| eth | [Node](eth/v1alpha1/node.proto#L33) | v1alpha1 | General information about the Ethereum node itself, the services it supports, chain information, and node version. |

Versioning should use [Semantic Versioning](https://semver.org/) with pre-releases using a pre-release version name such as alpha or beta.

## Building / Generating Client Libraries

Client libraries may be generated using [protoc](https://github.com/protocolbuffers/protobuf).

TODO - Add more detailed instructions here.

## RESTful Endpoints (gRPC Transcoding)

All of the gRPC services should define JSON over HTTP endpoints by specifying [HTTPRules](https://github.com/googleapis/googleapis/blob/master/google/api/http.proto). Client implementations may choose to bundle a REST service of gRPC with their client binaries or they may use a JSON encoding proxy such as [Envoy Proxy](https://www.envoyproxy.io/), [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway), or others. 

Example:
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

This enables an HTTP REST to gRPC mapping as below:

HTTP | gRPC
-----|-----
`GET /v1/messages/123456`  | `GetMessage(name: "messages/123456")`

For more information on gRPC transcoding, see the examples in [google/api/http.proto](https://github.com/googleapis/googleapis/blob/master/google/api/http.proto#L45).


## JSON Mapping

The majority of field primative types for Ethereum are `bytes` or `uint64`. The canonical JSON mapping for those fields are Base64 encoded string for `bytes` or a string representation of `uint64`. Since JavaScript loses precision for values over [MAX_SAFE_INTEGER](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Number/MAX_SAFE_INTEGER), uint64 must be string in JSON to accommodate those values. If the field value is set to the default value in the protobuf, this field will be omitted from the JSON encoding. 

For more details on JSON mapping for other types, view the relevent section in the [proto3 language guide](https://developers.google.com/protocol-buffers/docs/proto3#json).

## Helpful gRPC Tooling and Resources

- [Awesome gRPC](https://github.com/grpc-ecosystem/awesome-grpc)
- [Google's API Style Guide](https://cloud.google.com/apis/design/)

## License

Brought to you by [Prysmatic Labs](https://prysmaticlabs.com) with a permissive [Apache License 2.0](license.md).
