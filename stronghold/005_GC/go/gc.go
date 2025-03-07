package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type FastaRecord struct {
	ID       string
	Sequence string
}

type GCContentRecord struct {
	ID        string
	GCContent float32
}

// calculate the GC content of a sequence
func CalcGCContent(sequence string) float32 {
	var gc int
	for _, char := range sequence {
		if char == 'G' || char == 'C' {
			gc++
		}
	}
	return float32(gc) / float32(len(sequence)) * 100
}

func ParseFasta(scanner *bufio.Scanner) ([]FastaRecord, error) {
	var records []FastaRecord
	var currentID string
	var currentSeq strings.Builder

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, ">") {
			// new sequence, process previous
			if currentID != "" {
				records = append(records, FastaRecord{
					ID:       currentID,
					Sequence: currentSeq.String(),
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
			ID:       currentID,
			Sequence: currentSeq.String(),
		})
	}

	return records, scanner.Err()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	records, err := ParseFasta(scanner)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading FASTA file: %v\n", err)
		os.Exit(1)
	}

	gcChannel := make(chan GCContentRecord)

	for _, record := range records {
		record := record
		go func() {
			gcContent := CalcGCContent(record.Sequence)
			gcChannel <- GCContentRecord{
				ID:        record.ID,
				GCContent: gcContent,
			}
		}()
	}
	highestGCContent := GCContentRecord{}
	for range records {
		result := <-gcChannel
		if result.GCContent > highestGCContent.GCContent {
			highestGCContent = result
		}
	}

	fmt.Println(highestGCContent.ID)
	fmt.Println(highestGCContent.GCContent)
}
