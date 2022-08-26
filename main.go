package main

import (
	"fmt"
	"gee/gee"
	"net/http"
)

func main() {

	r := gee.New()
	r.Get("/", getHandle)
	r.POST("/hello", postHandle)
	r.Run(":8888")
}

func getHandle(w http.ResponseWriter, req *http.Request) {
	fmt.Println("get,path=", req.URL.Path)
	fmt.Fprintln(w, "get 200 ok,path=", req.URL.Path)
}

func postHandle(w http.ResponseWriter, req *http.Request) {
	fmt.Println("post,path=", req.URL.Path)
	fmt.Fprintln(w, "post 200 ok,path=", req.URL.Path)
}
