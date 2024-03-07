import os
import day_01
import day_02

fn main() {
	day_01.print_answers(os.read_file('inputs/day_01')!)
	day_02.print_answers(os.read_file('inputs/day_02')!)
}
