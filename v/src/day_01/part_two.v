module day_01

fn part_two(lines []string) int {
	mut total := 0
	for line in lines {
		total += get_calibration_value_two(line)
	}
	return total
}

fn get_calibration_value_two(code string) int {
	str_arr := ['zero', 'one', 'two', 'three', 'four', 'five', 'six', 'seven', 'eight', 'nine']

	mut first_pos := code.len_utf8()
	mut last_pos := -1

	mut digits := [0, 0]

	for i, word in str_arr {
		first_i := code.index(word) or { -1 }
		last_i := code.index_last(word) or { -1 }

		if first_i != -1 {
			first_pos, last_pos = update(first_pos, last_pos, mut digits, first_i, i)

			if first_i != last_i {
				first_pos, last_pos = update(first_pos, last_pos, mut digits, last_i,
					i)
			}
		}
	}

	for i, r in code {
		if !r.is_digit() {
			continue
		}
		first_pos, last_pos = update(first_pos, last_pos, mut digits, i, r.ascii_str().int())
	}

	return digits[0] * 10 + digits[1]
}

fn update(first_pos int, last_pos int, mut digits []int, index int, value int) (int, int) {
	mut new_first := first_pos
	mut new_last := last_pos

	if index < first_pos {
		new_first = index
		digits[0] = value
	}
	if index > last_pos {
		new_last = index
		digits[1] = value
	}
	
	return new_first, new_last
}
