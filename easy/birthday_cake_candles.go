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
 * Complete the 'birthdayCakeCandles' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY candles as parameter.
 */

func findHigherCandle(candles []int32) int32 {
	var higherCandle int32 = 0
	for _, v := range candles {
		if v > higherCandle {
			higherCandle = v
		}
	}

	return higherCandle
}

func birthdayCakeCandles(candles []int32) int32 {
	higherCandle := findHigherCandle(candles)
	var candleCount int32
	for _, v := range candles {
		if v == higherCandle {
			candleCount += 1
		}
	}

	return candleCount
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	candlesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	candlesTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var candles []int32

	for i := 0; i < int(candlesCount); i++ {
		candlesItemTemp, err := strconv.ParseInt(candlesTemp[i], 10, 64)
		checkError(err)
		candlesItem := int32(candlesItemTemp)
		candles = append(candles, candlesItem)
	}

	result := birthdayCakeCandles(candles)

	fmt.Println(result)
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
