package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Go runs on %s\n", runtime.GOOS)

	// switch os := runtime.GOOS; os {
	// case "darwin":
	// 	fmt.Println("OS X.")
	// case "linux":
	// 	fmt.Println("Linux. ")
	// default:
	// 	fmt.Print("%s.", os)
	// }
}
