#! /bin/bash

# Run: source ./scripts/protogen.sh from the root of the project

mkdir -p lib
rm -rf lib/go/proto

protoc \
  --go_out=lib/go \
  --go-grpc_out=lib/go \
  --grpc-gateway_out=lib/go \
  --go_opt=module=github.com/kinjalrk2k/Kaho-aur-Suno/lib/go \
  --go-grpc_opt=module=github.com/kinjalrk2k/Kaho-aur-Suno/lib/go \
  --grpc-gateway_opt=module=github.com/kinjalrk2k/Kaho-aur-Suno/lib/go \
  --proto_path=proto \
  **/*.proto