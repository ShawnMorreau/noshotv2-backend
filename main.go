package main

import (
	"fmt"
	"log"
	"net/http"
)

func Server(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello there")
}

func main() {
	handler := http.HandlerFunc(Server)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
