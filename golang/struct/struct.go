package main

import (
	"fmt"
	//"unsafe"
)

// 1 空结构体不占用任何空间
// func main() {
// 	fmt.Println(unsafe.Sizeof(struct{}{}))
// }



//Go 语言标准库没有提供 Set 的实现，通常使用 map 来代替。事实上，对于集合来说，只需要 map 的键，而不需要值。
//即使是将值设置为 bool 类型，也会多占据 1 个字节，那假设 map 中有一百万条数据，就会浪费 1MB 的空间。


// type Set map[string]struct{}

// func (s Set) Has(key string) bool {
// 	_, ok := s[key]
// 	return ok
// }

// func (s Set) Add(key string) {
// 	s[key] = struct{}{}
// }

// func (s Set) Delete(key string) {
// 	delete(s, key)
// }

// func main() {
// 	s := make(Set)
// 	s.Add("Tom")
// 	s.Add("Sam")
// 	fmt.Println(s.Has("Tom"))
// 	fmt.Println(s.Has("Jack"))
// }



//不发送数据的信道(channel)

// func worker(ch chan struct{}) {
// 	<-ch
// 	fmt.Println("do something")
// 	close(ch)
// }

// // 
// func main() {
// 	ch := make(chan struct{})
// 	go worker(ch)
// 	ch <- struct{}{}
	
// }


// 仅包含方法的结构体

type Door struct{}

func (d Door) Open() {
	fmt.Println("Open the door")
}

func (d Door) Close() {
	fmt.Println("Close the door")
}

func main(){
	d := Door{}
	d.Open()
	d.Close()
}