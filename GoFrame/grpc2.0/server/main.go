//go_learn_server，提供指令
package main

import (
	"go_learn/Public/golang/public/config"
	//"go_learn/GoFrame/grpc/server/service"
	//"go_learn/GoFrame/grpc/server/db"
	//"go_learn/GoFrame/grpc/server/redis"

	"fmt"
	"net"
	"os"
	//"os/exec"
	"os/signal"
	//"strings"
	"syscall"

	"github.com/beego/beego/v2/core/logs"
	"google.golang.org/grpc"
	//"google.golang.org/grpc/credentials"
)

func main() {
	// exit signal
	quitChain := make(chan os.Signal)
	// listen signal
	signal.Notify(quitChain, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// 启动rpc服务
	server_config := config.GetServerConfig()
	port := server_config.Port
	logs.Info(fmt.Sprintf("port is %d", port))
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
	logs.Info("listen success")
	var golearnServer *grpc.Server

	// 带tls校验
	// certFile := server_config.CertFile 
	// keyFile := server_config.KeyFile
	// if creds, err := credentials.NewServerTLSFromFile(certFile, keyFile); err != nil {
	// 	panic(err)
	// }else{
	// 	golearnServer = grpc.NewServer(grpc.Creds(creds))
	// }
	golearnServer = grpc.NewServer()
	service.RegisterService(golearnServer)
	go golearnServer.Serve(lis)
	logs.Info(fmt.Sprintf("run success"))


	// 阻塞退出信号
	for s := range quitChain {
		switch s {
		case syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			// 退出操作
			logs.Warn("stop go_learn_server")
			golearnServer.GracefulStop()
			logs.Warn("go_learn_server stopped")

			os.Exit(0)
		default:
			logs.Warn("other signal:", s)
		}
	}
}

// trigger ci 1
