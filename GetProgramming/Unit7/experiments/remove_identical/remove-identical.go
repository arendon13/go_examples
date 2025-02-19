package main

import "fmt"

func main() {
	chan1 := make(chan string)
	chan2 := make(chan string)

	go sourceGopher(chan1)
	go removeDuplicatesGopher(chan1, chan2)
	printGopher(chan2)
}

func sourceGopher(downstream chan string) {
	for _, v := range []string{"a", "b", "b", "c", "d", "e", "e", "e", "f", "f"} {
		downstream <- v
	}
	close(downstream)
}

func removeDuplicatesGopher(upstream, downstream chan string) {
	var prev string
	for {
		item, ok := <-upstream
		if !ok {
			close(downstream)
			return
		}

		if item != prev {
			downstream <- item
		}

		prev = item
	}
}

func printGopher(upstream chan string) {
	for v := range upstream {
		fmt.Println(v)
	}
}
