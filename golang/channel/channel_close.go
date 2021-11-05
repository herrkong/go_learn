package main

import (
	"fmt"
	//"log"
)

func main() {
	defer func(){
		if errs:=recover();errs!=nil{
			fmt.Println(errs)
		}
	}()

	ch := make(chan int, 2)

	ch<- 1
	ch <- 2

	a := <-ch

	fmt.Println(a)

	close(ch)

	// 协程关闭后 可以再读 但是不可再写
	b:= <-ch
	fmt.Println(b)

	ch<-3
	c:= <-ch
	fmt.Println(c)




	
}