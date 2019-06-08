use std::fs;

pub fn part1_input() -> usize {
    part1(&input())
}

fn input() -> String {
    fs::read_to_string("testdata/day6.txt").expect("error reading input for day 6")
}

fn part1(text: &str) -> usize {
    let mut l: Lights = Lights::new();
    for instruction in text.lines() {
        l.execute(&instruction)
    }
    l.count(ON)
}

const WIDTH: usize = 1000;
const HEIGHT: usize = 1000;

const ON: bool = true;
const OFF: bool = false;

#[derive(Debug)]
pub struct Grid {
    from: (usize, usize),
    through: (usize, usize),
}

pub struct Lights {
    // bit-vect is external, so go for underlying type bool
    lights: [[bool; WIDTH]; HEIGHT],
}

impl Default for Lights {
    fn default() -> Self {
        Self::new()
    }
}

pub fn f_on(_: bool) -> bool {
    ON
}

pub fn f_off(_: bool) -> bool {
    OFF
}

pub fn f_toggle(status: bool) -> bool {
    !status
}

// Convert string coordinate to tupel"500,500" to (500,500)
fn atot(coordinates: &str) -> (usize, usize) {
    let mut pi = coordinates.split(",");
    let x = pi.next().expect("missing x value");
    let x = usize::from_str_radix(x, 10).expect("cannot convert x");
    let y = pi.next().expect("missing y value");
    let y = usize::from_str_radix(y, 10).expect("cannot convert y");
    (x, y)
}

// No function types? type Mutator fn(bool) -> bool;

impl Lights {
    pub fn new() -> Self {
        Lights {
            lights: [[false; WIDTH]; HEIGHT],
        }
    }

    pub fn on(&mut self, g: Grid) {
        self.set(g, f_on);
    }

    pub fn off(&mut self, g: Grid) {
        self.set(g, f_off);
    }
    pub fn toggle(&mut self, g: Grid) {
        self.set(g, f_toggle);
    }

    pub fn count(&self, status: bool) -> usize {
        let mut n: usize = 0;
        for x in 0..WIDTH {
            for y in 0..HEIGHT {
                if self.lights[x][y] == status {
                    n += 1;
                }
            }
        }
        n
    }

    pub fn set(&mut self, g: Grid, f: fn(bool) -> bool) {
        // grids are [x..y[
        for x in g.from.0..g.through.0 + 1 {
            for y in g.from.1..g.through.1 + 1 {
                self.lights[x][y] = f(self.lights[x][y]);
            }
        }
    }

    // "turn on 0,0 through 999,999"
    // "toggle 0,0 through 999,0"
    // "turn off 499,499 through 500,500"
    pub fn execute(&mut self, instruction: &str) {
        let f: fn(bool) -> bool;
        let mut pi = instruction.split_ascii_whitespace();
        let op1 = pi.next().expect("missing operation #1");
        if op1 == "toggle" {
            f = f_toggle;
            println!("instruction: toggle");
        } else {
            let op2 = pi.next().expect("missing operation #2");
            if op2 == "on" {
                f = f_on;
                println!("instruction: turn on");
            } else {
                f = f_off;
                println!("instruction: turn off");
            }
        }
        let from = pi.next().expect("missing from value");
        let from = atot(from);
        pi.next().expect("missing through keyword");
        let through = pi.next().expect("missing through value");
        let through = atot(through);

        let g = Grid { from, through };
        println!("before {}", self.count(ON));
        println!("executing {:?}", g);
        self.set(g, f);
        println!("after {}", self.count(ON));
    }
}

#[cfg(test)]
mod tests {
    use super::{part1, Grid, Lights, HEIGHT, WIDTH};

    #[test]
    fn part1_input() {
        assert_eq!(400410, super::part1_input());
    }

    #[test]
    fn atot() {
        assert_eq!((0, 0), super::atot("0,0"));
        assert_eq!((499, 500), super::atot("499,500"));
    }

    #[test]
    fn part1_example1() {
        assert_eq!(WIDTH * HEIGHT, part1("turn on 0,0 through 999,999"));
    }

    #[test]
    fn part1_example2() {
        assert_eq!(1000, part1("toggle 0,0 through 999,0"));
    }

    #[test]
    fn part1_example3() {
        assert_eq!(0, part1("turn off 499,499 through 500,500"));
    }

    #[test]
    fn part1_example4() {
        assert_eq!(
            WIDTH * HEIGHT - 4,
            part1("turn on 0,0 through 999,999\nturn off 499,499 through 500,500")
        );
    }

    #[test]
    fn inclusive() {
        let g = Grid {
            from: (0, 0),
            through: (2, 2),
        };
        let mut l = Lights::new();
        l.on(g);
        assert_eq!(9, l.count(super::ON));
    }
}

#[cfg(test)]
mod benches {
    #[bench]
    fn part1_input(b: &mut test::Bencher) {
        b.iter(|| super::part1_input());
    }
}
