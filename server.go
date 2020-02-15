package main

import (
	"hello/api"
	"log"
	"net/http"
)

func hello() string {
    return "hello"
}

func main() {
	http.HandleFunc("/", api.HelloWorld)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
