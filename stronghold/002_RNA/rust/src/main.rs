use std::io;

fn translate(seq: &str) -> String {
    seq.chars().map(|n| if n == 'T' { 'U' } else { n }).collect()
}

fn main() {
    let mut s = String::new();
    io::stdin().read_line(&mut s)
        .ok()
        .expect("Failed to read line.");

    println!("{}", translate(&s))
}
