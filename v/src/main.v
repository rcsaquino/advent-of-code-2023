import day_01
import day_02
import os

fn main() {
	day_01.solution(get_input('inputs/day_01'))
	day_02.solution(get_input('inputs/day_02'))
}

fn get_input(path string) string {
	return os.read_file(path) or { panic(err) }
}
