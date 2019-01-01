package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

// technique 45
func main() {

}

func fileForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("upload.html")
		t.Execute(w, nil)
	} else {
		f, h, err := r.FormFile("file")
		if err != nil {
			panic(err)
		}
		defer f.Close()
		filename := "/tmp/" + h.Filename // location to store file
		out, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
		defer out.Close()
		io.Copy(out, f)
		fmt.Fprint(w, "Upload complete")
	}
}
