package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	validRun()

	fastaSeqs, err := ParseFasta(os.Args[1])
	if err != nil {
		fmt.Println("Error parsing FASTA: ", err)
	}

	sort.Strings(fastaSeqs)
	fmt.Println(fastaSeqs)
	fmt.Println(LongestCommonSubstring(fastaSeqs))
}

func validRun() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run lcsm.go <fasta_file>")
		os.Exit(1)
	}
}

func ParseFasta(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var records []string
	var currentSeq strings.Builder

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, ">") {
			// new sequence, process previous
			if currentSeq.Len() > 0 {
				records = append(records, currentSeq.String())
				currentSeq.Reset()
			}
		} else {
			currentSeq.WriteString(line)
		}
	}

	// add the last record
	if currentSeq.Len() > 0 {
		records = append(records, currentSeq.String())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return records, nil
}

func LongestCommonSubstring(seqArray []string) string {
	var lcs string
	shortestSeq := seqArray[0]
	remainingSeqs := seqArray[1:]

	for i := 0; i < len(shortestSeq); i++ {
		for j := len(shortestSeq); j > len(lcs)+i; j-- {
			s1 := seqArray[0][i:j]

			matched_all := true
			for _, s2 := range remainingSeqs {
				if !strings.Contains(s2, s1) {
					matched_all = false
					break
				}
			}

			if matched_all {
				lcs = s1
				break
			}
		}
	}

	return lcs
}
