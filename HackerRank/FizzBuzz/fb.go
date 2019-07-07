package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	fbCount := getUserInput()

	fizzBuzz(fbCount)

}

func getUserInput() int {

	scanner := bufio.NewScanner(os.Stdin)

	var input string
	if scanner.Scan() {
		input = scanner.Text()
	}
	if scanner.Err() != nil {
		log.Fatal("failed to scan input")
	}

	fbCount, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal("failed to convert to integer")
	}

	return fbCount

}

func fizzBuzz(count int) {

	for i := 1; i <= count; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

}
