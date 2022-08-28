package main

import (
	"log"
	"net/http"
)

func main() {

	//HandleFunc registers the handler function for the given pattern("/") in the DefaultServeMux.
	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		log.Println("Home Page!")
	})

	//Another Handler for pattern ("/settings")
	http.HandleFunc("/settings", func(http.ResponseWriter, *http.Request) {
		log.Println("Setting Page!")
	})

	//Run curl -v localhost:9090 in cmd to trigger ListenAndServe
	//ListenAndServe always returns an error, since it only returns when an unexpected error occurs. In order to log that error we wrap the function call with log.Fatal.
	//The handler is typically nil, in which case the DefaultServeMux is used.
	log.Fatal(http.ListenAndServe(":9090", nil))
}
