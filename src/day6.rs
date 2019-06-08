// max return value is WIDTH*HEIGHT = 1_000_000
pub fn part1_input() -> u32 {
    part1(&input())
}

// max return value is WIDTH * HEIGHT * sizeof(u8) = 256_000_000
pub fn part2_input() -> u32 {
    part2(&input())
}

fn input() -> String {
    std::fs::read_to_string("testdata/day6.txt").expect("error reading input for day 6")
}

fn part1(text: &str) -> u32 {
    using(text, IS_PART_1)
}

fn part2(text: &str) -> u32 {
    using(text, IS_PART_2)
}

fn using(text: &str, is: InstructionSet) -> u32 {
    let mut l: Lights = Lights::new();
    for instruction in text.lines() {
        l.execute(&instruction, is)
    }
    l.count()
}

// Underlying type that supports both part 1 (on/off) and part 2 (brightness)
// maximum brightness is undefined, so number range is a guess
type Light = u8;

const ON: Light = 1;
const OFF: Light = 0;

type Mutator = fn(Light) -> Light;

// our instructin sets know on, off and toggle
type InstructionSet = (Mutator, Mutator, Mutator);
const IS_PART_1: InstructionSet = (f_on, f_off, f_toggle);
const IS_PART_2: InstructionSet = (f_inc, f_dec, f_inc2);

fn f_on(_: Light) -> Light {
    ON
}

fn f_off(_: Light) -> Light {
    OFF
}

fn f_toggle(l: Light) -> Light {
    1 - l
}

fn f_inc(l: Light) -> Light {
    if l == 255 {
        panic!("overflow in inc");
    }
    l + 1
}

fn f_inc2(l: Light) -> Light {
    if l > 253 {
        panic!("overflow in inc2");
    }
    l + 2
}

fn f_dec(l: Light) -> Light {
    if l == 0 {
        return 0;
    }
    l - 1
}

#[derive(Debug)]
pub struct Grid {
    from: (usize, usize),
    through: (usize, usize),
}

const WIDTH: usize = 1000;
const HEIGHT: usize = 1000;

pub struct Lights {
    // bit-vect is external, so go for underlying type bool
    lights: [[Light; WIDTH]; HEIGHT],
}

impl Default for Lights {
    fn default() -> Self {
        Self::new()
    }
}

// Convert "500,500" to (500,500)
fn atot(coordinates: &str) -> (usize, usize) {
    let mut pi = coordinates.split(',');
    let x = pi.next().expect("missing x value");
    let x = usize::from_str_radix(x, 10).expect("cannot convert x");
    let y = pi.next().expect("missing y value");
    let y = usize::from_str_radix(y, 10).expect("cannot convert y");
    (x, y)
}

impl Lights {
    pub fn new() -> Self {
        Lights {
            // lights start off
            lights: [[OFF; WIDTH]; HEIGHT],
        }
    }

    // works for both on/off and brightness mode
    fn count(&self) -> u32 {
        let mut n = 0;
        for x in 0..WIDTH {
            for y in 0..HEIGHT {
                n += u32::from(self.lights[x][y])
            }
        }
        n
    }

    // "turn on 0,0 through 999,999"
    // "toggle 0,0 through 999,0"
    // "turn off 499,499 through 500,500"
    pub fn execute(&mut self, instruction: &str, is: InstructionSet) {
        let f: Mutator;
        let mut pi = instruction.split_ascii_whitespace();
        let op1 = pi.next().expect("missing operation #1");
        if op1 == "toggle" {
            f = is.2;
        } else {
            let op2 = pi.next().expect("missing operation #2");
            if op2 == "on" {
                f = is.0;
            } else {
                f = is.1;
            }
        }
        let from = pi.next().expect("missing from value");
        let from = atot(from);
        pi.next().expect("missing through keyword");
        let through = pi.next().expect("missing through value");
        let through = atot(through);

        let g = Grid { from, through };
        self.set(g, f);
    }

    fn set(&mut self, g: Grid, f: Mutator) {
        // grids are [x..y[
        for x in g.from.0..=g.through.0 {
            for y in g.from.1..=g.through.1 {
                self.lights[x][y] = f(self.lights[x][y]);
            }
        }
    }
}

#[cfg(test)]
mod tests {
    use super::part1;

    #[test]
    fn part1_input() {
        assert_eq!(400410, super::part1_input());
    }

    #[test]
    fn part2_input() {
        assert_eq!(15343601, super::part2_input());
    }

    #[test]
    fn atot() {
        assert_eq!((0, 0), super::atot("0,0"));
        assert_eq!((499, 500), super::atot("499,500"));
    }

    #[test]
    fn part1_example1() {
        assert_eq!(1_000_000, part1("turn on 0,0 through 999,999"));
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
            1_000_000 - 4,
            part1("turn on 0,0 through 999,999\nturn off 499,499 through 500,500")
        );
    }
}
