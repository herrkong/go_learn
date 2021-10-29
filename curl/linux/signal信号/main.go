package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)


func main(){

	// 创建 信号接收channel 和程序退出channel
	sigs := make(chan os.Signal,1)
	done := make(chan bool,1)


	//指定channel来接收特定信号
	signal.Notify(sigs,syscall.SIGINT,syscall.SIGTERM)


	//阻塞等待信号 直至信号通知退出
	go func(){
		sig := <-sigs
		fmt.Printf("sig=%v\n",sig)
		done <- true
	}()

	fmt.Printf("awaiting signal\n")
	<-done 
	fmt.Printf("exit")

}