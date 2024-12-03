package main

import (
	"os"
	"strings"
)

func ReverseComplement(dna string) string {
	var complement = map[rune]rune{
		'A': 'T',
		'C': 'G',
		'G': 'C',
		'T': 'A',
	}

	var result strings.Builder

	for i := len(dna) - 1; i >= 0; i-- {
		base := rune(dna[i])
		result.WriteRune(complement[base])
	}

	return result.String()
}

func main() {
	file, err := os.ReadFile("rosalind_revc.txt")
	if err != nil {
		println("error in reading file")
		os.Exit(1)
	}
	dna := string(file)
	println(ReverseComplement(dna))
}
