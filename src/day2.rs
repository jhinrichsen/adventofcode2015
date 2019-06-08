pub fn part1_input() -> usize {
    part1(&input())
}

pub fn part2_input() -> usize {
    part2(&input())
}

fn input() -> String {
    std::fs::read_to_string("testdata/day2.txt").expect("error reading day 2")
}

pub fn part1(text: &str) -> usize {
    text.lines().map(|line| part1_size(&parse(line))).sum()
}

pub fn part2(text: &str) -> usize {
    text.lines().map(|line| part2_ribbon(&parse(line))).sum()
}

fn part1_size(sizes: &[usize; 3]) -> usize {
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

// cubic feet of volume of the present
fn part2_ribbon_bow(sizes: &[usize; 3]) -> usize {
    sizes[0] * sizes[1] * sizes[2]
}

// smallest perimeter of any one face
fn part2_ribbon_present(sizes: &[usize; 3]) -> usize {
    let size1 = 2 * (sizes[0] + sizes[1]);
    let size2 = 2 * (sizes[1] + sizes[2]);
    let size3 = 2 * (sizes[0] + sizes[2]);
    let m = std::cmp::min;
    m(m(size1, size2), size3)
}

fn part2_ribbon(sizes: &[usize; 3]) -> usize {
    part2_ribbon_bow(sizes) + part2_ribbon_present(sizes)
}

#[cfg(test)]
mod tests {
    #[test]
    fn part1_input() {
        assert_eq!(1_598_415, super::part1_input());
    }

    #[test]
    fn part2_input() {
        assert_eq!(3_812_909, super::part2_input());
    }

    #[test]
    fn part1_examples() {
        let f = super::part1_size;
        assert_eq!(f(&[2, 3, 4]), 58);
        assert_eq!(f(&[1, 1, 10]), 43);
    }

    #[test]
    fn part1_line() {
        assert_eq!(super::part1("2x3x4"), 58);
    }

    #[test]
    fn part2_ribbon_present() {
        let f = super::part2_ribbon_present;
        assert_eq!(f(&[2, 3, 4]), 10);
        assert_eq!(f(&[1, 1, 10]), 4);
    }

    #[test]
    fn part2_ribbon_bow() {
        let f = super::part2_ribbon_bow;
        assert_eq!(f(&[2, 3, 4]), 24);
        assert_eq!(f(&[1, 1, 10]), 10);
    }

    #[test]
    fn part2_ribbon() {
        let f = super::part2_ribbon;
        assert_eq!(f(&[2, 3, 4]), 34);
        assert_eq!(f(&[1, 1, 10]), 14);
    }
}
