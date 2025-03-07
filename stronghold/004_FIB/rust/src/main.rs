use std::io;

fn parse_input(line: &str) -> Vec<usize> {
    line
        .trim()
        .split(' ')
        .flat_map(str::parse::<usize>)
        .collect::<Vec<usize>>()
}

fn rabbit_reproduction(n: usize, k: usize) -> usize {
    let mut tmp: [usize; 2] = [1usize, 1usize];
    for _ in 2..n {
        tmp = [tmp[1], tmp[1] + (tmp[0] * k)];
    }
    tmp[1]
}

fn main() {
    let mut s = String::new();
    io::stdin().read_line(&mut s).ok().expect("Failed to read line.");
    let nums = parse_input(&s);

    println!("{}", rabbit_reproduction(nums[0], nums[1]))
}