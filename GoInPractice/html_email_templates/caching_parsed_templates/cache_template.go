package main

import (
	"html/template"
	"net/http"
)

var t = template.Must(template.ParseFiles("../simple.html"))

// Page organizes the custom variables in the html template
type Page struct {
	Title, Content string
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "An Example",
		Content: "Have fun stormin' da castle.",
	}
	t.Execute(w, p)
}

// technique 33
func main() {
	http.HandleFunc("/", displayPage)
	http.ListenAndServe(":8080", nil)
}
