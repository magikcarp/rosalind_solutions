package main

import (
	"fmt"
	"os"
)

type CodonTable map[string]string

var StdCodonTable = CodonTable{
	"UUU": "F", "UUC": "F", "UUA": "L", "UUG": "L",
	"UCU": "S", "UCC": "S", "UCA": "S", "UCG": "S",
	"UAU": "Y", "UAC": "Y", "UAA": "*", "UAG": "*",
	"UGU": "C", "UGC": "C", "UGA": "*", "UGG": "W",

	"CUU": "L", "CUC": "L", "CUA": "L", "CUG": "L",
	"CCU": "P", "CCC": "P", "CCA": "P", "CCG": "P",
	"CAU": "H", "CAC": "H", "CAA": "Q", "CAG": "Q",
	"CGU": "R", "CGC": "R", "CGA": "R", "CGG": "R",

	"AUU": "I", "AUC": "I", "AUA": "I", "AUG": "M",
	"ACU": "T", "ACC": "T", "ACA": "T", "ACG": "T",
	"AAU": "N", "AAC": "N", "AAA": "K", "AAG": "K",
	"AGU": "S", "AGC": "S", "AGA": "R", "AGG": "R",

	"GUU": "V", "GUC": "V", "GUA": "V", "GUG": "V",
	"GCU": "A", "GCC": "A", "GCA": "A", "GCG": "A",
	"GAU": "D", "GAC": "D", "GAA": "E", "GAG": "E",
	"GGU": "G", "GGC": "G", "GGA": "G", "GGG": "G",
}

func Translate(s string, ct CodonTable) string {
	var protein string
	for i := 0; i < len(s); i += 3 {
		currentCodon := s[i : i+3]
		newAA, ok := ct[currentCodon]
		if !ok || newAA == "*" {
			break
		}
		protein += newAA
	}
	return protein
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run prot.go <input_file>")
	}

	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("error reading file", err)
	}

	rna := string(file)
	fmt.Println(Translate(rna, StdCodonTable))
}
