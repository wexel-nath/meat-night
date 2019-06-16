package main

import (
	"fmt"
	"github.com/wexel-nath/meat-night/pkg/initialize"
	"log"
	"net/http"

	"github.com/wexel-nath/meat-night/pkg/api"
	"github.com/wexel-nath/meat-night/pkg/config"
)

func main() {
	config.Configure()
	initialize.MaybeInsertDinners()

	startServer()
}

func getListenAddress() string {
	port := config.GetPort()

	if len(port) == 0 {
		log.Fatal("PORT must be set")
	}

	return ":" + port
}

func startServer() {
	address := getListenAddress()
	fmt.Println("Listening on " + address)
	log.Fatal(http.ListenAndServe(address, api.GetRouter()))
}
