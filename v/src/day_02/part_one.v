module day_02

fn part_one(lines []string) int {
	mut total := 0

	for line in lines {
		id, sets := parse_info(line)

		if is_possible(sets) {
			total += id
		}
	}

	return total
}

fn is_possible(sets []string) bool {
	red_cubes := 12
	green_cubes := 13
	blue_cubes := 14

	for set in sets {
		cubes := set.split(', ')

		for cube in cubes {
			data := cube.split(' ')

			count := data[0].int()
			color := data[1]

			if color == 'red' && red_cubes < count {
				return false
			}
			if color == 'green' && green_cubes < count {
				return false
			}
			if color == 'blue' && blue_cubes < count {
				return false
			}
		}
	}

	return true
}
