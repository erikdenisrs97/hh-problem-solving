package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	n, err := strconv.ParseInt(strings.TrimSpace(readLine(in)), 10, 64)
	checkError(err)

	arTemp := strings.Split(strings.TrimSpace(readLine(in)), " ")

	var ar []int32

	for i := 0; i < int(n); i++ {
		arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
		checkError(err)
		arItem := int32(arItemTemp)
		ar = append(ar, arItem)
	}

	result := simpleArraySum(ar)

	fmt.Println(result)

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func simpleArraySum(ar []int32) int32 {
	var result int32
	result = 0

	for _, v := range ar {
		result += v
	}

	return result
}
