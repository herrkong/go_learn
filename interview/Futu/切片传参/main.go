package main

import (
	"fmt"
)

func Append(s []int){
	s = append(s, 5)
}

func Add(s []int){
	for i := range s{
		s[i] = s[i] + 5
	}
}


func main() {
	var s = []int{1,2,3,4}
	Append(s)
	fmt.Println(s)

	// Append2(s)
	// fmt.Println(s)

	// Append3(&s)
	// fmt.Println(s)

	Add(s)
	fmt.Println(s)
	
}



func Append2(s []int) []int    {
	s = append(s, 5)
	return s
}

// 传入切片指针
func Append3(s *[]int) {
	*s = append(*s, 5)
}