package main

import (
	//"log"
	"sync"
	"fmt"
	
)

// 产生死锁的场景

//这里复制了一个锁，造成了死锁
func copyTest(mu sync.Mutex) {
	//外层加锁了 这里又加了一次 
    mu.Lock()
    defer mu.Unlock()
    fmt.Println("ok")
}

func main7() {
    var mu sync.Mutex
    mu.Lock()
    defer mu.Unlock()
    copyTest(mu)
}

