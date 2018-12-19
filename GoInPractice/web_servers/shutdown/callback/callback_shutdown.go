package main

import (
	"fmt"
	"net/http"
	"os"
)

// This way is not recommended
func main() {
	http.HandleFunc("/shutdown", shutdown)
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", nil)
}

func shutdown(res http.ResponseWriter, req *http.Request) {
	os.Exit(0)
}

func homePage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	fmt.Fprint(res, "The homepage.")
}
