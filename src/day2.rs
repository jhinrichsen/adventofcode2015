fn day2_part1_size(sizes: &[usize; 3]) -> usize {
    let s1 = sizes[0] * sizes[1];
    let s2 = sizes[1] * sizes[2];
    let s3 = sizes[0] * sizes[2];
    // avoid currying as we only have three parameters
    let m = std::cmp::min;
    2 * (s1 + s2 + s3) + m(m(s1, s2), s3)
}

fn parse(s: &str) -> [usize; 3] {
    let mut parts = s.split('x');
    let x = parts
        .next()
        .unwrap()
        .parse::<usize>()
        .expect("failed to unwrap x");
    let y = parts
        .next()
        .unwrap()
        .parse::<usize>()
        .expect("failed to unwrap y");
    let z = parts
        .next()
        .unwrap()
        .parse::<usize>()
        .expect("failed to unwrap z");
    [x, y, z]
}

// Convert XxYxZ into required square feet using day2_part2.
fn day2_part1(s: &str) -> usize {
    day2_part1_size(&parse(s))
}

// cubic feet of volume of the present
fn day2_part2_ribbon_bow(sizes: &[usize; 3]) -> usize {
    sizes[0] * sizes[1] * sizes[2]
}

// smallest perimeter of any one face
fn day2_part2_ribbon_present(sizes: &[usize; 3]) -> usize {
    let size1 = 2 * (sizes[0] + sizes[1]);
    let size2 = 2 * (sizes[1] + sizes[2]);
    let size3 = 2 * (sizes[0] + sizes[2]);
    let m = std::cmp::min;
    m(m(size1, size2), size3)
}

fn day2_part2_ribbon(sizes: &[usize; 3]) -> usize {
    day2_part2_ribbon_bow(sizes) + day2_part2_ribbon_present(sizes)
}

fn day2_part2(s: &str) -> usize {
    day2_part2_ribbon(&parse(s))
}

#[cfg(test)]
mod tests {
    #[test]
    fn day2_part1_examples() {
        let f = super::day2_part1_size;
        assert_eq!(f(&[2, 3, 4]), 58);
        assert_eq!(f(&[1, 1, 10]), 43);
    }

    #[test]
    fn day2_part1_line() {
        assert_eq!(super::day2_part1("2x3x4"), 58);
    }

    #[test]
    fn day2_part1() {
        let want = 1_598_415;
        let buf = std::fs::read_to_string("testdata/day2.txt").unwrap();
        let got: usize = buf.lines().map(|line| super::day2_part1(line)).sum();
        assert_eq!(want, got);
    }

    #[test]
    fn day2_part2_ribbon_present() {
        let f = super::day2_part2_ribbon_present;
        assert_eq!(f(&[2, 3, 4]), 10);
        assert_eq!(f(&[1, 1, 10]), 4);
    }

    #[test]
    fn day2_part2_ribbon_bow() {
        let f = super::day2_part2_ribbon_bow;
        assert_eq!(f(&[2, 3, 4]), 24);
        assert_eq!(f(&[1, 1, 10]), 10);
    }

    #[test]
    fn day2_part2_ribbon() {
        let f = super::day2_part2_ribbon;
        assert_eq!(f(&[2, 3, 4]), 34);
        assert_eq!(f(&[1, 1, 10]), 14);
    }

    #[test]
    fn day2_part2() {
        let want = 3_812_909;
        let buf = std::fs::read_to_string("testdata/day2.txt").unwrap();
        let got: usize = buf.lines().map(|line| super::day2_part2(line)).sum();
        assert_eq!(want, got);
    }
}
