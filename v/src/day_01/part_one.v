module day_01

fn part_one(lines []string) int {
	mut total := 0
	for line in lines {
		total += get_calibration_value_one(line)
	}

	return total
}

fn get_calibration_value_one(code string) int {
	mut digits := []int{}

	for r in code {
		if r.is_digit() {
			digits << r.ascii_str().int()
		}
	}

	return digits[0] * 10 + digits.last()
}
