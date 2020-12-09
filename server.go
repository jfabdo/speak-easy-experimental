package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jfabdo/fork-talk/src/api"
)

func main() {
	http.HandleFunc("/", api.Index)
	// http.HandleFunc("", api.Messaging)
	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
