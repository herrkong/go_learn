package main

import (
	"fmt"
	
)

// 循环数组

// 升序 数组 向右平移得到输入数组 k位 求k
// 

func Min(x ,y int) bool{
	return x < y
}

func Max(x,y int) bool{
	return x > y
}


func GetK(sl []int) int{
	for i:=0 ;i < len(sl) - 1;i++{
		if (sl[i+1] < sl[i]){
			return i + 1 
		}
	}
	return 0
}


func main() {
	a := []int{5,1,2,3,4}
	fmt.Println(GetK(a))
}