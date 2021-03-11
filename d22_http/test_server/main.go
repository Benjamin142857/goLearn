package main

import (
	"fmt"
	"net/http"
)

type TestHandle struct {
}

func (t TestHandle) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Println(666)
}

func f1(rw http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	n, err := rw.Write([]byte("Fuck you"))
	fmt.Println(n, err)
}

func run() {
	//var th TestHandle
	http.HandleFunc("/f1", f1)
	_ = http.ListenAndServe("127.0.0.1:8899", nil)
}

func main() {

	run()
}
