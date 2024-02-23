package day_02

import (
	"strconv"
	"strings"
)

func partOne(lines []string) int {
	total := 0

	for _, line := range lines {
		id, sets := parseInfo(line)

		if isPossible(sets) {
			total += id
		}

	}
	return total
}

func isPossible(sets []string) bool {
	const redCubes = 12
	const greenCubes = 13
	const blueCubes = 14

	for _, set := range sets {
		cubes := strings.Split(set, ", ")

		for _, cube := range cubes {
			data := strings.Split(cube, " ")

			count, err := strconv.Atoi(data[0])
			if err != nil {
				panic(err)
			}
			color := data[1]

			if color == "red" && redCubes < count {
				return false
			}
			if color == "green" && greenCubes < count {
				return false
			}
			if color == "blue" && blueCubes < count {
				return false
			}
		}
	}

	return true
}
