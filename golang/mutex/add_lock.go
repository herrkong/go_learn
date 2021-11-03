package main

import (
	"fmt"
	"sync"
)

// 加锁 累加10000

func main2() {
	wg := sync.WaitGroup{}
	lock := sync.Mutex{}
	n := 10
	count :=0
	wg.Add(n)
	for i:=0 ; i< n ; i++{
		go func(){
			for j:=0;j < 10000;j++{
				lock.Lock()
				count++
				lock.Unlock()
			}
			wg.Done()
		}()
	} 
	wg.Wait()
	fmt.Println(count)
}


