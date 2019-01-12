package main

import "fmt"

func main() {
	//nilai := []int{1, 2, 3, 4, 5}

	//nilaiPertama := nilai[0]
	//nilaiPertama = 20

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}

	fmt.Println(s)
}
