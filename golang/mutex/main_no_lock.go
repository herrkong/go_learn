package main

import (
	"fmt"
	"sync"
)

// 不加锁 累加10000

func main4() {
	wg := sync.WaitGroup{}
	n := 10
	count :=0
	wg.Add(n)
	for i:=0 ; i< n ; i++{
		go func(){
			defer wg.Done()
			for j:=0;j < 10000;j++{
				count++
			}
		}()
	} 
	wg.Wait()
	fmt.Println(count)
}


