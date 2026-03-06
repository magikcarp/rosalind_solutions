#!/usr/bin/env python3

import sys
from functools import reduce

def calculate_possible_codons(sequence: str) -> int:
    n_codons_per_aa: dict[str, int] = {
        'A': 4,
        'C': 2,
        'D': 2,
        'E': 2,
        'F': 2,
        'G': 4,
        'H': 2,
        'I': 3,
        'K': 2,
        'L': 6,
        'M': 1,
        'N': 2,
        'P': 4,
        'Q': 2,
        'R': 6,
        'S': 6,
        'T': 4,
        'V': 4,
        'W': 1,
        'Y': 2
    }
    return reduce(lambda acc, x: acc * n_codons_per_aa[x], sequence, 3) % 1_000_000

def main():
    protein_sequence = sys.stdin.readline().strip()
    print(calculate_possible_codons(protein_sequence))

if __name__ == "__main__":
    main()
