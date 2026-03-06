import std/strutils

var input = system.readLine(system.stdin)

type
  NuclCount = object
    A: int
    C: int
    G: int
    T: int

proc countNucleotides(sequence: string): NuclCount =
  for i in countup(0, sequence.len - 1):
    case sequence[i]
    of 'A':
      result.A += 1
    of 'C':
      result.C += 1
    of 'G':
      result.G += 1
    of 'T':
      result.T += 1
    else:
      continue

proc displayNuclCount(nc: NuclCount): string =
  return "$# $# $# $#" % [$nc.A, $nc.C, $nc.G, $nc.T]

echo displayNuclCount(countNucleotides(input))
