package main

import "fmt"

func banyakNilai(x, y int) (a, b int) {
	a = x * 2
	b = y * 2
	return
}

func main() {
	var foo, bar int

	foo = 10
	bar = 20
	fmt.Println(banyakNilai(foo, bar))
}
