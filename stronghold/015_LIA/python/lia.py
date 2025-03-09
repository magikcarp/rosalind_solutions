#!/usr/bin/env python3
# lia.py

import sys

def factorial(n: int) -> int:
    result = 1
    for i in range(2, n+1):
        result *= i
    return result

def choose(n: int, k: int) -> int:
    return (factorial(n) / (factorial(k) * factorial(n-k)))

def binomial_probability(n: int, k: int, p: float) -> float:
    return choose(n, k) * (p ** k) * ((1 - p) ** (n - k))

def lia(kth_generation: int, at_least: int) -> float:
    total_prob = 0
    pop_size = 2 ** kth_generation

    for i in range(at_least, pop_size+1):
        total_prob += binomial_probability(pop_size, i, 0.25)
    
    return total_prob

def main():
    for line in sys.stdin:
        nums = [int(i) for i in line.strip().split(" ")]
        result = lia(nums[0], nums[1])
        print(f"{result:.3f}")

if __name__ == '__main__':
    main()
