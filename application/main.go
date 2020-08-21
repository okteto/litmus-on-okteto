package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Starting hello-chaos server...")
	http.HandleFunc("/", helloServer)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func helloServer(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "Hello world from %s!", hostname)
}
