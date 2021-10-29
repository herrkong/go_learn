package main

import (
	"fmt"
	"sync"

)

func readchan(ch chan int){
	fmt.Println(<-ch)
}


func main() {
	defer func(){
		if errs := recover();errs != nil{
			fmt.Printf("errs=%v\n",errs)
		}
	}()

	// 带缓冲的chan
	//ch := make(chan int,1)
	// 不带缓冲chan 会报error 可能先读后写了 有人消费才会写

	
	// 同步
	ch := make(chan int,1)
	ch<- 2

	wg := sync.WaitGroup{}

	wg.Add(1)

	go func(){
		readchan(ch)
		wg.Done()
	}()


	wg.Wait()

}