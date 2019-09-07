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

    fn add(&mut self, c: SignalProvider) {
        match c {
            SignalProvider::Wire(id, signal) => {
                self.wires.insert(id, Some(signal));
            }
            SignalProvider::Gate(g) => {
                self.gates.push(g);
            }
            SignalProvider::Value(signal) => {}
        }
    }

    fn step(&mut self) {
        for gate in &self.gates {
            println!("triggering gate {:?}", gate);
            match gate {
                // Unary gates
                Gate::Identity(g) => {
                    let signal = self.wires.get(&g.input);
                    if signal.is_none() {
                        println!("no signal yet on {}", &g.input);
                        continue;
                    }
                    let signal = Some(signal.unwrap().unwrap());
                    self.wires.entry(g.output.clone()).or_insert(signal);
                }
                Gate::Not(g) => {
                    let s1 = self.wires.get(&g.input);
                    if s1.is_none() {
                        println!("no signal yet on {}", &g.input);
                        continue;
                    }
                    let signal = Some(!s1.unwrap().unwrap());
                    self.wires.entry(g.output.clone()).or_insert(signal);
                }
                // Binary gates
                Gate::And(g) => {
                    let s1 = self.wires.get(&g.input1);
                    if s1.is_none() {
                        println!("no signal yet on {:?}", &g.input1);
                        continue;
                    }
                    let s2 = self.wires.get(&g.input2);
                    if s2.is_none() {
                        println!("no signal yet on {:?}", &g.input2);
                        continue;
                    }
                    let signal = Some(s1.unwrap().unwrap() & s2.unwrap().unwrap());
                    self.wires.entry(g.output.clone()).or_insert(signal);
                }
                Gate::Or(g) => {
                    let s1 = self.wires.get(&g.input1);
                    if s1.is_none() {
                        println!("no signal yet on {:?}", &g.input1);
                        continue;
                    }
                    let s2 = self.wires.get(&g.input2);
                    if s2.is_none() {
                        println!("no signal yet on {:?}", &g.input2);
                        continue;
                    }
                    let signal = Some(s1.unwrap().unwrap() | s2.unwrap().unwrap());
                    self.wires.entry(g.output.clone()).or_insert(signal);
                }
                Gate::Lshift(g) => {
                    let s1 = self.wires.get(&g.input);
                    if s1.is_none() {
                        println!("no signal yet on {:?}", &g.input);
                        continue;
                    }
                    let signal = Some(s1.unwrap().unwrap() << g.operand);
                    self.wires.entry(g.output.clone()).or_insert(signal);
                }
                Gate::Rshift(g) => {
                    let s1 = self.wires.get(&g.input);
                    if s1.is_none() {
                        println!("no signal yet on {:?}", &g.input);
                        continue;
                    }
                    let signal = Some(s1.unwrap().unwrap() >> g.operand);
                    self.wires.entry(g.output.clone()).or_insert(signal);
                }
            }
        }
    }
}

// optionally resolve constants
fn res(s: &str) -> Value {
    match Signal::from_str_radix(s, 10) {
        Ok(n) => {
            return Value(n);
        }
        Err(_e) => {
            return Wire(s);
        }
    }
}

fn parse(line: &str) -> SignalProvider {
    let mut c: SignalProvider;
    let mut pi = line.split_ascii_whitespace();
    if line.contains("AND") || line.contains("OR") {
        // x AND y -> d
        let w1 = pi.next().expect("binary w1");
        let op = pi.next().expect("binary: missing op");;
        let w2 = pi.next().expect("binary w2");
        let _ = pi.next();
        let w3 = pi.next().expect("binary w3");
        let bg = BinaryGate {
            input1: res(w1),
            input2: res(w2),
            output: res(w3),
        };

        if op == "AND" {
            c = SignalProvider::Gate(Gate::And(bg));
        } else {
            c = SignalProvider::Gate(Gate::Or(bg));
        }
    } else if line.contains("NOT") {
        // NOT e -> f
        let _ = pi.next();
        let w1 = pi.next().expect("unary w1");
        pi.next();
        let w2 = pi.next().expect("binary w2");
        let ug = UnaryGate {
            input: res(w1),
            output: res(w2),
        };
        c = SignalProvider::Gate(Gate::Not(ug));
    } else if line.contains("SHIFT") {
        // p LSHIFT 2 -> q
        let w1 = pi.next().expect("operand w1");
        let op = pi.next().expect("operand missing");
        let w2 = pi.next().expect("operand w2");
        pi.next();
        let w3 = pi.next().expect("operand w3");
        let og = OperandGate {
            input: res(w1),
            operand: res(w2),
            output: res(w3),
        };

        if op == "LSHIFT" {
            c = SignalProvider::Gate(Gate::Lshift(og));
        } else {
            c = SignalProvider::Gate(Gate::Rshift(og));
        }
    } else {
        // 123 -> x
        println!("converting line {}", line);
        let w1 = pi.next().expect("wire signal");
        pi.next();
        let w2 = pi.next().expect("signal");
        let ug = UnaryGate {
            input: res(w1),
            output: res(w2),
        };
        c = SignalProvider::Gate(Gate::Identity(ug));
    }
    println!("parsed component {:?}", c);
    c
}

type WireId = String;

#[derive(Debug)]
enum Value {
    Constant(Signal),
    Wire(WireId),
}

#[derive(Debug)]
enum SignalProvider {
    Gate(Gate),
    Value,
}

#[derive(Debug)]
enum Gate {
    And(BinaryGate),
    Identity(UnaryGate),
    Not(UnaryGate),
    Or(BinaryGate),
    Lshift(OperandGate),
    Rshift(OperandGate),
}

#[derive(Debug)]
struct UnaryGate {
    input: Value,
    output: Value,
}

#[derive(Debug)]
struct BinaryGate {
    input1: Value,
    input2: Value,
    output: Value,
}

#[derive(Debug)]
struct OperandGate {
    input: Value,
    operand: Value,
    output: Value,
}

fn part1(text: &str) -> HashMap<WireId, Option<Signal>> {
    println!("part1");

    let mut board = Circuit::new();
    for line in text.lines() {
        println!("line: {}", line);
        let c = parse(line);
        board.add(c);
    }
    let want = board.wires.len();
    println!("want {} signals", want);
    loop {
        if board.wires.get("a").is_some() {
            return board.wires;
        }
        println!("got {} signals", board.wires.len());
        println!("stepping...");
        board.step();
    }
}

#[cfg(test)]
mod tests {
    use std::collections::HashMap;

    #[test]
    fn part1_input() {
        println!("part1_input()");
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

    #[test]
    fn part1_example() {
        assert_eq!(part1_example_result(), super::part1(&example()));
    }

    // The example for part 1 suggests that wires appear in order, but the input for part 1 shows
    // this is not the case
    #[test]
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
