package main

import (
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("rosalind_rna.txt")
	if err != nil {
		println("error in reading file")
		os.Exit(1)
	}
	dna := string(file)
	out_dna := strings.Replace(dna, "T", "U", -1)

	println(out_dna)
}
