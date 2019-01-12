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
### Menyingkat `if` statement
Untuk menyingkat if statment, variabel bisa di deklarasikan di dalam if statement
```
if variabel := value; expression {
	return variabel
}
```  

## Switch
* `switch` statmen adalah cara singkat untuk menulis sebuah sequence dari `if- else` statements.
`switch` akan menampilkan hanya satu expression yang bernilai true
* switch memerikas case dari atas ke bawah, dan akan berhenti ketika case bernilai benar
* switch tanpa kondisi, itu berarti `switch true`
```
switch variabel := value; variabel {
case value:
	statement
case value:
	statement
default:
	statement
}
```

## Defer (penudaan)
* ``defer`` statement yang menunda eksekusi, statement `defer` akan di eksekusi setelah semua statemen di dalam fungsi
selesai di eksekusi
* defer bisa di masukan ke dalam stack