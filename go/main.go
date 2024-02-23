package main

import (
	"advent-of-code-2023/solutions/day_01"
	"advent-of-code-2023/solutions/day_02"
	"fmt"
	"os"
)

func getInput(path string) string {
	res, err := os.ReadFile("inputs/day_01")
	if err != nil {
		panic(err)
	}
	return string(res)
}

func main() {
	fmt.Println("=== DAY 01 ===")
	day_01.Solution(getInput("inputs/day_01"))

	fmt.Println("=== DAY 02 ===")
	day_02.Solution(getInput("inputs/day_02"))
}
