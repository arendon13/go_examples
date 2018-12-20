package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/goodbye/", goodbye)
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", nil)
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Inigo Montoya"
	}
	fmt.Fprintf(res, "Hello, my name is %s", name)
}

func goodbye(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	path := req.URL.Path
	parts := strings.Split(path, "/")
	name := parts[2]
	if name == "" {
		name = "Inigo Montoya"
	}
	fmt.Fprintf(res, "Goodbye %s", name)
}

func homePage(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	fmt.Fprintf(res, "The homepage.")
}
