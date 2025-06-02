package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Mendel(k, m, n float32) float32 {

	tot := k + m + n
	twoHomRec := (n / tot) * ((n - 1) / (tot - 1))
	twoHet := (m / tot) * ((m - 1) / (tot - 1))
	hetRec := ((n / tot) * (m / (tot - 1))) + ((m / tot) * (n / (tot - 1)))
	recessiveProb := twoHomRec + (twoHet * 0.25) + (hetRec * 0.5)

	fmt.Println(twoHomRec)
	fmt.Println(twoHet)
	fmt.Println(hetRec)
	fmt.Println(recessiveProb)

	return 1 - recessiveProb
}

func main() {

	// usage
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run iprb.go <input_file>")
		os.Exit(1)
	}

	// read rosalind input file
	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("error in reading file", err)
		os.Exit(1)
	}
	nums := strings.Split(string(file), " ")

	k, err := strconv.ParseFloat(strings.TrimSpace(nums[0]), 64)
	if err != nil {
		fmt.Println("error in converting k to float32", err)
		os.Exit(1)
	}

	m, err := strconv.ParseFloat(strings.TrimSpace(nums[1]), 32)
	if err != nil {
		fmt.Println("error in converting m to float32", err)
		os.Exit(1)
	}
	n, err := strconv.ParseFloat(strings.TrimSpace(nums[2]), 32)
	if err != nil {
		fmt.Println("error in converting n to float32", err)
		os.Exit(1)
	}

	h := float32(k)
	i := float32(m)
	j := float32(n)

	fmt.Println(h, i, j)
	fmt.Println(Mendel(h, i, j))
}
