package main

import (
	"fmt"
	"net/http"
	"time"
)

type timeHandler struct {
	format string
}

func (t *timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(t.format)
	w.Write([]byte("the time is: " + tm))
}

func main() {
	mux := http.NewServeMux()
	// 注册默认的路由，未匹配到的地址，都会被它捕获
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("default page"))
	})
	// /not 路由，会产生 404
	mux.Handle("/not", http.NotFoundHandler())

	// foo 转跳 307，而 foo/ 不会转跳
	mux.Handle("/foo", http.RedirectHandler("http://example.org", http.StatusTemporaryRedirect))

	// 注册自定义 Handler
	t := &timeHandler{time.RFC1123}
	mux.Handle("/time", t)

	fmt.Println("listening ...")
	http.ListenAndServe(":8000", mux)
}
