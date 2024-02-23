package day_02

import (
	"strconv"
	"strings"
)

func partTwo(lines []string) int {
	total := 0

	for _, line := range lines {
		_, sets := parseInfo(line)

		total += getPower(sets)

	}
	return total
}

func getPower(sets []string) int {
	minRed := 0
	minBlue := 0
	minGreen := 0

	for _, set := range sets {
		cubes := strings.Split(set, ", ")

		for _, cube := range cubes {
			data := strings.Split(cube, " ")

			count, err := strconv.Atoi(data[0])
			if err != nil {
				panic(err)
			}

			color := data[1]

			if color == "red" && count > minRed {
				minRed = count
			}
			if color == "green" && count > minGreen {
				minGreen = count
			}
			if color == "blue" && count > minBlue {
				minBlue = count
			}
		}
	}

	return minRed * minBlue * minGreen
}
