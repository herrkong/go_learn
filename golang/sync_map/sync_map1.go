package main

import (
	"fmt"
	"sync"
	"time"
)



func dosth (m sync.Map) {
	i := 0
	for i < 10000 {
		m.Store(1,1)
		i++
	}
}

func main() {
	m := sync.Map{}
	//// Store sets the value for a key.
	m.Store(1,1)
	go dosth(m)
	go dosth(m)

	time.Sleep(1*time.Second)
// Load returns the value stored in the map for a key, or nil if no
// value is present.
// The ok result indicates whether value was found in the map.


	m.LoadOrStore(2,33)
	fmt.Println(m.Load(1))
	fmt.Println(m.Load(2))
}


