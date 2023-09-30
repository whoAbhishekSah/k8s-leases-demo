package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s from server %s", r.URL.Path[1:], os.Getenv("pod_name"))
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("serve started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
