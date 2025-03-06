import sys

def countNucs(seq: str) -> dict[str, int]:
    out = {}
    for nuc in seq: 
        if nuc not in out.keys():
            out[nuc] = 1
        else:
            out[nuc] += 1
    return out

def printNucCounts(d: dict[str, int]) -> None: 
    print(f"{d['A']} {d['C']} {d['G']} {d['T']}")

def main():
    for line in sys.stdin:
        try:
            printNucCounts(countNucs(line.strip().upper()))
        except Exception as e:
            print(f"{e}\nInvalid input: {line}", file=sys.stderr)

if __name__ == '__main__':
    main()
