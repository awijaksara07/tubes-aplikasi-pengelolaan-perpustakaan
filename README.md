# Sistem Perpustakaan Digital

Aplikasi Sistem Perpustakaan Digital merupakan program berbasis Command Line Interface (CLI) yang dibuat menggunakan bahasa pemrograman Golang sebagai tugas besar mata kuliah Algoritma dan Pemrograman 2.

Program ini digunakan untuk mengelola data buku, data anggota, serta transaksi peminjaman buku. Seluruh data disimpan menggunakan array of struct tanpa menggunakan database sehingga berbagai proses pengolahan data dilakukan langsung melalui algoritma yang telah dipelajari pada perkuliahan.

---

# Anggota Kelompok

* I Nyoman Ajita Wijaksara
* Muhammad Ikhsan

---

# Fitur Aplikasi

### Manajemen Data Buku

* Menambah data buku
* Menampilkan daftar buku
* Mengubah data buku
* Menghapus data buku
* Mencari buku berdasarkan ID maupun judul
* Mengurutkan data buku
* Menampilkan buku dengan stok terbesar dan terkecil

### Manajemen Data Anggota

* Menambah anggota
* Menampilkan anggota
* Mengubah data anggota
* Menghapus anggota

### Transaksi Peminjaman

* Peminjaman buku
* Pengembalian buku
* Perpanjangan lama peminjaman
* Menampilkan log peminjaman

### Statistik

* Total stok seluruh buku
* Jumlah judul buku
* Jumlah anggota
* Jumlah transaksi aktif
* Sebaran stok berdasarkan kategori

---

## Konsep Algoritma yang Digunakan

Program ini mengimplementasikan beberapa konsep yang dipelajari pada mata kuliah Algoritma dan Pemrograman, yaitu:

* Array of Struct
* Pencarian Nilai Ekstrem (Stok Terbesar & Terkecil)
* Procedure dan Function
* Binary Search
* Sequential Search
* Selection Sort
* Insertion Sort
* Rekursif

---

## Alur Program

1. Program menampilkan menu utama.
2. Pengguna memilih menu yang diinginkan.
3. Sistem menjalankan proses sesuai pilihan pengguna.
4. Setelah proses selesai, pengguna dapat kembali ke menu sebelumnya atau ke menu utama.
5. Program berhenti ketika pengguna memilih menu **Keluar**.

---

## Hubungan Antar Data

Program memiliki tiga entitas utama:

* Buku
* Anggota
* Peminjaman

Hubungan antara Buku dan Anggota bersifat many-to-many, sehingga dihubungkan oleh entitas Peminjaman yang menyimpan informasi transaksi peminjaman.

---

## Validasi Program

Beberapa validasi yang diterapkan pada aplikasi antara lain:

* Kapasitas array maksimum 100 data.
* ID buku dan anggota dibuat otomatis.
* Tahun terbit dan stok tidak boleh bernilai negatif.
* Lama peminjaman harus lebih dari 0 hari.
* Peminjaman hanya dapat dilakukan apabila ID anggota dan ID buku terdaftar.
* Buku tidak dapat dipinjam apabila stok habis.
* Pengembalian buku akan menambah stok secara otomatis.
* Perpanjangan hanya dapat dilakukan pada transaksi yang masih tersedia.

---

## Mata Kuliah

Algoritma dan Pemrograman 2

Program Studi S1 Informatika

Fakultas Informatika

Telkom University
