package main

import (
	"log"
	"net/http"

	"github.com/mozillazg/web-demo/api/views"
)

func main() {
	host := ""
	port := "8080"
	address := host + ":" + port
	log.Printf("Listen %s\n", address)
	http.HandleFunc("/", views.Handler)
	log.Fatal(http.ListenAndServe(address, nil))
}
