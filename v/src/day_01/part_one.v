module day_01

fn part_one(input string) int {
	mut digits := []int{}
	mut collection := []int{}

	for i, c in input {
		if c.is_digit() {
			digits << c.ascii_str().int()
		}

		if c.ascii_str() == '\n' || i == (input.len - 1) {
			value := digits.first() * 10 + digits.last()
			collection << value
			digits.clear()
		}
	}
	mut sum := 0
	for num in collection {
		sum += num
	}
	return sum
}
