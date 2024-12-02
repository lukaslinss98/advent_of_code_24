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

	resultPart1 := part1(sliceCol1, sliceCol2)
	fmt.Printf("result part 1 %d", resultPart1)

	resultPart2 := part2(sliceCol1, sliceCol2)
	fmt.Printf("result part 2 %d", resultPart2)
}

func part1(sliceCol1 []int, sliceCol2 []int) int {
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
	return sum
}

func part2(col1 []int, col2 []int) int {
	dict := make(map[int]int)
	for _, v := range col1 {
		occ := findOccurances(v, col2)
		dict[v] = occ
	}
	sum := 0
	for _, v := range col1 {
		sum += v * dict[v]
	}
	return sum
}

func findOccurances(num int, arr []int) int {
	occurances := 0
	for _, v := range arr {
		if v == num {
			occurances++
		}
	}
	return occurances
}
func convertToInt64(splitArray string) int64 {
	num, _ := strconv.ParseInt(splitArray, 10, 64)
	return num
}
