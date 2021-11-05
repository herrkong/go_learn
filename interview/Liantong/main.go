package main

import (
	"fmt"
	//"sync"
	
)

// // 传入切片 会改变其值
// func Change(a []int){
// 	a[0] = 1

// }

// // 传入数组 只是拷贝 不会改变其值
// func Change1(a [3]int){
// 	a[0] = 1

// }

// func main(){
// 	a:=[]int{9,8,7}
// 	b:=[3]int{9,8,7}

// 	Change(a)
// 	fmt.Println(a)

// 	Change1(b)
// 	fmt.Println(b)

// }

// 这里是传的切片 
// func Change (a []int){
// 	a[0] = 1

// }

// func main(){
// 	// cap = 3
// 	A := make([]int,3,3)

// 	fmt.Println(A)

// 	for i:= 0 ; i<3 ;i++{
// 		A[i] = i
// 	}

// 	// 0 1 
// 	B := A[0:2]

// 	fmt.Println(B)

// 	// 超过容量了 移动到新地址

// 	A = append(A, []int{3,4,5}...)

// 	fmt.Println(A)

// 	// 但是并没有改变其值
// 	Change(B)

// 	fmt.Println(A)


// }



// func test(a int) (value int){
// 	defer func(b int){
// 		value = b + 1

// 	}(a)
// 	a = a + 10
// 	return a

// }

// func main(){
// 	fmt.Println(test(10))
// }



// func test(index string,a,b int) int{
// 	ret := a + b 
// 	fmt.Println(index,a,b,ret)
// 	return ret
// }

// func main(){
// 	a :=1 
// 	b := 2
// 	defer test("1",a,test("10",a,b))
//     a = 0
// 	defer test("2",a,test("20",a,b))
// 	b = 1
// }



// type Animal interface{

// }

// type Cat struct{}

// func Factory(a Animal){
// 	if a != nil{
// 		fmt.Println("not nil")
// 	}else{
// 		fmt.Println("nil")
// 	}

// }

// func main(){
// 	var c * Cat
// 	Factory(c)
// }




// const (
// 	A,B = iota ,iota + 1
// 	_,_
// 	C,D
// )

// func main(){
// 	fmt.Println(A,B,C,D)
	
// }





// func main() {

// 	defer func() {fmt.Println("1")}()
// 	defer func() {fmt.Println("2")}()
// 	defer func() {fmt.Println("3")}()
// 	panic(4)
	
// }





// func main() {
// 	wg := sync.WaitGroup{}
// 	//wg.Add(20) // 开20个工作协程
// 	for i:= 0 ;i < 10 ;i++{
// 		wg.Add(1)
// 		go func(){
// 			defer func(){
// 				wg.Done()
// 			}()
// 			fmt.Println(i)
// 		}()
// 		//go f(i,&wg)
// 	}

// 	wg.Wait()
// }

// func main(){
// 	c1 := make(chan int,1)

// 	c1 <-1 

// 	select{
// 	case a := <- c1:
// 		fmt.Println(a)
// 	}
// }


// func main() {
// 	cache := make([]int,1) // 未初始化 0 
// 	cache = append(cache,2)
// 	fmt.Println(cache)
// }