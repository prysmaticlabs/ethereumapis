ethV1Path=eth/v1/
ethV1GatewayPath=eth/v1_gateway/
ethV1Files=$(ls eth/v1/*.proto)
for protoFile in $ethV1Files
do
  protoc -I. -Ibazel-ethereumapis/external/go_googleapis/ -Ibazel-ethereumapis/$ethV1Path -Ibazel-ethereumapis/external/com_google_protobuf/src --experimental_allow_proto3_optional "$protoFile" --go-cast_out=$ethV1Path
  protoc -I. -Ibazel-ethereumapis/external/go_googleapis/ -Ibazel-ethereumapis/$ethV1Path -Ibazel-ethereumapis/external/com_google_protobuf/src --experimental_allow_proto3_optional --grpc-gateway_out $ethV1GatewayPath "$protoFile"
done

ethV1PBFiles=$(ls eth/v1/github.com/prysmaticlabs/ethereumapis/eth/v1/*.pb.go)
for pbGoFile in $ethV1PBFiles
do
  filename="${pbGoFile##*/}"
  mv "$pbGoFile" "$ethV1Path$filename"
done

ethV1PBFiles=$(ls eth/v1_gateway/github.com/prysmaticlabs/ethereumapis/eth/v1/*.pb.go)
for pbGoFile in $ethV1PBFiles
do
  filename="${pbGoFile##*/}"
  mv "$pbGoFile" "$ethV1GatewayPath$filename"
done

rm -rf eth/v1/github.com/
rm -rf eth/v1_gateway/github.com/

bazel build //eth/v1:go_default_library

cp "bazel-out/k8-fastbuild/bin/${ethV1Path}generated.ssz.go" "$ethV1Path/generated.ssz.go"

ethV1Path=eth/v1alpha1/
ethV1GatewayPath=eth/v1alpha1_gateway/
ethV1Files=$(ls eth/v1alpha1/*.proto)
for protoFile in $ethV1Files
do
  protoc -I. -Ibazel-ethereumapis/external/go_googleapis/ -Ibazel-ethereumapis/$ethV1Path -Ibazel-ethereumapis/external/com_google_protobuf/src --experimental_allow_proto3_optional "$protoFile" --go-cast_out=$ethV1Path
done

ethV1PBFiles=$(ls eth/v1alpha1/github.com/prysmaticlabs/ethereumapis/eth/v1alpha1/*.pb.go)
for pbGoFile in $ethV1PBFiles
do
  filename="${pbGoFile##*/}"
  mv "$pbGoFile" "$ethV1Path$filename"
done

ethV1PBFiles=$(ls eth/v1_gateway/github.com/prysmaticlabs/ethereumapis/eth/v1alpha1/*.pb.go)
for pbGoFile in $ethV1PBFiles
do
  filename="${pbGoFile##*/}"
  mv "$pbGoFile" "$ethV1GatewayPath$filename"
done

rm -rf eth/v1alpha1/github.com/
rm -rf eth/v1alpha1_gateway/github.com/

bazel build //eth/v1alpha1:go_default_library

cp "bazel-out/k8-fastbuild/bin/${ethV1Path}generated.ssz.go" "$ethV1Path/generated.ssz.go"

