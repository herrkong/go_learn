package main

import (
	"fmt"
	"strings"
)

func main(){
	//只是声明了
	// var retMap map[string]string
	//直接初始化创建
	retMap := make(map[string]string)
	retMap["category"] = "receive"
	category := retMap["category"]
	if category != "receive"{
		fmt.Print("check different!")
	}else{
		fmt.Print("check the same!")
	}

	fmt.Println("\n")

	if !strings.EqualFold(category,"receive"){
		fmt.Print("check different using equalfold!")
	}else{
		fmt.Print("check the same using equalfold !")
	}

	fmt.Println("\n")


}