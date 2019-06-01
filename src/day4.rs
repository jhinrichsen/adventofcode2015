pub fn part1_input() -> usize {
    part1(&input())
}

pub fn part2_input() -> usize {
    part2(&input())
}

fn input() -> String {
    "bgvyzdsv".to_string()
}

fn part1(s: &str) -> usize {
    mine(s, 5)
}

fn part2(s: &str) -> usize {
    mine(s, 6)
}

fn mine(s: &str, zero_digits: usize) -> usize {
    let mut n = 0;
    let prefix = "0".repeat(zero_digits);
    while !md5(&format!("{}{}", s, n)).starts_with(&prefix) {
        n += 1;
    }
    n
}

fn md5(s: &str) -> String {
    let digest = md5::compute(s);
    format!("{:x}", digest)
}

#[cfg(test)]
mod tests {
    #[test]
    fn part1_example1() {
        assert_eq!(609_043, super::part1("abcdef"));
        assert!(super::md5("abcdef609043").starts_with("000001dbbfa"));
    }

    #[test]
    fn part1_example2() {
        assert_eq!(1_048_970, super::part1("pqrstuv"));
        assert!(super::md5("pqrstuv1048970").starts_with("000006136ef"));
    }

    #[test]
    fn part1_input() {
        assert_eq!(254_575, super::part1_input())
    }

    #[test]
    fn part2_input() {
        assert_eq!(1_038_736, super::part2_input())
    }
}

#[cfg(test)]
mod benches {
    use test::Bencher;

    #[bench]
    fn part1_input(b: &mut Bencher) {
        b.iter(super::part1_input);
    }

    #[bench]
    fn part2_input(b: &mut Bencher) {
        b.iter(super::part2_input);
    }
}
