package main

import (
	//"log"
	"sync"
	//"fmt"
	
)

// 循环等待 哲学家就餐问题

// A 等B B等C C等A 无限等待

func main() {
	var muA, muB sync.Mutex
    var wg sync.WaitGroup

    wg.Add(2)
    go func() {
        defer wg.Done()
        muA.Lock()
        defer muA.Unlock()
        //A依赖B
        muB.Lock()
        defer muB.Lock()
    }()

    go func() {
        defer wg.Done()
        muB.Lock()
        defer muB.Lock()
        //B依赖A
        muA.Lock()
        defer muA.Unlock()
    }()
    wg.Wait()


}