// gen.go
package grpcapi

//go:generate -command compile_proto protoc -I../protos
//go:generate compile_proto --go_out=plugins=grpc,paths=source_relative:. greeter.proto
