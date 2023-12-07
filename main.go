package main

import (
	dayone "adventofcode/src/dayOne"
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	dayone.PartOne()

	fmt.Println("RUN TIME - ")
	fmt.Println(time.Since(start))
}
