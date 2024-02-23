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

	firstDigitPosition := len(code)
	lastDigitPosition := -1

	digits := [2]int{0, 0}

	// gtlbhbjgkrb5sixfivefivetwosix
	for i := range strArr {
		firstIndex := strings.Index(code, strArr[i])
		lastIndex := strings.LastIndex(code, strArr[i])

		if firstIndex == -1 {
			continue
		}
		updateDigits(&firstDigitPosition, &lastDigitPosition, &digits, firstIndex, i)

		if firstIndex == lastIndex {
			continue
		}
		updateDigits(&firstDigitPosition, &lastDigitPosition, &digits, lastIndex, i)
	}

	for i, r := range code {
		if !unicode.IsDigit(r) {
			continue
		}
		updateDigits(&firstDigitPosition, &lastDigitPosition, &digits, i, int(r-'0'))
	}

	return digits[0]*10 + digits[1]
}

func updateDigits(firstDigitPosition *int, lastDigitPosition *int, digits *[2]int, index int, value int) {
	if index < *firstDigitPosition {
		*firstDigitPosition = index
		(*digits)[0] = value
	}
	if index > *lastDigitPosition {
		*lastDigitPosition = index
		(*digits)[1] = value
	}
}
