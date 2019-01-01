package main

import (
	"fmt"
	"net/http"
)

func main() {

}

// technique 43
func exampleFormHandler(w http.ResponseWriter, r *http.Request) {
	// name := r.FormValue("name")
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}
	name := r.FormValue("name")
	fmt.Println(name)
}

// technique 44
func exampleMultiFormHandler(w http.ResponseWriter, r *http.Request) {
	maxMemory := 16 << 20
	err := r.ParseMultipartForm(int64(maxMemory))
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range r.PostForm["names"] {
		fmt.Println(v)
	}
}
