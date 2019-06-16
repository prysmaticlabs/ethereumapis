# Ethereum APIs

This repository hosts service interface definitions of Ethereum (2.0) APIs.
These definitions support both REST and gRPC protocols. 

## Versioning Schedule

TODO

## Building / Generating Client Libraries

TODO

## RESTful Endpoints (gRPC Gateway)

All of the gRPC services define JSON over HTTP endpoints. Client implementations
may choose to bundle a REST service of gRPC with their client binaries or
they may use a JSON encoding proxy such as [Envoy Proxy](https://www.envoyproxy.io/),
[grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway), or others.

## JSON Mapping

The majority of field types for Ethereum are `bytes` or `uint64` type. The JSON
mapping for those fields are Base64 encoded string for `bytes` or a string
representation of `uint64`. If the field value is set to the default value in
the protobuf, this field will be omitted from the JSON encoding. 

For more details on JSON mapping for other types, view the relevent section in 
the [proto3 language guide](https://developers.google.com/protocol-buffers/docs/proto3#json).

## Helpful gRPC Tooling and Resources

- [Awesome gRPC](https://github.com/grpc-ecosystem/awesome-grpc)
- [Google's API Style Guide](https://cloud.google.com/apis/design/)