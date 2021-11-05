package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	//wg.Add(20) // 开20个工作协程
	for i:= 0 ;i < 10 ;i++{
		wg.Add(1)
		go func(){
			defer func(){
				wg.Done()
			}()
			fmt.Println(i)
		}()
		//go f(i,&wg)
	}

	wg.Wait()
}