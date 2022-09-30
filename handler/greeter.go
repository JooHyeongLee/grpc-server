package handler

import (
	"context"
	"golang-grpc-example/grpcapi"
	"log"
)

type GreeterHandler struct {
}

func (h GreeterHandler) SayHello(ctx context.Context, req *grpcapi.SayHelloRequest) (*grpcapi.SayHelloResponse, error) {
	log.Print("SayHello")
	return &grpcapi.SayHelloResponse{
		ResponseCode:    200,
		ResponseMessage: "Hi, Nice to meet you, too",
	}, nil
}

func (h GreeterHandler) SayGoodbye(ctx context.Context, req *grpcapi.SayGoodbyeRequest) (*grpcapi.SayGoodbyeResponse, error) {
	log.Print("SayGoodbye")
	return &grpcapi.SayGoodbyeResponse{
		ResponseCode: 200,
		ResponseMessage: "Good Goodbye",
	}, nil
}