package main

import (
	"fmt"
	"time"
)


func main() {

	height := int64(10)
	heighest := int64(20)
	goroutine_count := 20

	useMultiGoroutines := true

	if !useMultiGoroutines {
		HandleHeightAlone(height, heighest)
	} else {
		HandleHeightMultiGroutines(height, heighest, goroutine_count)
	}

}

func HandleHeightAlone(height int64, heighest int64) {

	begin_time := time.Now()

	for height < heighest {
		height = tasksingle(height)
	}
	end_time := time.Now()

	cost_time := end_time.Unix() - begin_time.Unix()

	fmt.Printf("\nHandleHeightAlone,cost_time=%d\n", cost_time)

}

func task(height int64, heightchan chan int64) (new_height int64) {
	time.Sleep(time.Second)
	fmt.Printf("height=%d\n", height+1)
	heightchan <- (height + 1)
	return height + 1
}

func tasksingle(height int64) (new_height int64) {
	time.Sleep(time.Second)
	fmt.Printf("height=%d\n", height+1)
	return height + 1
}

func HandleHeightMultiGroutines(height int64, heighest int64, goroutine_count int) {

	begin_time := time.Now()

	heightchan := make(chan int64,1)
	//quitchan := make(chan int, 1)

	//heightchan <- height

	for i := 0; i < goroutine_count; i++ { 
		go func(chan int64) {
			select {
			// case <-quitchan:
			// 	fmt.Printf("error quit chan")
			// 	break
			case <-heightchan:
				height := <-heightchan
				if height < heighest {
					task(height, heightchan)
				}
			}
		}(heightchan)
	}

	close(heightchan)

	end_time := time.Now()

	cost_time := end_time.Unix() - begin_time.Unix()

	fmt.Printf("\n,HandleHeightMultiGroutines,cost_time=%d", cost_time)

}


// func main() {
// 	deposit:= 10
// 	waiting := make(chan struct{})
// 	go func(){
// 		for i:=0; i < 100000000; i++{
// 			deposit++
// 		}
// 		waiting <- struct{}{}
// 	}()

// 	for i:=0; i < 100000000; i++{
// 		deposit--
// 	}

// 	<-waiting    // 等待子goroutine结束才结束打印余额，否则可能子goroutine都还没执行完1亿次循环就打印余额了，这样的话余额肯定不准。
// 	fmt.Println(deposit)
// }

