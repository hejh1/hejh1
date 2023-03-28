#!/bin/sh

export PATH=$PATH:/opt/homebrew/bin:/opt/homebrew/sbin

find . -name "*.pb.go" -type f -delete
find . -name "*.swift" -type f -delete
find . -name "*.py" -type f -delete
find . -name "*.js" -type f -delete

if ! command -v brew &> /dev/null
then
    echo "brew could not be found."
    exit
else
    brew update
fi

brew install protobuf
# Swift
brew install swift-protobuf grpc-swift
protoc --swift_opt=Visibility=Public --swift_out=. --proto_path=. $(find . -iname "*.proto") --grpc-swift_opt=Client=true,Server=false --grpc-swift_out=.
# Golang
brew install protoc-gen-go
brew install protoc-gen-go-grpc
protoc --go_out=. --proto_path=. --go_opt=paths=source_relative $(find . -iname "*.proto")
protoc --go-grpc_out=. --go-grpc_opt=paths=source_relative $(find . -iname "*.proto")
# Python
export GRPC_PYTHON_BUILD_SYSTEM_OPENSSL=1
export GRPC_PYTHON_BUILD_SYSTEM_ZLIB=1
python3 -m pip install protobuf==3.20.3 grpcio grpcio-tools
python3 -m pip install -U "betterproto[compiler]" --pre
python3 -m grpc_tools.protoc -I . --python_betterproto_out=pyproto *.proto
# Javascript
# brew install protoc-gen-js
# protoc --js_out=. --proto_path=. $(find . -iname "*.proto")

echo "Done generating protobuf files"
