#!/bin/zsh

# user proto
readonly APP_NAME="user"

protoc --go_out=../ --go_opt=paths=source_relative \
    --go-grpc_out=../ --go-grpc_opt=paths=source_relative \
    api/pb/proto/$APP_NAME.proto
mkdir -p api/services/$APP_NAME/pb
rm -f api/services/$APP_NAME/pb/${APP_NAME}_grpc.pb.go
rm -f api/services/$APP_NAME/pb/$APP_NAME.pb.go
cp api/pb/${APP_NAME}_grpc.pb.go api/services/$APP_NAME/pb/
cp api/pb/$APP_NAME.pb.go api/services/$APP_NAME/pb/