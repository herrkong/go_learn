package main

import (
    "fmt"
	"runtime"
)

func main() {
	// 主动让出时间片
	runtime.Gosched()
    endlessLoop3()
}

func endlessLoop1() {
    //i := 10
	// // go 写死循环会一直卡死
    // for {
	// 	// for循环中不做任何系统调用就好了
    //     fmt.Println("i =", i)
    // }
	//<-make(chan struct{})
	//<-make(chan struct{})

	





}

func endlessLoop2() {
    for i, j := 1, 10; i < j; i, j = i+1, j+1 {
        fmt.Println("i =", i, " j=", j) // j - i = 9
    }
}

func endlessLoop3() {
    i := 20
    select {
    default:
        fmt.Println("i =", i)
        endlessLoop3() // 递归
    }
}

func endlessLoop4() {
Here1:
    fmt.Println("Here1")
    goto Here2
Here2:
    fmt.Println("Here2")
    goto Here1
}

func endlessLoopEach5() {
    fmt.Println("Each5")
    defer endlessLoopOther5()
}

func endlessLoopOther5() {
    fmt.Println("Other5")
    defer endlessLoopEach5()
}

func endlessLoop6() {
    //定义一个匿名函数，函数的参数f是用一另外一个函数 endlessLoop6()的名称作为参数且另一个函数是无参数函数
    var loop = func(f func()) {
        fmt.Println("endlessLoop6")
        f()
        fmt.Println("1.本行代码永远都不会被执行") //本行上设置断点，调试运行本程序，启动和退出本程序时，均不会进入该断点
    }
    loop(endlessLoop6) //使用匿名函数去递归自己,此时执行f函数
    fmt.Println("2.本行代码永远都不会被执行") //本行上设置断点，调试运行本程序，启动和退出本程序时，均不会进入该断点
}
