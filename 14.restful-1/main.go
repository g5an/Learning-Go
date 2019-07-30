package main

import (
	"io"
	"net/http"
)

// 练习1：通过handlefunc方法注册路由
// func HelloServer(w http.ResponseWriter, req *http.Request) {
// 	io.WriteString(w, "hello, world! \n")
// }

// func main() {
// 	http.HandleFunc("/hello", HelloServer)
// 	err := http.ListenAndServe(":8082", nil)
// 	if err != nil {
// 		log.Fatal("ListenAndServe", err)
// 	}
// }

// 练习2：
type abctHot struct{}

func (a *abctHot) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.String() // 获得路径
	io.WriteString(w, path)
}

func main() {
	http.ListenAndServe(":8082", &abctHot{})
}
