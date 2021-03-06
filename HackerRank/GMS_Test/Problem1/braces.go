package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the braces function below.
func braces(values []string) []string {
	vals := []string{
		"yes",
		"no",
	}
	return vals
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	valuesCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	var values []string

	for i := 0; i < int(valuesCount); i++ {
		valuesItem := readLine(reader)
		values = append(values, valuesItem)
	}

	res := braces(values)

	for i, resItem := range res {
		fmt.Fprintf(writer, "%s", resItem)

		if i != len(res)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
