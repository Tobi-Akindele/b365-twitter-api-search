package main

import (
	"bet365/config"
	"bet365/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {

	config.LoadAppConfigs()
	serverHost, serverPort := config.GetServerAndPort()

	routes.HandleRequests()

	fmt.Printf("Starting service on [%v%v] \n", serverHost, serverPort)
	log.Fatal(http.ListenAndServe(serverPort, nil))
}
