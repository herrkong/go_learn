package main

import "fmt"

func main(){
  s := []int{5}
    
  s = append(s,7)
  fmt.Println("cap(s) =", cap(s), "ptr(s) =", &s[0])
    
  s = append(s,9)
  fmt.Println("cap(s) =", cap(s), "ptr(s) =", &s[0])
    
  x := append(s, 11)
  fmt.Println("cap(s) =", cap(s), "ptr(s) =", &s[0], "ptr(x) =", &x[0])
    
  y := append(s, 12)
  fmt.Println("cap(s) =", cap(s), "ptr(s) =", &s[0], "ptr(y) =", &y[0])
}