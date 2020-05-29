#!/bin/bash

# Script to copy pb.go files from bazel build folder to appropriate location.
# Bazel builds to bazel-bin/... folder, script copies them back to original folder where .proto is.

bazel query 'kind(go_library, //...) union kind(go_proto_library, //...)' | xargs bazel build

# Get locations of pb.go files.

file_list=()
while IFS= read -d $'\0' -r file ; do
    file_list=("${file_list[@]}" "$file")
done < <(find -L $(bazel info bazel-bin)/eth -type f -regextype sed -regex ".*pb\.\(gw\.\)\?go$" -print0)

arraylength=${#file_list[@]}
searchstring="/ethereumapis/"

# Copy pb.go files from bazel-bin to original folder where .proto is.
for (( i=0; i<${arraylength}; i++ ));
do
  destination=${file_list[i]#*$searchstring}
  echo ${destination}
  cp -R -L "${file_list[i]}" "$destination"
  echo Updated $destination
done

# Copy the generated.ssz.go file to live in the same package as 
# the generated protos.
file_list=()
while IFS= read -d $'\0' -r file ; do
    file_list=("${file_list[@]}" "$file")
done < <(find -L $(bazel info bazel-bin)/eth -type f -name "*ssz.go" -print0)

arraylength=${#file_list[@]}
searchstring="/bin/"

for (( i=0; i<${arraylength}; i++ ));
do
  destination=${file_list[i]#*$searchstring}
  echo ${destination}
  cp -R -L "${file_list[i]}" "$destination"
  echo Updated $destination
done
