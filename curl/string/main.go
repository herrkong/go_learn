package main

import "fmt"

func main()  {
	
	Alphabet := "23456789ABCDEFGHJKLMNPQRSTUVWXYZ"

	Alphabet_slice := []byte(Alphabet)

	fmt.Printf("Alphabet[20]=%v,Alphabet_slice[20]=%v\n",Alphabet[20],string(Alphabet_slice[20]))


}