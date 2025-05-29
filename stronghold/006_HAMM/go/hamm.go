package main

import (
	"fmt"
	"os"
	"strings"
)

func HammingDistance(s1, s2 string) int {
	var hd int
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			hd++
		}
	}
	return hd
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run hamm.go <fasta_file>")
		os.Exit(1)
	}

	seqs, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}
	seq_arr := strings.Split(string(seqs), "\n")

	fmt.Println(HammingDistance(seq_arr[0], seq_arr[1]))
}
