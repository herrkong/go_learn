package main

import (
	"fmt"
	
)



// 快速排序的思想
// pivot
// partition


func Partition(){

}

func quickSort(){

}

func QuickSort(data []int){
	lens := len(data)


}


func PrintSlice(data []int){
	for idx,d := range data{
		fmt.Printf("data[%v]=%v\n",idx,d)
	}
}


func main() {
	data := []int{6,2,7,7,3,8,5,1,9}
	PrintSlice(data)
	QuickSort(data)
	PrintSlice(data)

	
}