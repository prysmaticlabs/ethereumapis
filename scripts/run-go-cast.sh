bazel build //eth/v1alpha1:go_grpc_gateway_library

protoPath=eth/v1/
protoGatewayPath=eth/v1_gateway/
protoFiles=$(ls eth/v1/*.proto)
for protoFile in $protoFiles
do
  if [ "${protoFile}" != "eth/v1/swagger.proto" ]; then
    protoc -I. -Ibazel-ethereumapis/external/go_googleapis/ -Ibazel-ethereumapis/$protoPath -Ibazel-ethereumapis/external/com_google_protobuf/src --experimental_allow_proto3_optional --go-cast_out=$protoPath --go-grpc_out=$protoPath "$protoFile"
    protoc -I. -Ibazel-ethereumapis/external/go_googleapis/ -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/ -Ibazel-ethereumapis/$protoPath -Ibazel-ethereumapis/external/com_google_protobuf/src --experimental_allow_proto3_optional --grpc-gateway_out=$protoGatewayPath "$protoFile"
  fi
done

ethV1PBFiles=$(ls eth/v1/github.com/prysmaticlabs/ethereumapis/eth/v1/*.pb.go)
for pbGoFile in $ethV1PBFiles
do
  filename="${pbGoFile##*/}"
  mv "$pbGoFile" "$protoPath$filename"
done

ethV1PBFiles=$(ls eth/v1_gateway/github.com/prysmaticlabs/ethereumapis/eth/v1/*.pb.gw.go)
for pbGoFile in $ethV1PBFiles
do
  filename="${pbGoFile##*/}"
  mv "$pbGoFile" "$protoGatewayPath$filename"
done

rm -rf eth/v1/github.com/
rm -rf eth/v1_gateway/github.com/

bazel build //eth/v1:go_default_library

cp "bazel-out/k8-fastbuild/bin/${protoPath}generated.ssz.go" "$protoPath/generated.ssz.go"

protoPath=eth/v1alpha1/
protoGatewayPath=eth/v1alpha1_gateway/
protoFiles=$(ls eth/v1alpha1/*.proto)
for protoFile in $protoFiles
do
  if [ "${protoFile}" != "eth/v1alpha1/swagger.proto" ]; then
    protoc -I. -Ibazel-ethereumapis/external/go_googleapis/ -Ibazel-ethereumapis/$protoPath -Ibazel-ethereumapis/external/com_google_protobuf/src --experimental_allow_proto3_optional --go-cast_out=$protoPath --go-grpc_out=$protoPath "$protoFile"
    protoc -I. -Ibazel-ethereumapis/external/go_googleapis/ -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/ -Ibazel-ethereumapis/$protoPath -Ibazel-ethereumapis/external/com_google_protobuf/src --experimental_allow_proto3_optional --grpc-gateway_out=$protoGatewayPath "$protoFile"
  fi
done

protoc -I. -Ibazel-ethereumapis/external/go_googleapis/ -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/ -Ibazel-ethereumapis/$protoPath -Ibazel-ethereumapis/external/com_google_protobuf/src --openapiv2_out=$protoGatewayPath "${protoPath}/swagger.proto"


ethV1PBFiles=$(ls eth/v1alpha1/github.com/prysmaticlabs/ethereumapis/eth/v1alpha1/*.pb.go)
for pbGoFile in $ethV1PBFiles
do
  filename="${pbGoFile##*/}"
  mv "$pbGoFile" "$protoPath$filename"
done

ethV1PBFiles=$(ls eth/v1alpha1_gateway/github.com/prysmaticlabs/ethereumapis/eth/v1alpha1/*.pb.gw.go)
for pbGoFile in $ethV1PBFiles
do
  filename="${pbGoFile##*/}"
  mv "$pbGoFile" "$protoGatewayPath$filename"
done

rm -rf eth/v1alpha1/github.com/
rm -rf eth/v1alpha1_gateway/github.com/

bazel build //eth/v1alpha1:go_default_library

cp "bazel-out/k8-fastbuild/bin/${protoPath}generated.ssz.go" "$protoPath/generated.ssz.go"

