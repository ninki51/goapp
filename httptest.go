package main

import (
	"net/http"
	"regexp"
	"fmt"
)

type MyHandle func(http.ResponseWriter, *http.Request)

type BaseHttp interface {
	filter()
	connect()
}

type MyRouter struct {
	pattern string
	handler MyHandle
}

func (myRouter MyRouter) filter()  {
	fmt.Println("过滤http请求")
}

func (myRouter MyRouter) connect()  {
	fmt.Println("http请求")
	http.HandleFunc(myRouter.pattern, MyHandle(myRouter.handler))
}

func main(){
	var b BaseHttp
	b = MyRouter{"/", route}
	b.filter()
	b.connect()
	//http.HandleFunc("/", route)
	http.HandleFunc("/api", route)
	http.ListenAndServe(":8080", nil)
}

var num = regexp.MustCompile(`\d`)
var str = regexp.MustCompile(`\w`)

func route(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	switch {
	case num.MatchString(r.URL.Path):
		digits(w, r)
	case str.MatchString(r.URL.Path):
		paramStr(w, r)
	default:
		w.Write([]byte("位置匹配项"))
	}
}

func digits(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("receive digits"))
}

func paramStr(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("receive string"))
}
