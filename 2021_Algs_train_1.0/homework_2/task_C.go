package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func solve(arr []int, x int) int {
	minDiffs := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		minDiffs[i] = Abs(arr[i] - x)
	}
	
	minDiff := minDiffs[0]
	index := 0
	for i := 0; i < len(arr); i++ {
		if minDiffs[i] < minDiff {
			minDiff = minDiffs[i];
			index = i
		}
	}
	return arr[index]
}

func main() {
	// Create a scanner to read from standard input
	scanner := bufio.NewScanner(os.Stdin)

	// Read the first line containing a single integer
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	// Read the second line containing space-separated integers
	scanner.Scan()
	arrStr := strings.Split(scanner.Text(), " ")
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i], _ = strconv.Atoi(arrStr[i])
	}

	// Read the third line containing a single integer
	scanner.Scan()
	x, _ := strconv.Atoi(scanner.Text())

	fmt.Println(solve(arr, x))
}