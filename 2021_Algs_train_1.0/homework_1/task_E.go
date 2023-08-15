package main

import (
	"fmt"
)

func AppendIfMissing(slice []int, i int) []int {
    for _, ele := range slice {
        if ele == i {
            return slice
        }
    }
    return append(slice, i)
}

func get_p_n(k int, m int, apparts int) (a int, b int) {
	p := (k-1)/(apparts*m) + 1
	n := (k - (p - 1) * apparts * m - 1) / apparts + 1
	return p, n
}

func solve(k1 int, m int, k2 int, p2 int, n2 int) (p int, n int) {
	results_p := []int{}
	results_n := []int{}
	apparts_floor := 1
	for apparts_floor * (m * (p2 - 1) + (n2 - 1)) < k2 && apparts_floor < 100000 {
		p, n := get_p_n(k2, m, apparts_floor)
		if p == p2 && n == n2 {
			p1, n1 := get_p_n(k1, m, apparts_floor)
			results_p = AppendIfMissing(results_p, p1)
			results_n = AppendIfMissing(results_n, n1)
		}
		apparts_floor++
	}
	if len(results_p) == 0 {
		return -1, -1
	}
	var a,b int
	if len(results_p) > 1 {
		a = 0
	} else {
		a = results_p[0]
	}
	if m == 1 {
		b = 1
	} else {
		if len(results_n) > 1 {
			b = 0
		} else {
			b = results_n[0]
		}
	}
	return a,b
}

func main()  {
	var k1, m, k2, p2, n2 int
	fmt.Scanf("%d %d %d %d %d", &k1, &m, &k2, &p2, &n2)
	fmt.Println(solve(k1,m,k2,p2,n2))
}