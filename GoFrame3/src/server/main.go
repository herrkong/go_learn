package main

import (
	"context"
	"fmt"
	"github.com/herrkong/go_learn/GoFrame3"
	"log"
	"net"

	"google.golang.org/grpc"
)

var(
	port string = "8088"
	host string = "127.0.0.1"
)


type ChatServer struct{}


func (c * ChatServer) GetData(ctx context.Context,in * ChatFormat.Data) ( out *ChatFormat.Data,err error){
	log.Printf("Client Say:%s\n",in.message)
	var response string 
	fmt.Scanln(&response)
	out.message = response
	return out,nil
}


func main() {
	// 监听端口
	lis,err := net.Listen("tcp",host+":"+port)
	if err != nil {
		log.Fatalln("faile listen at: " + host + ":" + port)
	} else {
		log.Println("ChatServer is listening at: " + host + ":" + port)
	}

	// 新建一个grpc服务器
	grpcServer := grpc.NewServer()

	// 向grpc服务器注册 ChatServer
	ChatFormat.RegisterChatServerServer(grpcServer,&ChatServer{})

	grpcServer.Serve(lis)


}



	// // 退出信号
	// quitChain := make(chan os.Signal)
	// // 监听信号
	// signal.Notify(quitChain, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// // 启动rpc服务
	// port := pub_config.ExternalServerEnv.Server[pub_config.Program.Country].Port
	// logs.Info(fmt.Sprintf("port is %d", port))
	// lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	// if err != nil {
	// 	panic(err)
	// }
	// logs.Info("listen success")
	// var externalServer *grpc.Server
	// if pub_config.Program.Mode == pub_config.PROD {
	// 	certFile := pub_config.ExternalServerEnv.Server[pub_config.Program.Country].CertFile
	// 	keyFile := pub_config.ExternalServerEnv.Server[pub_config.Program.Country].KeyFile
	// 	if creds, err := credentials.NewServerTLSFromFile(certFile, keyFile); err != nil {
	// 		panic(err)
	// 	} else {
	// 		externalServer = grpc.NewServer(grpc.Creds(creds))
	// 	}
	// } else {
	// 	externalServer = grpc.NewServer()
	// }
	// server.RegisterService(externalServer)
	// go externalServer.Serve(lis)
	// logs.Info(fmt.Sprintf("run success"))

	// // Lark通报启动情况
	// exeFile, _ := exec.LookPath(os.Args[0])
	// fileInfo, _ := os.Stat(exeFile)
	// _ = lark.SendText(config.Lark.Url, config.Lark.BotId, fileInfo.Name()+"启动", strings.Join(os.Args, ` `))

	// // 阻塞退出信号
	// for s := range quitChain {
	// 	switch s {
	// 	case syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
	// 		// 退出操作
	// 		logs.Warn("停止external_server")
	// 		externalServer.GracefulStop()
	// 		logs.Warn("external_server已停止")

	// 		os.Exit(0)
	// 	default:
	// 		logs.Warn("其他信号:", s)
	// 	}
	// }