package main

import "fmt"

func banyakNilai(x, y int) (int, int) {
	x = x * 2
	y = y * 2
	return x, y
}

func main() {
	var foo, bar int

	foo = 10
	bar = 20
	fmt.Println(banyakNilai(foo, bar))
}
