package main

import (
	"fmt"
	"github.com/beego/beego/v2/client/httplib"
	"github.com/beego/beego/v2/core/logs"
)

func request(url string) string {
	req := httplib.Get(url)

	str, err := req.String()
	if err != nil {
		logs.Info(err)
	}
	return str
}


func main(){
	url := "http://beego.me/"
	str := request(url)
	fmt.Println(str)
}