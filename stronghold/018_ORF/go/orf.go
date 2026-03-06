package main

import (
	"bufio"
	"fmt"
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

type FastaRecord struct {
	ID       string
	Sequence string
}

func ParseOneFasta(scanner *bufio.Scanner) (FastaRecord, error) {
	var record FastaRecord
	var currentID string
	var currentSeq strings.Builder

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, ">") {
			currentID = strings.TrimPrefix(line, ">")
		} else {
			currentSeq.WriteString(line)
		}
	}

	record = FastaRecord{ID: currentID, Sequence: currentSeq.String()}
	return record, scanner.Err()
}

type CodonTable map[string]string

var DnaCodonTable = CodonTable{
	"TTT": "F", "TTC": "F", "TTA": "L", "TTG": "L",
	"TCT": "S", "TCC": "S", "TCA": "S", "TCG": "S",
	"TAT": "Y", "TAC": "Y", "TAA": "*", "TAG": "*",
	"TGT": "C", "TGC": "C", "TGA": "*", "TGG": "W",

	"CTT": "L", "CTC": "L", "CTA": "L", "CTG": "L",
	"CCT": "P", "CCC": "P", "CCA": "P", "CCG": "P",
	"CAT": "H", "CAC": "H", "CAA": "Q", "CAG": "Q",
	"CGT": "R", "CGC": "R", "CGA": "R", "CGG": "R",

	"ATT": "I", "ATC": "I", "ATA": "I", "ATG": "M",
	"ACT": "T", "ACC": "T", "ACA": "T", "ACG": "T",
	"AAT": "N", "AAC": "N", "AAA": "K", "AAG": "K",
	"AGT": "S", "AGC": "S", "AGA": "R", "AGG": "R",

	"GTT": "V", "GTC": "V", "GTA": "V", "GTG": "V",
	"GCT": "A", "GCC": "A", "GCA": "A", "GCG": "A",
	"GAT": "D", "GAC": "D", "GAA": "E", "GAG": "E",
	"GGT": "G", "GGC": "G", "GGA": "G", "GGG": "G",
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

type Orf struct {
	StartPosition int
	EndPosition   int
}

func FindOrfs(s string) []Orf {
	var orfs []Orf
	s_comp := ReverseComplement(s)
	for i := range len(s) {
		if i < len(s)-6 { // search forward
			if s[i:i+3] != "ATG" {
				continue
			}

		} else if i > 6 { // search reverse complement

		}
	}
	return orfs
}

func main() {
	reader := bufio.NewScanner(os.Stdin)
	fa_entry, err := ParseOneFasta(reader)
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(Translate(fa_entry.Sequence, DnaCodonTable))
}
