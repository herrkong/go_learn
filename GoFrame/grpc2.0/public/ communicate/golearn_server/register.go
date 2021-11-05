package external_server

import (
	"threshold/src/public/config"
	"threshold/src/public/data/external_data"
)

//向外部Server Pull指令
func RegisterPullCmdFromExternalServerEvent() {
	e := Event{
		Id:          IdPullCmdFromExternalServer,
		ServiceName: config.ExternalServerEnv.ServiceName,
		MethodName:  "pull_cmd_from_external_server",
	}
	e.Request = func(serverInfo config.ServerInfo, in interface{}) (interface{}, error) {
		reply := external_data.PullCmdFromExternalServerReply{}
		err := e.Send(serverInfo, in, &reply)
		return &reply, err
	}
	Com.Register(&e)
}

//向外部Server Push指令结果
func RegisterPushResultToExternalServerEvent() {
	e := Event{
		Id:          IdPushResultToExternalServer,
		ServiceName: config.ExternalServerEnv.ServiceName,
		MethodName:  "push_result_to_external_server",
	}
	e.Request = func(serverInfo config.ServerInfo, in interface{}) (interface{}, error) {
		reply := external_data.PushResultToExternalServerReply{}
		err := e.Send(serverInfo, in, &reply)
		return &reply, err
	}
	Com.Register(&e)
}

func init() {
	//注册各种事件
	RegisterPullCmdFromExternalServerEvent()
	RegisterPushResultToExternalServerEvent()
}
