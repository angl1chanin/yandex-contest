package main

import (
	"fmt"
	"math"
)

func main() {
	var a,b,n,m int
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	fmt.Scanf("%d", &n)
	fmt.Scanf("%d", &m)
	min1 := n + a*(n-1)
	max1 := n + a*(n-1) + 2*a
	min2 := m + b*(m-1)
	max2 := m + b*(m-1) + 2*b
	min := int(math.Max(float64(min1), float64(min2)))
	max := int(math.Min(float64(max1), float64(max2)))
	if min <= max {
		fmt.Println(min, max)
	} else {
		fmt.Println(-1)
	}
}