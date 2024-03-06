pub fn print_answers(input: String) {
    println!("---DAY 1---");
    solve_part_one(&input);
    solve_part_two(&input);
}

fn solve_part_one(input: &String) {
    let mut digits: Vec<i32> = Vec::new();

    let mut collection: Vec<i32> = Vec::new();

    for (i, c) in input.chars().enumerate() {
        if c.is_digit(10) {
            digits.push(c.to_digit(10).unwrap() as i32);
        }

        if c == '\n' || i == input.len() - 1 {
            let value = (digits.first().unwrap() * 10) + digits.last().unwrap();
            collection.push(value);
            digits.clear()
        }
    }

    println!("Part 1: {}", collection.iter().sum::<i32>());
}

fn solve_part_two(input: &String) {
    let spelled_numbers = [
        "zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
    ];

    let mut collection: Vec<i32> = Vec::new();

    let lines: Vec<&str> = input.split('\n').collect();

    for line in lines {
        let mut digits = [0, 0];
        let mut first_digit_index: i32 = -1;
        let mut last_digit_index: i32 = -1;

        for (actual_number, spelled_number) in spelled_numbers.iter().enumerate() {
            if line.contains(spelled_number) {
                let f_indices: Vec<i32> = line
                    .match_indices(spelled_number)
                    .map(|(i, _)| i as i32)
                    .collect();
                for f_index in f_indices {
                    if first_digit_index == -1 || first_digit_index > f_index {
                        first_digit_index = f_index;
                        digits[0] = actual_number as i32;
                    }
                    if last_digit_index == -1 || last_digit_index < f_index {
                        last_digit_index = f_index;
                        digits[1] = actual_number as i32;
                    }
                }
            }
        }

        for (i, c) in line.chars().enumerate() {
            let index = i as i32;
            if c.is_digit(10) {
                if first_digit_index == -1 || first_digit_index > index {
                    first_digit_index = index;
                    digits[0] = c.to_digit(10).unwrap() as i32;
                }
                if last_digit_index == -1 || last_digit_index < index {
                    last_digit_index = index;
                    digits[1] = c.to_digit(10).unwrap() as i32;
                }
            }
        }

        collection.push((digits[0] * 10) + digits[1]);
    }

    println!("Part 2: {}", collection.iter().sum::<i32>());
}
