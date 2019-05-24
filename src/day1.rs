fn day1_part1(s: &str) -> isize {
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

fn day1_part2(s: &str) -> usize {
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
    use std::fs::File;
    use std::io::Read;

    #[test]
    fn day1_part1_examples() {
        let f = super::day1_part1;
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
    fn day1_part1() {
        let mut buf = String::new();
        let mut f = File::open("testdata/day1.txt").unwrap();
        f.read_to_string(&mut buf).unwrap();
        assert_eq!(232, super::day1_part1(&buf));
    }

    #[test]
    fn day1_part2_examples() {
        assert_eq!(super::day1_part2(")"), 1);
        assert_eq!(super::day1_part2("()())"), 5);
    }

    #[test]
    fn day1_part2() {
        let mut buf = String::new();
        let mut f = File::open("testdata/day1.txt").unwrap();
        f.read_to_string(&mut buf).unwrap();
        assert_eq!(1783, super::day1_part2(&buf));
    }
}

#[cfg(test)]
mod benches {
    use std::fs::File;
    use std::io::Read;
    use test::Bencher;

    #[bench]
    fn day1_part1(b: &mut Bencher) {
        let mut buf = String::new();
        let mut f = File::open("testdata/day1.txt").unwrap();
        f.read_to_string(&mut buf).unwrap();
        b.iter(|| super::day1_part1(&buf));
    }

    #[bench]
    fn day1_part2(b: &mut Bencher) {
        let mut buf = String::new();
        let mut f = File::open("testdata/day1.txt").unwrap();
        f.read_to_string(&mut buf).unwrap();
        b.iter(|| super::day1_part2(&buf));
    }
}
