package main

import (
	"fmt"
	
)

func main() {
	cache := make([]int,1) // 未初始化 0 
	cache2 := []int{1}
	cache = append(cache,2)
	cache3 := append(cache2,2)
	fmt.Println(cache)
	fmt.Println(cache3)
}