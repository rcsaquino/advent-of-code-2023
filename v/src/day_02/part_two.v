module day_02

fn part_two(lines []string) int {
	mut total := 0

	for line in lines {
		_, sets := parse_info(line)
		total += get_power(sets)
	}

	return total
}

fn get_power(sets []string) int {
	mut min_red := 0
	mut min_blue := 0
	mut min_green := 0

	for set in sets {
		cubes := set.split(', ')

		for cube in cubes {
			data := cube.split(' ')

			count := data[0].int()
			color := data[1]

			if color == 'red' && count > min_red {
				min_red = count
			}
			if color == 'green' && count > min_green {
				min_green = count
			}
			if color == 'blue' && count > min_blue {
				min_blue = count
			}
		}
	}
	return min_red * min_green * min_blue
}
