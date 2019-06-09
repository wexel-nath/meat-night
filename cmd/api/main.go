package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/wexel-nath/meat-night/pkg/api"
)

func main() {
	startServer()
}

func startServer() {
	address := ":8080"
	fmt.Println("Listening on " + address)
	log.Fatal(http.ListenAndServe(address, api.GetRouter()))
}
