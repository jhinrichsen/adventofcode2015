use std::fs;

pub fn part1_input() -> isize {
    part1(&input())
}

pub fn part2_input() -> usize {
    part2(&input())
}

fn input() -> String {
    fs::read_to_string("testdata/day1.txt").expect("failed to read day 1")
}

fn part1(s: &str) -> isize {
    let mut floor = 0;
    for c in s.chars() {
        if c == '(' {
            floor += 1;
        } else {
            floor -= 1;
        }
    }
    floor
}

fn part2(s: &str) -> usize {
    let mut floor = 0;
    for (i, c) in s.chars().enumerate() {
        if c == '(' {
            floor += 1;
        } else {
            floor -= 1;
        }
        if floor == -1 {
            // result must be 1-based
            return i + 1;
        }
    }
    0
}

#[cfg(test)]
mod tests {
    #[test]
    fn part1_examples() {
        let f = super::part1;
        assert_eq!(f("(())"), 0);
        assert_eq!(f("()()"), 0);
        assert_eq!(f("((("), 3);
        assert_eq!(f("(()(()("), 3);
        assert_eq!(f("))((((("), 3);
        assert_eq!(f("())"), -1);
        assert_eq!(f("))("), -1);
        assert_eq!(f(")))"), -3);
        assert_eq!(f(")())())"), -3);
    }

    #[test]
    fn part1_input() {
        assert_eq!(232, super::part1_input());
    }

    #[test]
    fn part2_examples() {
        assert_eq!(super::part2(")"), 1);
        assert_eq!(super::part2("()())"), 5);
    }

    #[test]
    fn part2_input() {
        assert_eq!(1783, super::part2_input());
    }
}
