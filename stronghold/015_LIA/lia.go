package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	validRunCheck()

	fileContents, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file, ", err)
		os.Exit(1)
	}

	inputVals := Map(
		Map(
			strings.Split(string(fileContents), " "),
			strings.TrimSpace,
		),
		StrToInt,
	)

	fmt.Println(LIA(inputVals[0], inputVals[1]))
}

func validRunCheck() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run lcsm.go <fasta_file>")
		os.Exit(1)
	}
}

func Map[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}

func StrToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func Factorial(n int) int {
	f := 1
	for i := n; i > 1; i-- {
		f *= i
	}
	return f
}

func Choose(n, k int) int {
	return Factorial(n) / (Factorial(k) * Factorial(n-k))
}

func BinomialProb(n, k int, p float64) float64 {
	return float64(Choose(n, k)) *
		math.Pow(p, float64(k)) *
		math.Pow((1-p), float64(n-k))
}

func LIA(kthGeneration, N int) float64 {
	var totalProb float64
	popSize := int(math.Pow(2, float64(kthGeneration)))
	fmt.Println(popSize)
	for i := N; i <= popSize; i++ {
		newProb := BinomialProb(popSize, i, 0.25)
		totalProb += newProb
		fmt.Printf("%v Choose %v = %v\n", popSize, i, Choose(popSize, i))
	}
	return totalProb
}
