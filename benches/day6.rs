#![feature(test)]
extern crate test;

#[cfg(test)]
mod benches {
    use adventofcode2015::day6::{part1_input, part2_input};

    #[bench]
    fn part1(b: &mut test::Bencher) {
        b.iter(part1_input);
    }

    #[bench]
    fn part2(b: &mut test::Bencher) {
        b.iter(part2_input);
    }
}
