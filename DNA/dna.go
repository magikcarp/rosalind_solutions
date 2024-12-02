package main

import (
	"os"
)

func main() {
	file, err := os.ReadFile("rosalind_dna.txt")
	if err != nil {
		println("error reading input file")
		os.Exit(1)
	}
	dna := string(file)
	var a, c, g, t int

	for _, nuc := range dna {
		switch nuc {
		case 'A':
			a += 1
		case 'C':
			c += 1
		case 'G':
			g += 1
		case 'T':
			t += 1
		}
	}

	println(a, c, g, t)
}
