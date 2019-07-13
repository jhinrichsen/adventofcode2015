use std::collections::HashMap;

pub fn part1_input() -> Signal {
    let wires = part1(&input());
    wires.get("a").expect("wire a has no signal").unwrap()
}

fn input() -> String {
    std::fs::read_to_string("testdata/day7.txt").expect("missing day 7")
}

type Signal = u16;
type Wires = HashMap<String, Option<Signal>>;

struct Circuit {
    wires: Wires,
    gates: Vec<Gate>,
}

impl Circuit {
    fn new() -> Circuit {
        Circuit {
            wires: HashMap::new(),
            gates: Vec::new(),
        }
    }

    fn wire(&mut self, c: Component) {
        match c {
            Component::Wire(id, signal) => {
                self.wires.insert(id, Some(signal));
            }
            Component::Gate(g) => {
                self.gates.push(g);
            }
        }
    }

    fn step(&mut self) {
        for gate in &self.gates {
            match gate {
                Gate::And(g) => {
                    let s1 = self.wires.get(&g.input1).expect("missing input 1 wire");
                    if s1.is_none() {
                        println!("no signal yet on {}", &g.input1);
                        continue;
                    }
                    let s2 = self.wires.get(&g.input2).expect("missing input 2 wire");
                    if s2.is_none() {
                        println!("no signal yet on {}", &g.input2);
                        continue;
                    }
                    let signal = Some(s1.unwrap() & s2.unwrap());
                    // self.wires.entry(g.output.clone()).or_insert(signal);
                    self.wires.entry(g.output.clone()).or_insert(signal);
                }
                _ => panic!("missing impl"),
            }
        }
    }
}

fn parse(s: &str) -> Component {
    let mut c: Component;
    let mut pi = s.split_ascii_whitespace();
    if s.contains("AND") || s.contains("OR") {
        // x AND y -> d
        let w1 = pi.next().expect("binary w1");
        let op = pi.next().expect("binary: missing op");;
        let w2 = pi.next().expect("binary w2");
        let _ = pi.next();
        let w3 = pi.next().expect("binary w3");
        let bg = BinaryGate {
            input1: w1.to_string(),
            input2: w2.to_string(),
            output: w3.to_string(),
        };

        if op == "AND" {
            c = Component::Gate(Gate::And(bg));
        } else {
            c = Component::Gate(Gate::Or(bg));
        }
    } else if s.contains("NOT") {
        // NOT e -> f
        let _ = pi.next();
        let w1 = pi.next().expect("unary w1");
        let w2 = pi.next().expect("binary w2");
        let ug = UnaryGate {
            input: w1.to_string(),
            output: w2.to_string(),
        };
        c = Component::Gate(Gate::Not(ug));
    } else if s.contains("SHIFT") {
        // p LSHIFT 2 -> q
        let w1 = pi.next().expect("operand w1");
        let op = pi.next().expect("operand missing");;
        let w2 = pi.next().expect("operand w2");
        let og = OperandGate {
            input: w1.to_string(),
            operand: Signal::from_str_radix(w2, 10).expect("operand: illegal value"),
            output: w2.to_string(),
        };

        if op == "LSHIFT" {
            c = Component::Gate(Gate::Lshift(og));
        } else {
            c = Component::Gate(Gate::Rshift(og));
        }
    } else {
        // 123 -> x
        let w1 = pi.next().expect("wire signal");
        let _ = pi.next();
        let w2 = pi.next().expect("signal");
        let w2 = Signal::from_str_radix(w2, 10).expect("bad signal");
        c = Component::Wire(w1.to_string(), w2);
    }
    c
}

type WireId = String;

enum Component {
    Wire(String, Signal),
    Gate(Gate),
}

enum Gate {
    And(BinaryGate),
    Not(UnaryGate),
    Or(BinaryGate),
    Lshift(OperandGate),
    Rshift(OperandGate),
}

struct UnaryGate {
    input: WireId,
    output: WireId,
}

struct BinaryGate {
    input1: WireId,
    input2: WireId,
    output: WireId,
}

struct OperandGate {
    input: WireId,
    operand: Signal,
    output: WireId,
}

fn part1(text: &str) -> HashMap<WireId, Option<Signal>> {
    let mut board = Circuit::new();
    for line in text.lines() {
        let c = parse(line);
        board.wire(c);
    }
    loop {
        if board.wires.get("a").is_some() {
            return board.wires;
        }
        board.step();
    }
}

#[cfg(test)]
mod tests {
    use std::collections::HashMap;

    // #[test]
    fn part1_input() {
        assert_eq!(42, super::part1_input());
    }

    fn example() -> String {
        std::fs::read_to_string("testdata/day7_example.txt").expect("cannot read day 7")
    }

    fn part1_example_result() -> super::Wires {
        let mut r = HashMap::new();
        r.insert("d".to_string(), Some(72));
        r.insert("e".to_string(), Some(507));
        r.insert("f".to_string(), Some(492));
        r.insert("g".to_string(), Some(114));
        r.insert("h".to_string(), Some(65412));
        r.insert("i".to_string(), Some(65079));
        r.insert("x".to_string(), Some(123));
        r.insert("y".to_string(), Some(456));
        r
    }

    // #[test]
    fn part1_example() {
        assert_eq!(part1_example_result(), super::part1(&example()));
    }

    // The example for part 1 suggests that wires appear in order, but the input for part 1 shows
    // this is not the case
    // #[test]
    fn part1_example_rev() {
        let text = &example()
            .lines()
            .rev()
            // retrofit the newline info
            .map(|line| format!("{}\n", line))
            .collect::<String>();
        assert_eq!(part1_example_result(), super::part1(text));
    }
}
