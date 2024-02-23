package day_01

import (
	"fmt"
	"strings"
)

func Solution(input string) {
	lines := strings.Split(input, "\r\n")

	fmt.Println("Part 1:", partOne(lines))
	fmt.Println("Part 2:", partTwo(lines))
}
