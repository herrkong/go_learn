// package main

// import (
// 	//"log"
// )

// go 1.17 优化抛出的错误堆栈

// func main() {
// 	example(make([]string, 1, 2), "煎鱼", 3)
// }

// //go:noinline
// func example(slice []string, str string, i int) error {
// 	panic("脑子进煎鱼了")
// }

package main
import (
	"fmt"
	"runtime/debug"
)
 
// 异常处理函数1
func panicDeal1() {
	fmt.Println("panicDeal1,begin")
	if err := recover(); err != nil {
		fmt.Println("err1:", err)          // 打印出异常(由于panicDeal2()中的recover函数已经捕获了异常,所以这里捕获不到异常,不会得到执行)
		fmt.Println(string(debug.Stack())) // 打印出堆栈信息
	}
	fmt.Println("panicDeal1,end")
}
 
// 异常处理函数2
func panicDeal2() {
	fmt.Println("panicDeal2,begin")
	if err := recover(); err != nil {
		fmt.Println("err2", err)           // 打印出异常
		fmt.Println(string(debug.Stack())) // 打印出堆栈
	}
	fmt.Println("panicDeal2,end")
}
 
func test() {
	fmt.Println("1111")
 
	// 必须先声明defer,否则不能捕获panic异常
	defer panicDeal1()
 
	// 触发panic时,逆序执行,也就是先执行 panicDeal2(),再执行 panicDeal1()
	defer panicDeal2()
 
	fmt.Println("2222")
 
	// 空指针赋值,产生崩溃
	var p *int
	*p = 1
 
	// 这里的代码得不到执行
	fmt.Println("3333")
}
 
func main() {
	test()
}