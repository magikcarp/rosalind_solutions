import sys


class FastaSequence:
    def __init__(self, id: str, seq: str):
        self.id = id
        self.seq = seq
        self.gc = self.gc_content()
    
    def gc_content(self):
        gc = 0
        for nuc in self.seq:
            if nuc == "C" or nuc == "G":
                gc += 1
        return (gc / len(self.seq))


def parse_fasta(fa_content: list[str]) -> list[FastaSequence]:
    results = []
    curr_id = ""
    curr_seq = ""
    
    for line in fa_content:
        if line[0] == ">": # new sequence starting
            if curr_id != "":
                results.append(FastaSequence(curr_id, curr_seq))
                curr_seq = ""
            curr_id = line[1:]
        else:
            curr_seq += line

    # add the last sequence
    if curr_id != "":
        results.append(FastaSequence(curr_id, curr_seq))

    return results

def main():
    lines = [line.strip() for line in sys.stdin.readlines()]
    seqs = parse_fasta(lines)

    best = seqs[0]
    for i in range(len(seqs)):
        if seqs[i].gc > best.gc:
            best = seqs[i]

    print(f"{best.id}\n{best.gc * 100}")

if __name__ == "__main__":
    main()
