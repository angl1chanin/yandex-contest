package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readNumsFromLine() []int {
	// new scanner
	sc := bufio.NewScanner(os.Stdin)
	// scan from stdin
	sc.Scan()
	// get the text from scanner
	line := sc.Text()
	// split the line by spaces
	arr := strings.Split(line, " ")
	intArr := []int{}
	// convert each string number to integer
	for _, v := range arr {
		intV, _ := strconv.Atoi(v)
		intArr = append(intArr, intV)
	}
	return intArr
}

func checkSequence(seq []int) bool {
	for i := 0; i < len(seq)-1; i++ {
		if seq[i] >= seq[i+1] {
			return false
		}
	}
	return true
}

func main() {
	nums := readNumsFromLine()
	res := checkSequence(nums)
	if res {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}