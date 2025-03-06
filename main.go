//A simple HTTP server in golang

package main

import (
	"fmt"
	"log"
	"net/http"
)

//port of server

const portNum string = ":8080"

// some handler function , which handles the request and process them
// Fprintf prints according to format specifier , writes in w

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage")
}

func Info(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Information Page")
}

func main() {
	log.Println("Starting the server")

	//calling / register the handler function and making paths

	http.HandleFunc("/", Home)
	http.HandleFunc("/info", Info)

	//logging
	log.Println("Started on port", portNum)
	fmt.Println("To close connection CTRL+C :-)")

	//spinning the server
	// log.Fatal means immediate termination of the program

	err := http.ListenAndServe(portNum, nil)
	if err != nil {
		log.Fatal(err)
	}
}
