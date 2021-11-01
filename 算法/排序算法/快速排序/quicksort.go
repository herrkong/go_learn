package main

import (
	//"fmt"
	"log"
	"sort"
	
)



// 快速排序的思想
// pivot
// partition

// 通过一趟排序将要排序的数据分割成独立的两部分，其中一部分的所有数据都比另外一部分的所有数据都要小，
// 然后再按此方法对这两部分数据分别进行快速排序，整个排序过程可以递归进行，以此达到整个数据变成有序序列。


func Partition(){

}

func quickSort(){

}

func Min(x ,y int)(int,int){
	if x <= y {
		return x,y
	}
	return y,x
}

func Max(x,y int) (int,int){
	if x >= y{
		return x,y
	}
	return y,x
}


func QuickSort(data []int){
	mid := len(data) / 2
	pivot := data[mid]
	log.Printf("pivot=%d\n",pivot)
	for i:=0 ;i< len(data) ; i++{
		if i <= mid{
			pivot,data[i] = Max(pivot,data[i])
		}else{
			pivot,data[i] = Min(pivot,data[i])
		}
		log.Printf("i=%d,pivot=%d,data[i]=%d\n",i,pivot,data[i])
	}

}


func PrintSlice(data []int){
	for idx,d := range data{
		log.Printf("data[%v]=%v\n",idx,d)
	}
}



// type NeedSort []int

// type need int

// func (data NeedSort) Len()(int){
// 	return len(data)
// }

// func (data NeedSort) Less(i,j int) { 
// 	return data[i] < data[j]
// }


func main() {
	data := []int{6,2,7,7,3,8,5,1,9}
	//PrintSlice(data)
	//QuickSort(data)
	sort.Ints(data)
	log.Println("-----------")
	PrintSlice(data)

	
}