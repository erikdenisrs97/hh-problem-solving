package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func timeConversion(s string) string {
	meridian := s[len(s)-2:]
	hour := s[:2]
	newTime := s[:len(s)-2]

	if hour == "12" && meridian == "PM" {
		return newTime
	} else if hour == "12" && meridian == "AM" {
		newTime = "00" + newTime[2:]
		return newTime
	} else if hour != "12" && meridian == "PM" {
		h, err := strconv.Atoi(newTime[:len(newTime)-6])
		if err != nil {
			return err.Error()
		}
		h += 12
		newTime = strconv.Itoa(h) + newTime[2:]
		return newTime
	} else {
		return newTime
	}

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

	fmt.Fprintf(writer, "%s\n", result)

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
