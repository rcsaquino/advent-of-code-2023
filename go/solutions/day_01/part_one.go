package day_01

import (
	"unicode"
)

func partOne(lines []string) int {
	total := 0
	for _, line := range lines {
		total += getCalibrationValueOne(line)
	}
	return total
}

func getCalibrationValueOne(code string) int {
	digits := []int{}
	for _, r := range string(code) {
		if unicode.IsDigit(r) {
			digits = append(digits, int(r-'0'))
		}
	}
	return digits[0]*10 + digits[len(digits)-1]
}
