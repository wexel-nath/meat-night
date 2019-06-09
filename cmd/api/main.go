package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/wexel-nath/meat-night/pkg/api"
)

func main() {
	startServer()
}

func getListenAddress() string {
	port := os.Getenv("PORT")

	if len(port) == 0 {
		log.Fatal("$PORT must be set")
	}

	return ":" + port
}

func startServer() {
	address := getListenAddress()
	fmt.Println("Listening on " + address)
	log.Fatal(http.ListenAndServe(address, api.GetRouter()))
}
