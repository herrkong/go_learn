package main

import "fmt"

// go 内存问题  超过cap之后会开辟新的内存 


func main(){

	slice := make([]int,2,5)
	fmt.Printf("slice=%+v,&slice=%p,len(slice)=%v,&slice[0]=%+v\n",slice,&slice,len(slice),&slice[0])
	slice = append(slice,2)
	fmt.Printf("slice=%+v,&slice=%p,len(slice)=%v,&slice[0]=%+v\n",slice,&slice,len(slice),&slice[0])
	slice = append(slice,2)
	fmt.Printf("slice=%+v,&slice=%p,len(slice)=%v,&slice[0]=%+v\n",slice,&slice,len(slice),&slice[0])
	slice = append(slice,2)
	fmt.Printf("slice=%+v,&slice=%p,len(slice)=%v,&slice[0]=%+v\n",slice,&slice,len(slice),&slice[0])
	slice = append(slice,2)
	fmt.Printf("slice=%+v,&slice=%p,len(slice)=%v,&slice[0]=%+v\n",slice,&slice,len(slice),&slice[0])
	slice = append(slice,2)
	fmt.Printf("slice=%+v,&slice=%p,len(slice)=%v,&slice[0]=%+v\n",slice,&slice,len(slice),&slice[0])
	slice = append(slice,2)
	fmt.Printf("slice=%+v,&slice=%p,len(slice)=%v,&slice[0]=%+v\n",slice,&slice,len(slice),&slice[0])




}