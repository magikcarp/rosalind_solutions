package main

import (
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	dna, _ := reader.ReadString('\n')
	var a, c, g, t int

	for _, nuc := range dna {
		switch nuc {
		case 'A':
			a += 1
		case 'C':
			c += 1
		case 'G':
			g += 1
		case 'T':
			t += 1
		}
	}

	println(a, c, g, t)
}
