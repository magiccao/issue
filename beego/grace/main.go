package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/astaxie/beego/grace"
)

var addr string = ":8989"

func main() {
	f, err := os.Open("tmp")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	srv := grace.NewServer(addr, http.HandlerFunc(handler))
	err = srv.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}

func handler(rw http.ResponseWriter, r *http.Request) {
}
