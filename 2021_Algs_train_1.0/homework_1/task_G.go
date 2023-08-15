package main

import "fmt"

func main() {
	var n,k,m int
	fmt.Scanf("%d %d %d", &n, &k, &m)
	if !(n >= k && k >= m) {
		fmt.Println(0)
		return
	}
	ost := 0
	d := 0
	for n >= k {
		p := n / k
		n %= k
		ost = p * (k % m)
		d += (p*k-ost) / m
		n += ost
	}
	fmt.Println(d)
}