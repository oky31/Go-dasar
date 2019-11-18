package main

import "fmt"

type camera struct {
	merk   string
	series string
	pixel  int
	lensa  string
}

func (cam camera) ambilGambar() {
	fmt.Printf("get picture with pixel = %v \n", cam.pixel)
}

func main() {
	cam := camera{
		"Canon",
		"60D",
		12,
		"Fix",
	}
	fmt.Println(cam)
	cam.ambilGambar()
}
