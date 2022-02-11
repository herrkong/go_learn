package main
 
import (
	"fmt"
	"time"
)
 
//发送者
func sender(c chan int) {
	for i := 0; i < 20; i++ {
		c <- i
		if i >= 5 {
			time.Sleep(time.Second * 7)
		} else {
			time.Sleep(time.Second)
		}
	}
}
 
func main() {
	c := make(chan int)
	go sender(c)
	timeout := time.After(time.Second * 3)
	for {
		select {
		case d := <-c:
			fmt.Println(d)
		case <-timeout:
			fmt.Println("执行定时操作任务")
		case dd := <-time.After(time.Second * 3):
			fmt.Println(dd.Format("2006-01-02 15:04:05"), "执行超时动作")
		}
		fmt.Println("for end")
	}
}