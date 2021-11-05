package external_server

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"threshold/src/public/config"
)

type ID int

const (
	IdPullCmdFromExternalServer ID = iota
	IdPushResultToExternalServer
)

type Event struct {
	Id          ID
	ServiceName string
	MethodName  string
	Request     func(config.ServerInfo, interface{}) (interface{}, error)
	Handler     func(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error)
}

func (self *Event) Send(serverInfo config.ServerInfo, in, out interface{}) error {
	address := fmt.Sprintf("%s:%d", serverInfo.Host, serverInfo.Port)
	logs.Info(fmt.Sprintf("send rpc request, address: %s, %+v", address, in))
	var conn *grpc.ClientConn
	var err error
	if config.Program.Mode == config.PROD {
		if creds, err := credentials.NewClientTLSFromFile(serverInfo.CertFile, serverInfo.Host); err != nil {
			return err
		} else {
			if conn, err = grpc.Dial(address, grpc.WithTransportCredentials(creds)); err != nil {
				return err
			}
		}
	} else {
		if conn, err = grpc.Dial(address, grpc.WithInsecure()); err != nil {
			return err
		}
	}
	defer conn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), config.ExternalServerEnv.GrpcReadTimeout)
	defer cancel()
	if err := conn.Invoke(ctx, fmt.Sprintf("/%s/%s", self.ServiceName, self.MethodName), in, out); err != nil {
		return err
	}
	return nil
}

type Communicate struct {
	Events []*Event
}

func (self *Communicate) Register(e *Event) {
	self.Events = append(self.Events, e)
}

func (self *Communicate) Find(id ID) *Event {
	for _, event := range self.Events {
		if event.Id == id {
			return event
		}
	}
	return nil
}

func (self *Communicate) Request(id ID, serverInfo config.ServerInfo, in interface{}) (interface{}, error) {
	if event := self.Find(id); event == nil {
		return nil, fmt.Errorf("not find event, id=%d", id)
	} else {
		return event.Request(serverInfo, in)
	}
}

var (
	Com Communicate
)
