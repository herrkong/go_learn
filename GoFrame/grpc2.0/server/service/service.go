package service

import (
	//"fmt"
	//"golang.org/x/net/context"
	"google.golang.org/grpc"
	//"threshold/src/public/communicate/external_server"
	//"threshold/src/public/config"
	//"threshold/src/public/data/external_data"
)

type ExternalService struct {
}

func RegisterService(server *grpc.Server) {
	var methods []grpc.MethodDesc
	for _, event := range external_server.Com.Events {
		methods = append(methods, grpc.MethodDesc{
			MethodName: event.MethodName,
			Handler:    event.Handler,
		})
	}
	server.RegisterService(&grpc.ServiceDesc{
		ServiceName: config.ExternalServerEnv.ServiceName,
		HandlerType: (*ExternalServiceEr)(nil),
		Methods:     methods,
	}, &ExternalService{})
}
