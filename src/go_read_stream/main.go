package main

import (
	"fmt"
	"net/http"

	"go_read_stream/handler"
)

var port = "8080"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.Top)
	mux.HandleFunc("/data", handler.Data)

	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), mux); err != nil {
		fmt.Printf("server error occured. err: %v", err)
	}
}
