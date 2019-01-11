# Belajar Go

# Instalasai Di ubuntu
```
sudo apt install golang-go
sudo apt install gccgo-go
```

# Packages
Setiap Go Program terdiri dari paket-paket
Dan Program utama ada di package ``package main``

## Import Package
Cara yang baik gunakan 
```
import "fmt"
import "math"
```

# Function
* Fungsi bisa berisi 0 atau banyak parameter dan me-returnkan value dengan tipe data yang terdefinisi
* Jika parameter difungsi memiliki tipe yang sama abaikan deklarasi tipe untuk yang lain `x int, y int` menjadi `x, y int`
* Fungsi bisa mereturnkan lebih dari 1 variabel
* "naked" return, dimana sebuah fungsi tidak mendefinikan return value-nya, di gunakan hanya untuk fungsi yang pendek

# Variabel
* variabel di deklarasiakan dengan keyword `var`, variabel bisa di deklarasikan pada level `function` atau `package`
* tipe data bisa di abaikan, dengan catatan variabel harus sudah di inisialisasi
* di dalam fungsi bisa mengunakan `:=` sebagai ganti keyword `var`, hal ini tidak berlaku di luar fungsi

# Kata Kunci
* Setiap Go Program terdiri dari paket-paket
* untuk memangil setiap fungsi di dalam package harus di awali dengan huruf besa, misal `math.Pi`

# Compile progarm
```
go bulid
```