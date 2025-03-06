package main

import (
	"bufio"
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
	reader := bufio.NewReader(os.Stdin)
	dna, _ := reader.ReadString('\n')

	println(ReverseComplement(dna))
}
