package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {

	m := New(httpGetBody)
	url := "http://gopl.io"

	start := time.Now()
	value, err := m.Get(url)
	if err != nil {
		log.Print(err)
	}
	fmt.Printf("%s, %s, %d bytes\n",
		url, time.Since(start), len(value.([]byte)))

	start = time.Now()
	value, err = m.Get(url)
	if err != nil {
		log.Print(err)
	}
	fmt.Printf("%s, %s, %d bytes\n",
		url, time.Since(start), len(value.([]byte)))

}

// Memo caches the results of calling a func
type Memo struct {
	f     Func
	cache map[string]result
}

// Func is the type of the function to memoize
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

// New creates a new reference to Memo
func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

// Get is function that will get a http request
// NOTE: not concurrency-safe
func (memo *Memo) Get(key string) (interface{}, error) {

	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	return res.value, res.err

}

func httpGetBody(url string) (interface{}, error) {

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)

}
