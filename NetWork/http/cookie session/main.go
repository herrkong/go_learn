package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

/*
读写 Cookie
当客户端第一次访问服务端时，服务端为客户端发一个凭证（全局唯一的字符串）
服务端将字符串发给客户端（写Cookie的过程）
HTTP 请求头（发送给服务端Cookie）
HTTP 响应头（在服务端通知客户端保存Cookie）
*/

// 写 cookie
func writeCookie(w http.ResponseWriter, r *http.Request)  {
	expiration := time.Now()
	expiration = expiration.AddDate(0, 0, 3)
	cookie := http.Cookie{Name:"username", Value:"darwin", Expires:expiration}
	http.SetCookie(w, &cookie)
	fmt.Fprintf(w,"write cookie success")
}

// 读 cookie
func readCookie(w http.ResponseWriter, r *http.Request)  {
	cookie, _ := r.Cookie("username")
	fmt.Fprint(w, cookie)
}
	
func main()  {
	http.HandleFunc("/writeCookie", writeCookie)
	http.HandleFunc("/readCookie", readCookie)

	fmt.Println("服务器已经启动，写 cookie 地址：http://localhost:8800/writeCookie ，读 cookie 地址：http://localhost:8800/readCookie")

	// 启动 HTTP 服务，并监听端口号，开始监听，处理请求，返回响应
	err := http.ListenAndServe(":8800", nil)
	if err != nil {
		log.Fatal("ListenAndServe",err)
	}
}
