#!/bin/sh

export PATH=$PATH:/opt/homebrew/bin:/opt/homebrew/sbin

find . -name "*.pb.go" -type f -delete
find . -name "*.pb.swift" -type f -delete
find . -name "*.js" -type f -delete

which -s brew
if [[ $? != 0 ]] ; then
    # Install Homebrew
    ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
else
    brew update
fi

brew install protobuf

brew install swift-protobuf
protoc --swift_out=. --proto_path=. $(find . -iname "*.proto")

brew install protoc-gen-go
protoc --go_out=. --proto_path=. --go_opt=paths=source_relative $(find . -iname "*.proto")

# brew install protoc-gen-js
# protoc --js_out=. --proto_path=. $(find . -iname "*.proto")

echo "Done generating protobuf files"
