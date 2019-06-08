pub fn part1_input() -> usize {
    part1(&input())
}

pub fn part2_input() -> usize {
    part2(&input())
}

const NICE: bool = true;
const NAUGHTY: bool = !NICE;

fn input() -> String {
    std::fs::read_to_string("testdata/day5.txt").expect("cannot read day 5")
}

fn part1(text: &str) -> usize {
    text.lines().filter(|line| is_nice_part1(line)).count()
}

fn part2(text: &str) -> usize {
    text.lines().filter(|line| is_nice_part2(line)).count()
}

fn is_nice_part1(s: &str) -> bool {
    p1(s) && p2(s) && p3(s)
}

fn is_nice_part2(s: &str) -> bool {
    p4(s) && p5(s)
}

// does not contain ab, cd, pq, or xy
fn p1(s: &str) -> bool {
    let sl = s.to_ascii_lowercase();
    !(sl.contains("ab") || sl.contains("cd") || sl.contains("pq") || sl.contains("xy"))
}

// at least three vowels
fn p2(s: &str) -> bool {
    let mut n = 0;
    for c in s.chars() {
        match c {
            'A' | 'a' => n += 1,
            'E' | 'e' => n += 1,
            'I' | 'i' => n += 1,
            'O' | 'o' => n += 1,
            'U' | 'u' => n += 1,
            _ => {}
        }
    }
    n >= 3
}

// at least one letter that appears twice in a row
fn p3(s: &str) -> bool {
    let buf = s.as_bytes();
    for i in 1..buf.len() {
        if buf[i - 1] == buf[i] {
            return NICE;
        }
    }
    NAUGHTY
}

// a pair of any two letters that appears at least twice in the string without overlapping
fn p4(s: &str) -> bool {
    if s.len() < 4 {
        return false;
    }
    for i in 0..s.len() - 3 {
        let pair = &s[i..i + 2];
        let right = &s[i + 2..s.len()];
        if right.contains(pair) {
            return NICE;
        }
    }
    NAUGHTY
}

// at least one letter which repeats with exactly one letter between them
fn p5(s: &str) -> bool {
    if s.len() < 3 {
        return false;
    }
    let buf = s.as_bytes();
    for i in 1..buf.len() - 1 {
        if buf[i - 1] == buf[i + 1] {
            return NICE;
        }
    }
    NAUGHTY
}

#[cfg(test)]
mod tests {

    #[test]
    fn part1_example1() {
        assert_eq!(super::NICE, super::is_nice_part1("ugknbfddgicrmopn"));
    }

    #[test]
    fn part1_example2() {
        assert_eq!(super::NICE, super::is_nice_part1("aaa"));
    }

    #[test]
    fn part1_example3() {
        assert_eq!(super::NAUGHTY, super::is_nice_part1("jchzalrnumimnmhp"));
    }

    #[test]
    fn part1_example4() {
        assert_eq!(super::NAUGHTY, super::is_nice_part1("haegwjzuvuyypxyu"));
    }

    #[test]
    fn part1_example5() {
        assert_eq!(super::NAUGHTY, super::is_nice_part1("dvszwmarrgswjxmb"));
    }

    #[test]
    fn part1_input() {
        assert_eq!(238, super::part1_input())
    }

    #[test]
    fn part2_p4() {
        assert_eq!(super::NICE, super::p4("xyxy"));
        assert_eq!(super::NICE, super::p4("aabcdefgaa"));
        assert_eq!(super::NAUGHTY, super::p4("aaa"));
    }

    // the initial implementation considers  these three strings naughty, but
    // they are nice
    #[test]
    fn part2_missing_p4() {
        assert_eq!(super::NICE, super::is_nice_part2("komgvqejojpnykol"));
        assert_eq!(super::NICE, super::is_nice_part2("wkkypomlvyglpfpf"));
        assert_eq!(super::NICE, super::is_nice_part2("xojwroydfeoqupup"));
    }

    #[test]
    fn part2_p5() {
        assert_eq!(super::NICE, super::p5("xyxy"));
        assert_eq!(super::NICE, super::p5("abcdefeghi"));
        assert_eq!(super::NICE, super::p5("aaa"));
    }

    #[test]
    fn part2_examples() {
        assert_eq!(super::NICE, super::is_nice_part2("qjhvhtzxzqqjkmpb"));
        assert_eq!(super::NICE, super::is_nice_part2("xxyxx"));
        assert_eq!(super::NAUGHTY, super::is_nice_part2("uurcxstgmygtbstg"));
        assert_eq!(super::NAUGHTY, super::is_nice_part2("ieodomkazucvgmuy"));
    }

    #[test]
    fn part2_input() {
        assert_eq!(69, super::part2_input())
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
