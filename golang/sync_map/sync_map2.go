package main

import (
	//"log"
	"time"
	"fmt"
	
)

func do (m map[int]int) {
	i := 0
	for i < 10000 {
		m[1]=1
		i++
	}
}

func main4() {
	m := map[int]int {1:1}
	go do(m)
	go do(m)

	time.Sleep(1*time.Second)
	fmt.Println(m)
}

