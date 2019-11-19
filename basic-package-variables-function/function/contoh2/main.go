package main

import "fmt"

func swap(x, y string) (string, string, string) {
	z := "kata baru"
	return y, x, z
}

func main() {
	a, b, c := swap("tama", "makan")
	fmt.Println(a, b, c)
}
