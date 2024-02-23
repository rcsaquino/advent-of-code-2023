package day_02

import (
	"fmt"
	"strconv"
	"strings"
)

func Solution(input string) {
	lines := strings.Split(input, "\r\n")

	fmt.Println("Part 1:", partOne(lines))
	fmt.Println("Part 2:", partTwo(lines))
}

func parseInfo(line string) (int, []string) {
	idAndSets := strings.Split(line, ": ")
	id, err := strconv.Atoi(strings.Split(idAndSets[0], " ")[1])
	if err != nil {
		panic(err)
	}
	sets := strings.Split(idAndSets[1], "; ")
	return id, sets
}
