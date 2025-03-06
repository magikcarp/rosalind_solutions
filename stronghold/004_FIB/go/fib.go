package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RabitReproduction(n, k int) int {
	months := make([]int, n)
	months[0] = 1
	months[1] = 1
	for i := 2; i < len(months); i++ {
		months[i] = months[i-1] + (months[i-2] * k)
	}
	return months[n-1]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	nums := strings.Split(string(input), " ")

	n, err := strconv.Atoi(strings.TrimSpace(nums[0]))
	if err != nil {
		println("failed to convert n to int")
		os.Exit(1)
	}
	k, err := strconv.Atoi(strings.TrimSpace(nums[1]))
	if err != nil {
		fmt.Println("failed to convert k to int", err)
		os.Exit(1)
	}

	println(RabitReproduction(n, k))
}
