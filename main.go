package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world! " + strconv.Itoa((int)(time.Now().Unix()))))
}

func file(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Query().Get("path")
	fmt.Printf("path=%v", path)
	content, _ := os.ReadFile(path)
	w.Write([]byte(content))
}

func main() {
	fmt.Println("Start server")
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/file", file)
	err := http.ListenAndServe(":18080", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
