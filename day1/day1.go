package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day1/inputDay1.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var sliceCol1 []int
	var sliceCol2 []int

	for scanner.Scan() {
		line := scanner.Text()
		splitArray := strings.Split(line, "   ")

		num64Col1 := convertToInt64(splitArray[0])
		numCol1 := int(num64Col1)

		num64Col2 := convertToInt64(splitArray[1])
		numCol2 := int(num64Col2)

		sliceCol1 = append(sliceCol1, numCol1)
		sliceCol2 = append(sliceCol2, numCol2)
	}
	sort.Ints(sliceCol1)
	sort.Ints(sliceCol2)

	var sum int
	for i := range sliceCol1 {
		valCol1 := sliceCol1[i]
		valCol2 := sliceCol2[i]
		fmt.Println(valCol1, valCol2)

		if valCol1 >= valCol2 {
			sum += valCol1 - valCol2
		} else {
			sum += valCol2 - valCol1
		}
	}
	fmt.Println(sum)
}

func convertToInt64(splitArray string) int64 {
	num, _ := strconv.ParseInt(splitArray, 10, 64)
	return num
}
