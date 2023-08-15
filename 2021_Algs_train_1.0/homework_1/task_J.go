package main

import "fmt"

func main() {
	var a, b, c, d, e, f float32
	fmt.Scanf("%f", &a)
	fmt.Scanf("%f", &b)
	fmt.Scanf("%f", &c)
	fmt.Scanf("%f", &d)
	fmt.Scanf("%f", &e)
	fmt.Scanf("%f", &f)
	// система имеет единственное решение
	// метод крамера
	det := (a * d) - (b * c)
	detx := (e * d) - (b * f)
	dety := (a * f) - (e * c)
	if det != 0 {
		fmt.Println(2, detx/det, dety/det)
	}
}
