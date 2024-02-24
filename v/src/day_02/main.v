module day_02

pub fn solution(input string) {
	lines := input.split('\r\n')

	println('Part 1: ${part_one(lines)}')
	println('Part 2: ${part_two(lines)}')
}

fn parse_info(line string) (int, []string) {
	id_and_sets := line.split(': ')

	id := id_and_sets[0].split(' ')[1].int()
	sets := id_and_sets[1].split('; ')

	return id, sets
}
