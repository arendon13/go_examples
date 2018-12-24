package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

// technique 20
func main() {
	var file io.ReadCloser
	file, err := OpenCSV("data.csv")
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	defer file.Close()

	// Do something with file
}

// OpenCSV will open a CSV file and prepare it for
// pre-processing
func OpenCSV(filename string) (file *os.File, err error) {
	defer func() {
		// this will set err before return to main()
		if r := recover(); r != nil {
			file.Close()
			err = r.(error)
		}
	}()
	file, err = os.Open(filename)
	if err != nil {
		fmt.Printf("Failed to open file\n")
		return file, err
	}
	RemoveEmptyLines(file)

	return file, err
}

// RemoveEmptyLines will strip CSV of any empty lines in file
func RemoveEmptyLines(f *os.File) {
	panic(errors.New("Failed parse"))
}
