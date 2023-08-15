package main

import (
	"fmt"
	"sort"
)

func main() {
	var a,b,c,d,e int
	fmt.Scanf("%d",&a)
	fmt.Scanf("%d",&b)
	fmt.Scanf("%d",&c)
	fmt.Scanf("%d",&d)
	fmt.Scanf("%d",&e)
	if d > e {
		d, e = e, d
	}
	brick := []int{a,b,c}
	sort.Ints(brick)

	a,b,c = brick[0], brick[1], brick[2]

	if a <= d && b <= e {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}