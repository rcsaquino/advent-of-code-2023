package day_01

import (
	"strings"
	"unicode"
)

func partTwo(lines []string) int {
	total := 0
	for _, line := range lines {
		total += getCalibrationValueTwo(line)
	}
	return total
}

func getCalibrationValueTwo(code string) int {
	strArr := [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	firstDigitPos := len(code)
	lastDigitPos := -1

	digits := [2]int{0, 0}

	for i := range strArr {
		firstIndex := strings.Index(code, strArr[i])
		lastIndex := strings.LastIndex(code, strArr[i])

		if firstIndex != -1 {
			updateDigits(&firstDigitPos, &lastDigitPos, &digits, firstIndex, i)

			if firstIndex != lastIndex {
				updateDigits(&firstDigitPos, &lastDigitPos, &digits, lastIndex, i)
			}
		}
	}

	for i, r := range code {
		if !unicode.IsDigit(r) {
			continue
		}
		updateDigits(&firstDigitPos, &lastDigitPos, &digits, i, int(r-'0'))
	}

	return digits[0]*10 + digits[1]
}

func updateDigits(firstDigitPos *int, lastDigitPos *int, digits *[2]int, index int, value int) {
	if index < *firstDigitPos {
		*firstDigitPos = index
		(*digits)[0] = value
	}
	if index > *lastDigitPos {
		*lastDigitPos = index
		(*digits)[1] = value
	}
}
