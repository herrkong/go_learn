package main

import (
	//"log"
)

// go 1.17 优化抛出的错误堆栈

func main() {
	example(make([]string, 1, 2), "煎鱼", 3)
}

//go:noinline
func example(slice []string, str string, i int) error {
	panic("脑子进煎鱼了")
}

