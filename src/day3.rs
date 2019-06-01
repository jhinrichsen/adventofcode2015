pub fn part1_input() -> usize {
    part1(&input())
}

pub fn part2_input() -> usize {
    part2(&input())
}

fn input() -> String {
    std::fs::read_to_string("testdata/day3.txt").expect("cannot read day3")
}

fn part1(s: &str) -> usize {
    // for a tuple of 2 integers, HashSet is 25% faster than BTreeSet
    let mut houses = std::collections::HashSet::new();
    let mut pos = (0, 0);
    // deliver to current position
    houses.insert(pos);
    for c in s.chars() {
        // `move` is a keyword
        pos = mv(pos, c);
        houses.insert(pos);
    }
    houses.len()
}

fn part2(s: &str) -> usize {
    let mut houses = std::collections::HashSet::new();
    // index santa = 0, robo santa = 1
    let mut poss = [(0, 0), (0, 0)];
    // santa starts
    let mut who = 0;
    // deliver to current position
    houses.insert(poss[who]);
    for c in s.chars() {
        // `move` is a keyword
        poss[who] = mv(poss[who], c);
        houses.insert(poss[who]);
        // take turns
        who = 1 - who;
    }
    houses.len()
}

fn mv(pos: (isize, isize), c: char) -> (isize, isize) {
    match c {
        '>' => (pos.0 + 1, pos.1),
        'v' => (pos.0, pos.1 + 1),
        '<' => (pos.0 - 1, pos.1),
        '^' => (pos.0, pos.1 - 1),
        // ignore noise in radio transmission
        _ => (pos.0, pos.1),
    }
}

#[cfg(test)]
mod tests {
    #[test]
    fn part1_examples() {
        assert_eq!(2, super::part1(">"));
        assert_eq!(4, super::part1("^>v<"));
        assert_eq!(2, super::part1("^v^v^v^v^v"));
    }

    #[test]
    fn part1_input() {
        assert_eq!(2081, super::part1_input());
    }

    #[test]
    fn part2_examples() {
        assert_eq!(3, super::part2("^v"));
        assert_eq!(3, super::part2("^>v<"));
        assert_eq!(11, super::part2("^v^v^v^v^v"));
    }

    #[test]
    fn part2_input() {
        assert_eq!(2341, super::part2_input());
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
