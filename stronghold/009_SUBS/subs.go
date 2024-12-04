package main

import (
	"fmt"
	"os"
	"strings"
)

func SubStringLocations(s, t string) []int {
	locs := make([]int, 0)

	for i := 0; i < len(s)-len(t); i++ {
		if s[i:i+len(t)] == t {
			locs = append(locs, i+1)
		}
	}

	return locs
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run subs.go <fasta_file>")
		os.Exit(1)
	}

	seqs, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}
	seq_arr := strings.Split(string(seqs), "\n")

	fmt.Println(SubStringLocations(seq_arr[0], seq_arr[1]))
}
