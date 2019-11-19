package main

import "fmt"

func foo(x int) int {
	return x * 2
}

func main() {
	var y int
	y = 5
	fmt.Println(foo(y))
}
