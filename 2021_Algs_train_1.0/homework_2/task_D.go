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

func main() {
	arr := readNumsFromLine()
	k := 0
	for i := 1; i < len(arr)-1; i++ {
		if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
			k++
		}
	}
	fmt.Println(k)
}