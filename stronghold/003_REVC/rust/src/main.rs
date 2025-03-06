use std::{char, io};

fn complement(n: char) -> char {
    match n {
        'A' => 'T',
        'C' => 'G',
        'G' => 'C',
        'T' => 'A',
        _   => 'N',
    }
}

fn reverse_complement(seq: &str) -> String {
    seq.trim().chars().rev().map(|n| complement(n)).collect()
}

fn main() {
    let mut s = String::new();
    io::stdin().read_line(&mut s)
        .ok()
        .expect("Failed to read line.");
    
    println!("{}", reverse_complement(&s))
}
