package main

import (
	"fmt"
	
)

//Golang中通道是进行数据同步一个重要手段，当主进程读取空通道，或者向没有协程读取的通道写入时候，
//都会发生死锁现象（编译时候提示fatal error: all goroutines are asleep - deadlock!）。下面列出几个常见死锁情况。


// 对无缓冲通道先行写入 写读后写 都会死锁

// func main() {
// 	ch := make(chan int)
// 	fmt.Println(<-ch)
// 	ch <- 100
	
// 	// ch <- 100
// 	// fmt.Println(<-ch)
// }


// func main() {
// 	ch := make(chan int)
// 	ch <- 100
// 	go func() {
// 		fmt.Println(<-ch)
// 	}()
// }


// 有缓冲通道
func main(){
	ch := make(chan int,1)
	// 先读后写 会死锁 
	fmt.Println(<-ch)
	ch<-1
	

}