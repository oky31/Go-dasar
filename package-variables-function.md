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

# Deklarasi Variabel
* variabel di deklarasiakan dengan keyword `var`, variabel bisa di deklarasikan pada level `function` atau `package`
* tipe data bisa di abaikan, dengan catatan variabel harus sudah di inisialisasi
* di dalam fungsi bisa mengunakan `:=` sebagai ganti keyword `var`, hal ini tidak berlaku di luar fungsi
* variable yang di deklarasikan tanpa di beri nilai akan otomanis berisi 0

Contoh :  
```

```
## Deklarasi tanpa di inisialisasi
deklarasi variabel tanpa di inisialisasi memili default value berdasarkan masing masing tipe
```
0 -> untuk tipe numeric
false -> untuk tipe logika
"" -> untuk tipe string
```

# Tipe Data
Tipe data dasar 
```
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```

* setiap tipe data yang di deklarasikan tanpa default value akan memiliki default value 0, false, atau ""

## Konversi Tipe Data
Untuk mengkonversi tipe data gunakan `T(v)` contoh :
```
var i int = 42
var f float64 = float64(i)
```

## Type inference (kesimpulan data tipe)
Ketika tipe data tidak di deklarasikan di variabel, variabel akan otomatis menyesuaikan tipe data dengan nilai yang di isikan di dalamnya


# Constanta
* Gunakan keword `const` untuk mendefinisikan constanta, dan tidak bisa di deklarasikan dengan sintax `:=`
* Constanta yang di deklarasikan tanpa tipe data akan mengambil tipe data yang dibutuhkan berdasarkan konteksnya


# Kata Kunci
* Setiap Go Program terdiri dari paket-paket
* untuk memangil setiap fungsi di dalam package harus di awali dengan huruf besa, misal `math.Pi`

# Compile progarm
```
go bulid    #compile program
go run      #menjalankan tanpa di compile
```