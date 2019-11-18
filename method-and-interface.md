# method
* Go tidak memiliki class. Namun, anda bisa mendefinisikan method pada tipe.
* Sebuah method adalah sebuah fungsi dengan argumen khusus receiver.
* Method bisa di deklarasikan di `type`  lain selain `struct`

ada 2 jenis method di go :
* value receiver -> hanya bisa menghitung nilai
* pointer receiver -> bisa menguban nilai di dalam struct

## Pointer receiver
untuk mengubah nilai yang ada di tipe bentukan, gunakan pointer receiver `*T` di mana `T`
adalah tipe bentukan


2 alasan untuk mengunakan Pointer receiver :  
* method dapat mengubah nilai yang ditunjuk oleh receiver
* menghindari menyalin nilai setiap kali method dipanggil

# Interface
.......


