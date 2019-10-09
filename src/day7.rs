use std::collections::HashMap;

pub fn part1_input() -> Signal {
    let wires = part1(&input());
    *wires.get("a").expect("wire a has no signal")
}

fn input() -> String {
    std::fs::read_to_string("testdata/day7.txt").expect("missing day 7")
}

type Signal = u16;
type Wires = HashMap<String, Signal>;

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
            SignalProvider::Gate(g) => {
                self.gates.push(g);
            }
            SignalProvider::Value => {}
        }
    }

    fn get(&self, v: &Value) -> Option<Signal> {
        match v {
            Value::Constant(c) => Some(*c),
            Value::Wire(id) => self.wires.get(id).cloned(),
        }
    }

    // the circuit is complete if all gates have triggered, and all output
    // signals have been generated.
    fn complete(&mut self) -> bool {
        let mut triggered = 0;
        for gate in self.gates.iter() {
            // there's no way to access the underlying generic type, i.e.
            // UnaryGate or BinaryGate other than match() ing to discrete type
            // luckily, pattern matching is exhaustive.
            let output = match gate {
                Gate::Identity(g) => &g.output,
                Gate::And(g) => &g.output,
                Gate::Or(g) => &g.output,
                Gate::Not(g) => &g.output,
                Gate::Lshift(g) => &g.output,
                Gate::Rshift(g) => &g.output,
            };
            if self.wires.get(output).is_some() {
                triggered += 1;
            }
        }
        triggered == self.gates.len()
    }

    fn step(&mut self) {
        for gate in self.gates.iter() {
            match gate {
                // Unary gates
                Gate::Identity(g) => {
                    let signal = self.get(&g.input);
                    if signal.is_none() {
                        continue;
                    }
                    let signal = signal.unwrap();
                    self.wires.entry(g.output.clone()).or_insert(signal);
                }
                Gate::Not(g) => {
                    let s1 = self.get(&g.input);
                    if s1.is_none() {
                        continue;
                    }
                    let signal = !s1.unwrap();
                    self.wires.entry(g.output.clone()).or_insert(signal);
                }
                // Binary gates
                Gate::And(g) => {
                    let s1 = self.get(&g.input1);
                    if s1.is_none() {
                        continue;
                    }
                    let s2 = self.get(&g.input2);
                    if s2.is_none() {
                        continue;
                    }
                    let signal = s1.unwrap() & s2.unwrap();
                    self.wires.entry(g.output.clone()).or_insert(signal);
                }
                Gate::Or(g) => {
                    let s1 = self.get(&g.input1);
                    if s1.is_none() {
                        continue;
                    }
                    let s2 = self.get(&g.input2);
                    if s2.is_none() {
                        continue;
                    }
                    let signal = s1.unwrap() | s2.unwrap();
                    self.wires.entry(g.output.clone()).or_insert(signal);
                }
                Gate::Lshift(g) => {
                    let s1 = self.get(&g.input1);
                    if s1.is_none() {
                        continue;
                    }
                    let s2 = self.get(&g.input2);
                    if s2.is_none() {
                        continue;
                    }
                    let signal = s1.unwrap() << s2.unwrap();
                    self.wires.entry(g.output.clone()).or_insert(signal);
                }
                Gate::Rshift(g) => {
                    let s1 = self.get(&g.input1);
                    if s1.is_none() {
                        continue;
                    }
                    let s2 = self.get(&g.input2);
                    if s2.is_none() {
                        continue;
                    }
                    let signal = s1.unwrap() >> s2.unwrap();
                    self.wires.entry(g.output.clone()).or_insert(signal);
                    // triggered = true;
                }
            }
        }
    }
}

// optionally resolve constants
fn res(s: &str) -> Value {
    match Signal::from_str_radix(s, 10) {
        Ok(n) => Value::Constant(n),
        Err(_) => Value::Wire(s.to_string()),
    }
}

fn parse(line: &str) -> SignalProvider {
    let c: SignalProvider;
    let mut pi = line.split_ascii_whitespace();
    if line.contains("AND") || line.contains("OR") || line.contains("SHIFT") {
        // x AND y -> d
        let w1 = pi.next().expect("binary w1");
        let op = pi.next().expect("binary: missing op");;
        let w2 = pi.next().expect("binary w2");
        let _ = pi.next();
        let w3 = pi.next().expect("binary w3");
        let bg = BinaryGate {
            input1: res(w1),
            input2: res(w2),
            output: w3.to_string(),
        };

        if op == "AND" {
            c = SignalProvider::Gate(Gate::And(bg));
        } else if op == "OR" {
            c = SignalProvider::Gate(Gate::Or(bg));
        } else if op == "LSHIFT" {
            c = SignalProvider::Gate(Gate::Lshift(bg));
        } else if op == "RSHIFT" {
            c = SignalProvider::Gate(Gate::Rshift(bg));
        } else {
            panic!(format!("unknown gate: {}", op))
        }
    } else if line.contains("NOT") {
        // NOT e -> f
        let _ = pi.next();
        let w1 = pi.next().expect("unary w1");
        pi.next();
        let w2 = pi.next().expect("binary w2");
        let ug = UnaryGate {
            input: res(w1),
            output: w2.to_string(),
        };
        c = SignalProvider::Gate(Gate::Not(ug));
    } else {
        // 123 -> x
        let w1 = pi.next().expect("wire signal");
        pi.next();
        let w2 = pi.next().expect("signal");
        let ug = UnaryGate {
            input: res(w1),
            output: w2.to_string(),
        };
        c = SignalProvider::Gate(Gate::Identity(ug));
    }
    println!("parsed component {:?}", c);
    c
}

type WireId = String;

#[derive(Clone, Debug, PartialEq, Eq, Hash)]
enum Value {
    Constant(Signal),
    Wire(WireId),
}

#[derive(Debug, PartialEq, Eq, Hash)]
enum SignalProvider {
    Gate(Gate),
    Value,
}

#[derive(Debug, PartialEq, Eq, Hash)]
enum Gate {
    And(BinaryGate),
    Identity(UnaryGate),
    Not(UnaryGate),
    Or(BinaryGate),
    Lshift(BinaryGate),
    Rshift(BinaryGate),
}

#[derive(Debug, PartialEq, Eq, Hash)]
struct UnaryGate {
    input: Value,
    output: WireId,
}

#[derive(Debug, PartialEq, Eq, Hash)]
struct BinaryGate {
    input1: Value,
    input2: Value,
    output: WireId,
}

fn part1(text: &str) -> Wires {
    println!("part1");

    let mut board = Circuit::new();
    for line in text.lines() {
        println!("line: {}", line);
        let c = parse(line);
        board.add(c);
    }
    let mut complete = false;
    while !complete {
        board.step();
        complete = board.complete();
    }
    board.wires
}

#[cfg(test)]
mod tests {
    use std::collections::HashMap;

    #[test]
    fn part1_input() {
        assert_eq!(16076, super::part1_input());
    }

    fn example() -> String {
        std::fs::read_to_string("testdata/day7_example.txt").expect("cannot read day 7")
    }

    fn part1_example_result() -> super::Wires {
        let mut r = HashMap::new();
        r.insert("d".to_string(), 72);
        r.insert("e".to_string(), 507);
        r.insert("f".to_string(), 492);
        r.insert("g".to_string(), 114);
        r.insert("h".to_string(), 65412);
        r.insert("i".to_string(), 65079);
        r.insert("x".to_string(), 123);
        r.insert("y".to_string(), 456);
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

    #[test]
    fn hashmap_equality() {
        let mut h1 = HashMap::new();
        h1.insert("a", 1);
        h1.insert("b", 2);
        h1.insert("c", 3);
        let mut h2 = HashMap::new();
        h2.insert("c", 3);
        h2.insert("a", 1);
        h2.insert("b", 2);
        assert_eq!(h1, h2);
    }
}
