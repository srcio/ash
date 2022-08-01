package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! \nTime now is: %v\n", time.Now().Format(time.RFC3339))
}

func headers(w http.ResponseWriter, r *http.Request) {
	r.BasicAuth()
	var headers string
	for k, v := range r.Header {
		headers += fmt.Sprintf("%v: %v\n", k, v)
	}

	fmt.Fprintf(w, "Headers: \n%v\n", headers)
}

func main() {
	log.Println("server started.")
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8081", nil)
}
