module day_01

fn part_two(input string) int {
	spelled_numbers := ['zero', 'one', 'two', 'three', 'four', 'five', 'six', 'seven', 'eight',
		'nine']

	mut collection := []int{}

	lines := input.split('\n')

	for line in lines {
		mut digits := [0, 0]
		mut first_digit_index := -1
		mut last_digit_index := -1

		for actual_number, spelled_number in spelled_numbers {
			if line.contains(spelled_number) {
				mut f_index := -1
				for {
					f_index = line.index_after(spelled_number, f_index + 1)
					if f_index == -1 {
						break
					}

					if first_digit_index == -1 || first_digit_index > f_index {
						first_digit_index = f_index
						digits[0] = actual_number
					}
					if last_digit_index == -1 || last_digit_index < f_index {
						last_digit_index = f_index
						digits[1] = actual_number
					}
				}
			}
		}

		for i, c in line {
			if c.is_digit() {
				if first_digit_index == -1 || first_digit_index > i {
					first_digit_index = i
					digits[0] = c.ascii_str().int()
				}
				if last_digit_index == -1 || last_digit_index < i {
					last_digit_index = i
					digits[1] = c.ascii_str().int()
				}
			}
		}

		collection << digits[0] * 10 + digits[1]
	}

	mut sum := 0
	for num in collection {
		sum += num
	}

	return sum
}
