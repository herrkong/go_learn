package main

import (
	"container/list"
	"fmt"
	//"log"
)


func main() {

	l := list.List{}
	l.Init()
	l.PushBack(5)
	l.PushBack(6)
	l.PushBack(7)

    fmt.Println(l)

	
	
	
}