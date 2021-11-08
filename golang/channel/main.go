package main

import (
    "fmt"
    
)

//  有缓冲channel 先写channel 再读channal 
func main() {
    ch := make(chan int,1)
    ch <-1
    select{
    case a:=<-ch:{
            fmt.Println(a)
        }
    }
}