package main

import (
	//"fmt"
	"log"
	//"sort"
	
)

// 快速排序的思想
// pivot
// partition

// 通过一趟排序将要排序的数据分割成独立的两部分，其中一部分的所有数据都比另外一部分的所有数据都要小，
// 然后再按此方法对这两部分数据分别进行快速排序，整个排序过程可以递归进行，以此达到整个数据变成有序序列。


// func Partition(){

// }

// func quickSort(){

// }

// func Min(x ,y int)(int,int){
// 	if x <= y {
// 		return x,y
// 	}
// 	return y,x
// }

// func Max(x,y int) (int,int){
// 	if x >= y{
// 		return x,y
// 	}
// 	return y,x
// }


func QuickSort(data []int,left int ,right int){
	if left > right{
		return 
	}
	i,j,pivot := left,right, data[left]

	for i < j{

		for data[j] >= pivot && i < j{
			j--
		}	

		for data[i] <= pivot && i < j{
			i++
		}

		data[i],data[j]= data[j],data[i]
	}
	data[i],data[left] = data[left],data[i]
	QuickSort(data,left,i-1)
	QuickSort(data,i+1,right)
}


func PrintSlice(data []int){
	for idx,d := range data{
		log.Printf("data[%v]=%v\n",idx,d)
	}
}


func main() {
	data := []int{6,2,7,7,3,8,5,1,9}
	PrintSlice(data)
	QuickSort(data,0,len(data)-1)
	//sort.Ints(data)
	log.Println("-----------")
	PrintSlice(data)

	
}