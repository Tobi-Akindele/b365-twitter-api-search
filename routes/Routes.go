package routes

import (
	"bet365/controllers"
	"net/http"
)

func HandleRequests() {
	http.HandleFunc("/ping", controllers.Ping)
	http.HandleFunc("/twitterbot", controllers.TwitterBot)
}
