package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	dna, _ := reader.ReadString('\n')

	out_dna := strings.Replace(dna, "T", "U", -1)

	println(out_dna)
}
