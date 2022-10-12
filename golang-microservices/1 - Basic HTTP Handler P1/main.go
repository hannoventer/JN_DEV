package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	//HandleFunc(pattern,function to trigger) | rw & r are shorter var names.
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Home Page!")

		d, err := ioutil.ReadAll(r.Body) //lib that reads all data from request(r)'s body

		if err != nil {
			//htt.Error is a convenient method to return a response message and code to "rw"
			http.Error(rw, "Oops", http.StatusBadRequest)
			return //exit handler function
		}

		//send the request body(r) back to the client via response writer(rw)
		fmt.Fprintf(rw, "You sent the following data: %s\n", d)
	})

	//Another Handler for pattern ("/settings")
	http.HandleFunc("/settings", func(http.ResponseWriter, *http.Request) {
		log.Println("Setting Page!")
	})

	//ListenAndServe always returns an error, it only returns when an unexpected error occurs.
	//In order to log that error we wrap the function call with log.Fatal.
	//The handler is typically nil, in which case the DefaultServeMux is used.
	log.Fatal(http.ListenAndServe(":9090", nil))
}

//Test Code with the following command in cmd (not powershell!):
//Run curl -v -d "MyData" localhost:9090 in cmd to send a HTTP GET request from CMD and trigger ListenAndServe
//-v verbose mode / additional info
//-d states that the following text will be the data to send (request body)
//"MyData" - the data you are sending
