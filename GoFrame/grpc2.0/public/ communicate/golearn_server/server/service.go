package server

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"threshold/src/public/communicate/external_server"
	"threshold/src/public/config"
	"threshold/src/public/data/external_data"
)

type ExternalServiceEr interface {
	//向外部Server pull指令
	PullCmdFromExternalServerReply(ctx context.Context, in *external_data.PullCmdFromExternalServerRequest) (*external_data.PullCmdFromExternalServerReply, error)
	//向外部Server push指令执行结果
	PushResultToExternalServerReply(ctx context.Context, in *external_data.PushResultToExternalServerRequest) (*external_data.PushResultToExternalServerReply, error)
}

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

func init() {
	//向外部Server Pull指令
	if event := external_server.Com.Find(external_server.IdPullCmdFromExternalServer); event == nil {
		panic("not find IdPullCmdFromExternalTransfer event")
	} else {
		event.Handler = func(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
			in := new(external_data.PullCmdFromExternalServerRequest)
			if err := dec(in); err != nil {
				return nil, err
			}
			if interceptor == nil {
				return (srv.(*ExternalService)).PullCmdFromExternalServerReply(ctx, in)
			}
			info := &grpc.UnaryServerInfo{
				Server:     srv,
				FullMethod: fmt.Sprintf("/%s/%s", event.ServiceName, event.MethodName),
			}
			handler := func(ctx context.Context, req interface{}) (interface{}, error) {
				return srv.(*ExternalService).PullCmdFromExternalServerReply(ctx, req.(*external_data.PullCmdFromExternalServerRequest))
			}
			return interceptor(ctx, in, info, handler)
		}
	}
	//向外部Server Push指令结果
	if event := external_server.Com.Find(external_server.IdPushResultToExternalServer); event == nil {
		panic("not find IdPushResultToExternalServer event")
	} else {
		event.Handler = func(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
			in := new(external_data.PushResultToExternalServerRequest)
			if err := dec(in); err != nil {
				return nil, err
			}
			if interceptor == nil {
				return (srv.(*ExternalService)).PushResultToExternalServerReply(ctx, in)
			}
			info := &grpc.UnaryServerInfo{
				Server:     srv,
				FullMethod: fmt.Sprintf("/%s/%s", event.ServiceName, event.MethodName),
			}
			handler := func(ctx context.Context, req interface{}) (interface{}, error) {
				return srv.(*ExternalService).PushResultToExternalServerReply(ctx, req.(*external_data.PushResultToExternalServerRequest))
			}
			return interceptor(ctx, in, info, handler)
		}
	}
}
