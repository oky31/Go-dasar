# Pointer
Go memiliki pointer,  A pointer holds the memory address of a value.
```
* untuk menunjuk nilai yang tersimpan di dalam variabel
& untuk menunjuk alamat memori yang tersimpan di variabel
```

# struct
* `struct` adalah kumpulan field yang memiliki berbagai macam tipe data
* karna di dalam go tidak ada class, maka `struct` bisa di pakai sebagia attribut


cara mendefinisikan 
```
type namaStruct struct {
    variable1 tipeData1
    variable2 tipeData2
}
```

cara mengunakan 
```
variable := namaStruct{nilaiTipe1, nilaiTipe2, .., nilaiTipeN}
```
* Gunakan `.` untuk mengakses setiap fild di dalam struct

## Pointer Ke structs
setiap filed di dalam structs bisa di akses dengan pointer

untuk meng-akses field di dalam `struct` 
```
(*variablePenyimpanAlamat).Field
        atau
variablePenyimpanAlamat.Field

```

# Array
array `[n]T` dimana `n` panjang array, dan `T` tipe data

* array memiliki ukuran yang tetap

# Slice (array dinamis)
* memiliki ukuran yang dinamis
* `[]T` , sebuah slice bertipe data T
* `a[bawah:atas]`, untuk memotong slice
* slice tidak menyimpan data, hanya mengacu bagian dari array
* mengubah elemen pada slice, berarti mengubah elemen pada array
* untuk menabahkan elemen ke dalam sline `func append(s []T, vs ...T) []T`
* `a[low:hight]` ambil elemen mulai dari index low sampai index hight tetapi tidak termasuk height

# Range
* `range` untuk perulangan di dalam slice,  `range` mengembalikan 2 nilai yaitu index dan value
```
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
}
```

* `_` untuk melewati index atau value di dalam perulangan

# Map
* Sebuah map memetakan sebuah kunci (key) dengan nilainya
* untuk membuat maps `varName := make(map[key-type]val-type)`
* mengisi atau mengubah elemen map `m[key] = elemen`
* mengambil nilai `elem = m[key]`
* menghapus `delete(m,key)`
* menguji sebuah key ada di dalam map `elem, ok := m[key]`

# Fungsi menjadi parameter fungsi lain