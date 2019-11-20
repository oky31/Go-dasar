
# Function
* Fungsi bisa berisi 0 atau banyak parameter dan me-returnkan value dengan tipe data yang terdefinisi
* Mendefinikan tipe data yang di kembaliakan oleh fungsi, di letakan setelah parameter ex: `func add(x int) int {}` 
* Jika parameter difungsi memiliki tipe yang sama abaikan deklarasi tipe untuk yang lain `x int, y int` menjadi `x, y int`
* Fungsi bisa mereturnkan lebih dari 1 variabel
* "naked" return atau "named return parameters", dimana sebuah fungsi tidak mendefinikan return value-nya dialam fungsi, tetapi di returnkan di tipe data fungsi , di gunakan hanya untuk fungsi yang pendek


## Mendefinisikan fungsi

Mendefinisikan fungsi
```
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
```

## Fungsi mengembalikan banyak nilai
Di dalam go fungsi bisa mengebalikan lebih dari 1 nilai
contoh :

```
package main

import "fmt"

func banyak_nilai(x, y int) (int int) {
    x = x * 2
    y = y * 2
    return x,y
}

func main(){
    foo, bar int

    foo = 10
    bar = 20
    fmt.Println(banyak_nilai(foo, bar))
}
```

## "naked" pengembalian nilai
```
pckage main

import "fmt"

function banyak_nilai(x, y int) (x, y int) {
    x = x * 2
    y = y * 2
    return
}

func main() {
    foo, bar int
    
    foo = 10
    bar = 20
    fmt.Println(banyak_nilai(foo,bar))
}
```