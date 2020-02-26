package main

import (
	"os"
	"fmt"
	"net/http"
)

// EchoServer ...
func EchoServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	config := os.Getenv("CONFIG")
	w.Write([]byte(fmt.Sprintf("pong from %s", config)))
}

func main() {
	http.HandleFunc("/ping", EchoServer)
	fmt.Println("Listen at :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}
