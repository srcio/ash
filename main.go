package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	pathBase string
	port     string
	version  string
)

func init() {
	pathBase = os.Getenv("ASH_PATH_BASE")
	if pathBase != "" && !strings.HasPrefix(pathBase, "/") {
		pathBase = "/" + pathBase
	}
	port = os.Getenv("ASH_PORT")
	if port == "" {
		port = "80"
	}
	version = os.Getenv("ASH_VERSION")
	if version == "" {
		version = "v1"
	}
}

func headers(w http.ResponseWriter, r *http.Request) {
	var headers string
	for k, v := range r.Header {
		headers += fmt.Sprintf("%v: %v\n", k, v)
	}
	fmt.Fprintf(w, "Headers \n"+headers)
}

func greet(w http.ResponseWriter, r *http.Request) {
	host, _ := os.Hostname()
	if host == "" {
		host = "-"
	}
	log.Println("GET [200] /")
	fmt.Fprintf(w, "Hello World! \nTime now is: %v\nServer: %s\nash version: %s\n", time.Now().Format(time.RFC3339), host, version)
}

func main() {
	rootPath := pathBase
	if pathBase == "" {
		rootPath = "/"
	}

	http.HandleFunc(pathBase+"/headers", headers)
	http.HandleFunc(rootPath, greet)
	log.Println("ash server started.")
	log.Fatalln(http.ListenAndServe(":"+port, nil))
}
