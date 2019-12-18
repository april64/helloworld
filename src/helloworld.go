package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("hello world!")
	server := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := server.ListenAndServe()
	fmt.Println(err)
}
