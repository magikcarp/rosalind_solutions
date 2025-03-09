#!/usr/bin/env python3
# mprt.py

import sys
import re
from re import Pattern # imported for type hinting
from urllib.request import urlopen

def get_uniprot_fasta(id: str) -> str:
    result = urlopen('https://www.uniprot.org/uniprotkb/' + id + '.fasta')
    return result.read().decode('utf-8')

def parse_fasta(fasta: str) -> str:
    result = re.sub(r"^>(.*)\n", "", fasta, 1)
    result = re.sub(r"\n", "", result)
    return result

def find_motif(sequence: str, motif: Pattern) -> list[int]:
    result = []
    hits = re.finditer(motif, sequence)
    for hit in hits:
        result.append(hit.start()+1)
    return result

def main():
    target = re.compile(r"(?=N[A-O,Q-Z](S|T)[A-O,Q-Z])")
    for id in sys.stdin:
        protein_fasta = get_uniprot_fasta(id.strip().split('_')[0])
        sequence = parse_fasta(protein_fasta)
        hits = find_motif(sequence, target)
        if len(hits) > 0:
            print(id.strip())
            print(" ".join([str(x) for x in hits]))

if __name__ == '__main__':
    main()
