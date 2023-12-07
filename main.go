package main

import (
	fileHelper "adventofcode/src/helper"
	"fmt"
	"strconv"
)

const (
	ONE   = '1'
	TWO   = '2'
	THREE = '3'
	FOUR  = '4'
	FIVE  = '5'
	SIX   = '6'
	SEVEN = '7'
	EIGTH = '8'
	NINE  = '9'
)

func findNumbers(line string) [2]int {
	numbers := [9]rune{ONE, TWO, THREE, FOUR, FIVE, SIX, SEVEN, EIGTH, NINE}

	var numbersFind []int

	for i, l := range line {
		for _, r := range numbers {
			if l == r {
				numbersFind = append(numbersFind, i)
				continue
			}
		}

	}

	return [2]int{numbersFind[0], numbersFind[len(numbersFind)-1]}
}

func processLine(line string) int {

	numbersInLine := findNumbers(line)

	stringNumber := fmt.Sprintf("%c%c", line[numbersInLine[0]], line[numbersInLine[1]])

	intNumber, _ := strconv.Atoi(stringNumber)

	return intNumber

}

func main() {

	err, lines := fileHelper.ReadFile("./inputs/d1c1.txt")

	if err != nil {
		fmt.Println(err)
	}

	count := 0

	for _, line := range lines {
		count += processLine(line)
	}

	fmt.Println(count)

}
