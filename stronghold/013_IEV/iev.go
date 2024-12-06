package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
1. AA-AA
2. AA-Aa
3. AA-aa
4. Aa-Aa
5. Aa-aa
6. aa-aa
*/

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run iev.go <input_file>")
		os.Exit(1)
	}

	fileContents, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file: ", err)
	}

	popPairs := Map(
		Map(
			strings.Split(string(fileContents), " "),
			strings.TrimSpace,
		),
		StrToInt,
	)

	fmt.Println(IEV(popPairs))
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

func ToFloat32(i int) float32 {
	return float32(i)
}

func IEV(h []int) float32 {
	g := Map(h, ToFloat32)
	return float32((g[0] * 2) + (g[1] * 2) + (g[2] * 2) + (g[3] * 1.5) + (g[4]))
}
