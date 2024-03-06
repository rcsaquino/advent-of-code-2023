use std::fs;
mod solutions;

fn main() {
    solutions::day_01::print_answers(fs::read_to_string("inputs/day_01").unwrap());
}
