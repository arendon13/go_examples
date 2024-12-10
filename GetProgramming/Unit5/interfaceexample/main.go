package main

import "fmt"

// DbAccesor will connect with the database
type DbAccesor interface {
	Save()
	Fetch() string
	Delete()
	Update()
}

func create(db DbAccesor) {
	db.Save()
}

func read(db DbAccesor) {
	value := db.Fetch()
	fmt.Println(value)
}

func update(db DbAccesor) {
	db.Update()
}

func delete(db DbAccesor) {
	db.Delete()
}

type post struct {
	id       string
	platform string
}

var myPost = post{
	id:       "someId",
	platform: "Twitter",
}

func (p post) Save() {
	fmt.Println()
}

func main() {

}
