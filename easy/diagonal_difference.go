package main

import (
	"bufio"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func diagonalDifference(arr [][]int32) int32 {
	var x int32 = 0
	var y int = len(arr) - 1
	var sumX int32
	var sumY int32
	var diff int32
	for _, v := range arr {
		sumX += v[x]
		sumY += v[y]
		x += 1
		y -= 1
	}
	diff = sumX - sumY

	return int32(math.Abs(float64(diff)))
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var arr [][]int32
	for i := 0; i < int(n); i++ {
		arrRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != int(n) {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}

	diagonalDifference(arr)

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
