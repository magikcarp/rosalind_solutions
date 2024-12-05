package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run fibd.go <input_file>")
		os.Exit(1)
	}
	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error with reading file, ", err)
		os.Exit(1)
	}
	numsAsStr := strings.Split(string(input), " ")

	nMonths, err := strconv.Atoi(strings.TrimSpace(numsAsStr[0]))
	if err != nil {
		println("failed to convert n to int")
		os.Exit(1)
	}
	lifespan, err := strconv.Atoi(strings.TrimSpace(numsAsStr[1]))
	if err != nil {
		println("failed to convert n to int")
		os.Exit(1)
	}

	fmt.Println(mortalRabbitFibonacci(nMonths, lifespan))
}

func mortalRabbitFibonacci(month, lifespan int) int {
	months := make([]int, month)
	months[0] = 1
	months[1] = 1

	for i := 2; i < month; i++ {
		if i < lifespan {
			months[i] = months[i-1] + months[i-2]
		} else if i == lifespan || i == lifespan+1 {
			months[i] = months[i-1] + months[i-2] - 1
		} else {
			months[i] = months[i-1] + months[i-2] - months[i-(lifespan+1)]
		}
	}

	return months[month-1]
}
