package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run grph.go <fasta_file>")
		os.Exit(1)
	}

	seqs, err := ParseFasta(os.Args[1])
	if err != nil {
		fmt.Println("Error reading fasta file", err)
		os.Exit(1)
	}

	for i, s := range seqs {
		for j, t := range seqs {
			if i == j {
				continue
			}

			if isIncident(s.Seq, t.Seq, 3) {
				fmt.Println(s.ID, t.ID)
			}
		}
	}
}

type FastaRecord struct {
	ID  string
	Seq string
}

func ParseFasta(filePath string) ([]FastaRecord, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var records []FastaRecord
	var currentID string
	var currentSeq strings.Builder

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, ">") {
			// new sequence, process previous
			if currentID != "" {
				records = append(records, FastaRecord{
					ID:  currentID,
					Seq: currentSeq.String(),
				})
				currentSeq.Reset()
			}
			currentID = strings.TrimPrefix(line, ">")
		} else {
			currentSeq.WriteString(line)
		}
	}

	// add the last record
	if currentID != "" {
		records = append(records, FastaRecord{
			ID:  currentID,
			Seq: currentSeq.String(),
		})
	}

	return records, scanner.Err()
}

func isIncident(s, t string, k int) bool {
	if s[len(s)-k:] == t[:k] {
		return true
	} else {
		return false
	}
}
