package dayone

import (
	fileHelper "adventofcode/src/helper"
	"fmt"
	"strconv"
)

func findNumbers(line string) [2]int {
	numbers := [9]rune{'1', '2', '3', '4', '5', '6', '7', '8', '9'}

	var numbersFind []int

	for i, l := range line {
		for _, r := range numbers {
			if l == r {
				numbersFind = append(numbersFind, i)
				break
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

func PartOne() {
	err, lines := fileHelper.ReadFile("./inputs/d1c1.txt")
	if err != nil {
		fmt.Println(err)
	}

	count := 0

	for _, line := range lines {
		count += processLine(line)
	}

	fmt.Println("-------------------------")
	fmt.Printf("Day One Result %d\n", count)
	fmt.Println("-------------------------")
}
