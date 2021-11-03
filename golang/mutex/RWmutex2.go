package main

import (
    "fmt"
    "sync"
    "time"
)

var count = 0

func read1(m *sync.RWMutex, i int) {
    fmt.Println(i, "reader start")
    m.RLock()
    fmt.Println(i, "reading count:", count)
    time.Sleep(1 * time.Millisecond)
    m.RUnlock()

    fmt.Println(i, "reader over")
}

func write(m *sync.RWMutex, i int) {
    fmt.Println(i, "writer start")
    m.Lock()
    count++
    fmt.Println(i, "writing count", count)
    time.Sleep(1 * time.Millisecond)
    m.Unlock()

    fmt.Println(i, "writer over")
}

func main6() {
    var m sync.RWMutex
    for i := 1; i <= 3; i++ {
        go write(&m, i)
    }
    for i := 1; i <= 3; i++ {
        go read1(&m, i)
    }

    time.Sleep(1 * time.Second)
    fmt.Println("final count:", count)
}

