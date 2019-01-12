package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

type Mahasiswa struct {
	nama string
	umur int
}

func main() {
	mahasiswa := Mahasiswa{"oky saputra", 25}

	fmt.Println(mahasiswa)

}
