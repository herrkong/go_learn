package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)	
	go foo(&wg)

	fmt.Println("before wait")
	wg.Wait()
	fmt.Println("after wait")
}

func foo(wg *sync.WaitGroup) {
	fmt.Println("before sleep")
	time.Sleep(2 * time.Second)
	fmt.Println("after sleep")
	wg.Done()
}
