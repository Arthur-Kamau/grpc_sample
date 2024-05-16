# GRPC Command
Run at project root.

## Golang
## in project root
protoc  --proto_path=proto --go_out=app/models --go_opt=paths=source_relative   --go-grpc_out=app/models --go-grpc_opt=paths=source_relative   app.proto
