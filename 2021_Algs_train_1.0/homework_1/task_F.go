package main

import (
	"fmt"
)

func tableSquare(w1,h1,w2,h2 int) (min_w, min_h int) {
	var height int
	if h1 > h2 {
		height = h1
	} else {
		height = h2
	}
	return w1+w2, height
}

func findMinPair(s1,s2,s3,s4 [2]int) (w,h int) {
	minMul := s1[0]*s1[1]
	w, h = s1[0], s1[1]

	pairs := [][2]int{s2, s3, s4}
	for _, pair := range pairs {
		mul := pair[0] * pair[1]
		if mul < minMul {
			minMul = mul
			w, h = pair[0], pair[1]
		}
	}

	return w, h
}

func main()  {
	var w1,h1,w2,h2 int
	fmt.Scanf("%d %d %d %d", &w1, &h1, &w2, &h2)
	rw1, rh1 := tableSquare(w1,h1,w2,h2)
	r1 := [2]int{rw1, rh1}
	rw2, rh2 := tableSquare(h1,w1,w2,h2)
	r2 := [2]int{rw2, rh2}
	rw3, rh3 := tableSquare(h1,w1,h2,w2)
	r3 := [2]int{rw3, rh3}
	rw4, rh4 := tableSquare(w1,h1,h2,w2)
	r4 := [2]int{rw4, rh4}
	w,h := findMinPair(r1,r2,r3,r4)
	fmt.Println(w, h)
}