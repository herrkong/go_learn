package main

import (
	//"log"
	"sync"
	"fmt"
	"time"
)

// 读写锁 相比一般的锁来说 读协程占用锁时 其他协程可以不用等待锁 直接共享资源

// 读写锁的读锁可以重入，在已经有读锁的情况下，可以任意加读锁。
// 在读锁没有全部解锁的情况下，写操作会阻塞直到所有读锁解锁。
// 写锁定的情况下，其他协程的读写都会被阻塞，直到写锁解锁。


//如果我们可以明确区分reader和writer的协程场景，且是大师的并发读、少量的并发写，有强烈的性能需要，我们就可以考虑使用读写锁RWMutex替换Mutex

func read(m *sync.RWMutex, i int) {
    fmt.Println(i, "reader start")
    m.RLock()
    fmt.Println(i, "reading")
    time.Sleep(1 * time.Second)
    m.RUnlock()

    fmt.Println(i, "reader over")
}


func main1() {
	//wg:= sync.WaitGroup{}
    var m sync.RWMutex
    go read(&m, 1)
    go read(&m, 2)
    go read(&m, 3)

    time.Sleep(2 * time.Second)
}
