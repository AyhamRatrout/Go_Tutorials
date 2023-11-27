package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", rootMessageHandler)
	http.HandleFunc("/home", homeMessageHandler)

	port := ":8080"
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Printf("Error ocurred trying to run the web service: %s\n", err)
		log.Fatal(err)
	}

}

func rootMessageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, web service world!")
}

func homeMessageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./home.html")
}
