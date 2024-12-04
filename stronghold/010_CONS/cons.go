package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run cons.go <fasta_file>")
		os.Exit(1)
	}

	seqs, err := ParseFasta(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file: ", err)
		os.Exit(1)
	}

	bases := map[rune]int{
		'A': 0,
		'C': 1,
		'G': 2,
		'T': 3,
	}

	ncm := NewCountMatrix(bases, len(seqs[0]))

	for _, seq := range seqs {
		for i, nuc := range seq {
			ncm.UpdateCount(nuc, i)
		}
	}

	fmt.Println(ncm.ConsensusSequence())
	fmt.Println(ncm.String())
}

type CountMatrix struct {
	BaseMap map[rune]int
	Data    [][]int
}

func NewCountMatrix(bm map[rune]int, seqLength int) *CountMatrix {
	matrix := make([][]int, len(bm)*seqLength)
	for i := 0; i < len(bm); i++ {
		matrix[i] = make([]int, seqLength)
	}

	return &CountMatrix{
		BaseMap: bm,
		Data:    matrix,
	}
}

func (cm *CountMatrix) UpdateCount(base rune, location int) {
	cm.Data[cm.BaseMap[base]][location]++
}

func (cm CountMatrix) ConsensusSequence() string {
	reverseBaseMap := map[int]rune{}
	for base, val := range cm.BaseMap {
		reverseBaseMap[val] = base
	}

	consensus := make([]rune, len(cm.Data[0]))

	for j := 0; j < len(cm.Data[0]); j++ {
		maxCount := 0
		maxIndex := -1

		for i := 0; i < len(cm.BaseMap); i++ {
			if cm.Data[i][j] > maxCount {
				maxCount = cm.Data[i][j]
				maxIndex = i
			}
		}

		consensus[j] = reverseBaseMap[maxIndex]
	}

	return string(consensus)
}

func (cm CountMatrix) String() string {
	var outTable string
	bases := make([]rune, 0)
	for k := range cm.BaseMap {
		bases = append(bases, k)
	}
	sort.Slice(bases, func(i, j int) bool {
		return bases[i] < bases[j]
	})

	for _, base := range bases {
		outTable += string(base) + ": "
		for j := 0; j < len(cm.Data[cm.BaseMap[base]]); j++ {
			outTable += strconv.Itoa(cm.Data[cm.BaseMap[base]][j]) + " "
		}
		outTable += "\n"
	}

	return outTable
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
