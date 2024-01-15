package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func sumArray(arr []int32) int64 {
	var sum int64
	for _, v := range arr {
		sum += int64(v)
	}

	return sum
}

func findMinMax(arr []int64) (int64, int64) {
	var min int64 = arr[0]
	var max int64 = arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max

}

func miniMaxSum(arr []int32) {
	var sums []int64
	for i := 0; i < len(arr); i++ {
		arrayToSum := make([]int32, len(arr))
		copy(arrayToSum, arr)
		arrayToSum = append(arrayToSum[:i], arrayToSum[i+1:]...)
		sum := sumArray(arrayToSum)
		sums = append(sums, sum)
	}

	minimum, maximum := findMinMax(sums)

	fmt.Println(minimum, maximum)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
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
