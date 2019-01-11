# Flow Control

## For 
Go hanya memiliki satu looping construct, yaitu `for`   
3 Komponen yang di pisahkan oleh semicolon `;` yaitu :
```
- init state
- condition expression
- post statement
```

`init state` dan `post statement` adalah optional di dalam badan for

## While
Karna go hanya memiliki `for` loop, untuk `while` gunakan `for` tanpa `init state` dan `post statement` 
```
sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
```

## IF
Expressi di `if` statement tidak mengunakan tanda `()`,  membutuhkan tanda `{}` untuk isi perulangan
```
package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
```