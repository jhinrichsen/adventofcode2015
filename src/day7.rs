use std::collections::HashMap;

type Registers = HashMap<String, u16>;

fn part1(text: &str) -> Registers {
    let mut registers: Registers = HashMap::new();
    for line in text.lines() {
        println!("{}, anyone?", line);
        let mut pi = line.split_ascii_whitespace();
        if line.contains("AND") {
            // x AND y -> d
            let r1 = pi.next().expect("AND r1");
            let v1 = registers.get(r1).expect("AND missing r1");
            let _ = pi.next();
            let r2 = pi.next().expect("AND r2");
            let v2 = registers.get(r2).expect("AND missing r2");
            let _ = pi.next();
            let r3 = pi.next().expect("AND r3");
            let v = v1 & v2;
            registers.insert(r3.to_string(), v);
            println!("{} = {}", r3, v);
        } else if line.contains("OR") {
            // x OR y -> d
            let r1 = pi.next().expect("OR r1");
            let v1 = registers.get(r1).expect("OR missing r1");
            let _ = pi.next();
            let r2 = pi.next().expect("OR r2");
            let v2 = registers.get(r2).expect("OR missing r2");
            let _ = pi.next();
            let r3 = pi.next().expect("OR r3");
            let v = v1 | v2;
            registers.insert(r3.to_string(), v);
            println!("{} = {}", r3, v);
        } else if line.contains("LSHIFT") {
            // x LSHIFT 2 -> d
            let r1 = pi.next().expect("LSHIFT r1");
            let v1 = registers.get(r1).expect("LSHIFT missing r1");
            let _ = pi.next();
            let v2 = pi.next().expect("LSHIFT missing value");
            let v2 = u16::from_str_radix(v2, 10).expect("SET bad number");
            let _ = pi.next();
            let r2 = pi.next().expect("LSHIFT r2");
            let v = v1 << v2;
            registers.insert(r2.to_string(), v);
            println!("{} = {}", r2, v);
        } else if line.contains("RSHIFT") {
            // x RSHIFT 2 -> d
            let r1 = pi.next().expect("RSHIFT r1");
            let v1 = registers.get(r1).expect("RSHIFT missing r1");
            let _ = pi.next();
            let v2 = pi.next().expect("RSHIFT missing value");
            let v2 = u16::from_str_radix(v2, 10).expect("SET bad number");
            let _ = pi.next();
            let r2 = pi.next().expect("RSHIFT r2");
            let v = v1 >> v2;
            registers.insert(r2.to_string(), v);
            println!("{} = {}", r2, v);
        } else if line.contains("NOT") {
            // NOT x -> h
            let _ = pi.next().expect("NOT missing");
            let r1 = pi.next().expect("NOT r1");
            let r1 = registers.get(r1).expect("NOT missing r1");
            let _ = pi.next();
            let r2 = pi.next().expect("NOT r2");
            let v = !r1;
            registers.insert(r2.to_string(), v);
            println!("{} = {}", r2, v);
        } else {
            // 123 -> x
            let s = pi.next().expect("SET missing value");
            let v = u16::from_str_radix(s, 10).expect("SET bad number");
            let _ = pi.next();
            let r1 = pi.next().expect("SET missing r1");
            registers.insert(r1.to_string(), v);
            println!("{} = {}", r1, v);
        }
    }
    registers
}

fn example() -> String {
    std::fs::read_to_string("testdata/day7_example.txt").expect("cannot read day 7")
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1_example() {
        let mut want: Registers = HashMap::new();
        want.insert("d".to_string(), 72);
        want.insert("e".to_string(), 507);
        want.insert("f".to_string(), 492);
        want.insert("g".to_string(), 114);
        want.insert("h".to_string(), 65412);
        want.insert("i".to_string(), 65079);
        want.insert("x".to_string(), 123);
        want.insert("y".to_string(), 456);
        let got = super::part1(&example());
        assert_eq!(want, got);
    }
}
