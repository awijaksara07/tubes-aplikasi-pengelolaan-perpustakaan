package main

import "fmt"

const NMAX int = 100

// Type Buku berisi data untuk menampung buku berdasar ID, Kategori, Judul, Penulis, Tahun
type Buku struct {
	ID       int
	Kategori string
	Judul    string
	Penulis  string
	Tahun    int
	Stok     int
}

// Type anggota berisi data untuk menampung IDAnggota, Nama, Kelas
type Anggota struct {
	IDAnggota int
	Nama      string
	Kelas     string
}

// Type Peminjaman berisi data untuk menampung IDPinjam, IDAnggota, IDBuku, LamaHari
type Peminjaman struct {
	IDPinjam  int
	IDAnggota int
	IDBuku    int
	LamaHari  int
}

type tabBuku [NMAX]Buku
type tabAnggota [NMAX]Anggota
type tabPeminjaman [NMAX]Peminjaman

func greetings() {
	fmt.Println("\n===========================================================")
	fmt.Println("|      SELAMAT DATANG DI SISTEM PERPUSTAKAAN DIGITAL      |")
	fmt.Println("===========================================================")
}


func isiDummyBuku(A *tabBuku, n *int) {
	A[0] = Buku{1, "Informatika", "Algoritma", "Rinaldi", 2022, 5}
	A[1] = Buku{2, "Informatika", "BasisData", "Kadir", 2021, 3}
	A[2] = Buku{3, "Matematika", "Kalkulus", "Stewart", 2020, 8}
	A[3] = Buku{4, "Sains", "FisikaDasar", "Halliday", 2019, 2}
	A[4] = Buku{5, "Sains", "KimiaDasar", "Chang", 2018, 4}

	*n = 5
}

func main() {
	var dataBuku tabBuku
	var dataAnggota tabAnggota
	var dataPinjam tabPeminjaman

	var nBuku, nAnggota, nPinjam int
	var pilihan int

	pilihan = -1
	greetings()

	isiDummyBuku(&dataBuku, &nBuku)
	for pilihan != 0 {
		menuUtama()
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			menuManajemenBuku(&dataBuku, &nBuku)
		case 2:
			menuManajemenAnggota(&dataAnggota, &nAnggota)
		case 3:
			menuPeminjamanBuku(&dataPinjam, &nPinjam, &dataBuku, &nBuku, &dataAnggota, nAnggota)
		case 4:
			tampilkanStatistik(dataBuku, nBuku, nPinjam, nAnggota)
		case 0:
			fmt.Println("\nProgram selesai. Terima kasih!")
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func menuUtama() {
	/*
	IS.
	*/
	fmt.Println("\n===================================================")
	fmt.Println("|                   MENU UTAMA                    |")
	fmt.Println("===================================================")
	fmt.Println("| 1. Manajemen Data Buku                          |")
	fmt.Println("| 2. Manajemen Data Anggota                       |")
	fmt.Println("| 3. Transaksi Peminjaman Buku                    |")
	fmt.Println("| 4. Statistik & Ringkasan Data                   |")
	fmt.Println("| 0. Keluar Aplikasi                              |")
	fmt.Println("===================================================")
}

func menuBuku() {
	fmt.Println("\n===================================================")
	fmt.Println("|                  MANAJEMEN BUKU                 |")
	fmt.Println("===================================================")
	fmt.Println("| 1. Tambah Data Buku Baru                        |")
	fmt.Println("| 2. Tampilkan Semua Daftar Buku                  |")
	fmt.Println("| 3. Ubah Informasi Data Buku                     |")
	fmt.Println("| 4. Hapus Data Buku Dari Sistem                  |")
	fmt.Println("| 5. Pencarian Data Buku (Searching)              |")
	fmt.Println("| 6. Pengurutan Data Buku (Sorting)               |")
	fmt.Println("| 7. Tampilkan stok Terbanyak/Terkecil            |")
	fmt.Println("| 0. Kembali ke Menu Utama                        |")
	fmt.Println("===================================================")
}

func menuManajemenBuku(A *tabBuku, n *int) {
	var pilihan int
	pilihan = -1

	for pilihan != 0 {
		menuBuku()
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahBuku(A, n)
		case 2:
			cetakDataBuku(*A, *n)
		case 3:
			ubahBuku(A, *n)
		case 4:
			hapusBuku(A, n)
		case 5:
			menuCariBuku(A, *n)
		case 6:
			menuUrutBuku(A, *n)
		case 7:
			pilihanStok(*A, *n)
		case 0:
			fmt.Println("Kembali ke menu utama...")
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func pilihanStok(A tabBuku, n int) {
	var pilihan int
	var terbesar bool

	if isEmpty(n){
		fmt.Println("Data buku kosong. Silakan isi terlebih dahulu.")
	}else{
		fmt.Println("\n======================================")
		fmt.Println("| 1. Tampilkan Stok Buku Terbesar    |")
		fmt.Println("| 2. Tampilkan Stok Buku Terkecil    |")
		fmt.Println("======================================")

		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)

		terbesar = pilihan == 1

		cariStokTerbesarTerkecil(A, n, terbesar)
	}
	
}

func cariStokTerbesarTerkecil(A tabBuku, n int, terbesar bool){
	var idx, i  int
	var status string

	idx = 0

	if terbesar{
		for i = 1; i<n; i++ {
			if A[i].Stok > A[idx].Stok{
				idx = i
			}
		}

		if A[idx].Stok == 0 {
			status = "Habis"
		} else if A[idx].Stok <= 2 {
			status = "Menipis"
		} else {
			status = "Tersedia"
		}

		fmt.Println("\nStok Buku Terbesar:")
		fmt.Println("------------------------------------")
		fmt.Println("ID Buku    :", A[idx].ID)
		fmt.Println("Kategori   :", A[idx].Kategori)
		fmt.Println("Judul Buku :", A[idx].Judul)
		fmt.Println("Penulis    :", A[idx].Penulis)
		fmt.Println("Tahun      :", A[idx].Tahun)
		fmt.Println("Stok       :", A[idx].Stok)
		fmt.Println("Status     :", status)
		fmt.Println("------------------------------------")
	}else{
		for i = 1; i<n; i++ {
			if A[i].Stok < A[idx].Stok{
				idx = i
			}
		}

		if A[idx].Stok == 0 {
			status = "Habis"
		} else if A[idx].Stok <= 2 {
			status = "Menipis"
		} else {
			status = "Tersedia"
		}

		fmt.Println("\nStok Buku Terkecil:")
		fmt.Println("------------------------------------")
		fmt.Println("ID Buku    :", A[idx].ID)
		fmt.Println("Kategori   :", A[idx].Kategori)
		fmt.Println("Judul Buku :", A[idx].Judul)
		fmt.Println("Penulis    :", A[idx].Penulis)
		fmt.Println("Tahun      :", A[idx].Tahun)
		fmt.Println("Stok       :", A[idx].Stok)
		fmt.Println("Status     :", status)
		fmt.Println("------------------------------------")
	}

}

func binarySearchBukuByID(A tabBuku, n int, id int) int {
	var left, right, mid, idx int

	idx = -1
	left = 0
	right = n - 1
	for left <= right && idx == -1 {
		mid = (left + right) / 2
		if A[mid].ID == id {
			idx = mid
		} else if A[mid].ID < id {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return idx
}

func binarySearchAnggotaByID(A tabAnggota, n int, id int) int {
	var left, right, mid, idx int

	idx = -1
	left = 0
	right = n - 1
	for left <= right && idx == -1 {
		mid = (left + right) / 2
		if A[mid].IDAnggota == id {
			idx = mid
		} else if A[mid].IDAnggota < id {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return idx
}

func binarySearchPinjamByID(P tabPeminjaman, n int, id int) int {
	var left, right, mid, idx int

	idx = -1
	left = 0
	right = n - 1
	for left <= right && idx == -1 {
		mid = (left + right) / 2
		if P[mid].IDPinjam == id {
			idx = mid
		} else if P[mid].IDPinjam < id {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return idx
}

func idBukuTerakhir(A tabBuku, n int) int {
	if n == 0 {
		return 0
	}
	return A[n-1].ID
}

func idAnggotaTerakhir(A tabAnggota, n int) int {
	if n == 0 {
		return 0
	}
	return A[n-1].IDAnggota
}

func idPinjamTerakhir(P tabPeminjaman, n int) int {
	if n == 0 {
		return 0
	}
	return P[n-1].IDPinjam
}

func tambahBuku(A *tabBuku, n *int) {
	var banyak, i int

	if isFull(*n) {
		fmt.Println("Data buku sudah penuh, tidak bisa menambah data lagi.")
	} else {
		fmt.Printf("Maksimal bisa tambah %d buku\n", NMAX-*n)
		fmt.Print("Berapa buku yang ingin ditambah: ")
		fmt.Scan(&banyak)

		if banyak > NMAX-*n {
			banyak = NMAX - *n
			fmt.Println("Jumlah disesuaikan ke batas maksimal slot yang tersedia.")
		}

		for i = 0; i < banyak; i++ {
			fmt.Printf("\n===== INPUT BUKU KE-%d =====\n", i+1)

			A[*n].ID = idBukuTerakhir(*A, *n) + 1
			fmt.Println("ID Buku (Otomatis) :", A[*n].ID)

			fmt.Print("Judul Buku         : ")
			fmt.Scan(&A[*n].Judul)

			fmt.Print("Kategori Buku      : ")
			fmt.Scan(&A[*n].Kategori)

			fmt.Print("Penulis Buku       : ")
			fmt.Scan(&A[*n].Penulis)

			fmt.Print("Tahun Terbit       : ")
			fmt.Scan(&A[*n].Tahun)

			fmt.Print("Jumlah Stok Buku   : ")
			fmt.Scan(&A[*n].Stok)

			*n++
		}
		fmt.Println("\nData berhasil ditambahkan ke dalam sistem.")
	}
}

func cetakDataBuku(A tabBuku, n int) {
	var i int
	var status string

	if isEmpty(n) {
		fmt.Println("Data buku kosong. Silakan isi terlebih dahulu.")
	} else {
		fmt.Println("\n============================================= DAFTAR BUKU PERPUSTAKAAN ==============================================")
		fmt.Printf("| %-3s | %-6s | %-15s | %-25s | %-20s | %-6s | %-5s | %-12s |\n", "No", "ID", "Kategori", "Judul Buku", "Penulis", "Tahun", "Stok", "Status")
		fmt.Println("=====================================================================================================================")

		for i = 0; i < n; i++ {
			if A[i].Stok == 0 {
				status = "Habis"
			} else if A[i].Stok <= 2 {
				status = "Menipis"
			} else {
				status = "Tersedia"
			}
			fmt.Printf("| %-3d | %-6d | %-15s | %-25s | %-20s | %-6d | %-5d | %-12s |\n", i+1, A[i].ID, A[i].Kategori, A[i].Judul, A[i].Penulis, A[i].Tahun, A[i].Stok, status)
		}
		fmt.Println("=====================================================================================================================")
	}
}

func ubahBuku(A *tabBuku, n int) {
	var id, idx int

	if isEmpty(n) {
		fmt.Println("Data kosong, modifikasi dibatalkan.")
	} else {
		fmt.Print("Masukkan ID buku yang ingin diubah: ")
		fmt.Scan(&id)

		idx = binarySearchBukuByID(*A, n, id)

		if idx == -1 {
			fmt.Println("ID tidak ditemukan.")
		} else {
			fmt.Println("\n===== MASUKKAN DATA BARU (PERUBAHAN) =====")
			fmt.Print("Judul Baru   : ")
			fmt.Scan(&A[idx].Judul)

			fmt.Print("Kategori Baru: ")
			fmt.Scan(&A[idx].Kategori)

			fmt.Print("Penulis Baru : ")
			fmt.Scan(&A[idx].Penulis)

			fmt.Print("Tahun Baru   : ")
			fmt.Scan(&A[idx].Tahun)

			fmt.Print("Stok Baru    : ")
			fmt.Scan(&A[idx].Stok)

			fmt.Println("\nData buku berhasil diperbarui.")
		}
	}
}

func hapusBuku(A *tabBuku, n *int) {
	var id, idx, i int

	if isEmpty(*n) {
		fmt.Println("Data kosong, penghapusan dibatalkan.")
	} else {
		fmt.Print("Masukkan ID buku yang ingin dihapus: ")
		fmt.Scan(&id)

		idx = binarySearchBukuByID(*A, *n, id)

		if idx == -1 {
			fmt.Println("ID data buku tidak ditemukan.")
		} else {
			for i = idx; i < *n-1; i++ {
				A[i] = A[i+1]
			}
			*n--
			fmt.Println("\nData buku berhasil dihapus dari sistem.")
		}
	}
}

func menuCariBuku(A *tabBuku, n int) {
	var pilihan, kriteriaInt, hasil int
	var kriteriaString string

	if isEmpty(n) {
		fmt.Println("Data kosong, pencarian tidak dapat dilakukan.")
	} else {
		fmt.Println("\n=====================================")
		fmt.Println("|         MENU PENCARIAN BUKU       |")
		fmt.Println("=====================================")
		fmt.Println("| 1. Cari ID Buku (Binary Search)   |")
		fmt.Println("| 2. Cari Judul (Sequential Search) |")
		fmt.Println("| 0. Kembali                        |")
		fmt.Println("=====================================")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)

		hasil = -1
		if pilihan == 1 {
			fmt.Print("Masukkan ID Buku: ")
			fmt.Scan(&kriteriaInt)
			hasil = binarySearchBukuByID(*A, n, kriteriaInt)
		} else if pilihan == 2 {
			fmt.Print("Masukkan Judul Buku: ")
			fmt.Scan(&kriteriaString)
			hasil = sequentialSearchJudul(*A, n, kriteriaString)
		}

		if pilihan == 1 || pilihan == 2 {
			cetakHasilCariBuku(*A, hasil)
		}
	}
}

func cetakHasilCariBuku(A tabBuku, idx int) {
	var status string
	if idx == -1 {
		fmt.Println("Data yang dicari tidak ditemukan.")
	} else {
		if A[idx].Stok == 0 {
			status = "Habis"
		} else if A[idx].Stok <= 2 {
			status = "Menipis"
		} else {
			status = "Tersedia"
		}
		fmt.Println("\nDATA DITEMUKAN:")
		fmt.Println("------------------------------------")
		fmt.Println("ID Buku    :", A[idx].ID)
		fmt.Println("Kategori   :", A[idx].Kategori)
		fmt.Println("Judul Buku :", A[idx].Judul)
		fmt.Println("Penulis    :", A[idx].Penulis)
		fmt.Println("Tahun      :", A[idx].Tahun)
		fmt.Println("Stok       :", A[idx].Stok)
		fmt.Println("Status     :", status)
		fmt.Println("------------------------------------")
	}
}

func menuUrutBuku(A *tabBuku, n int) {
	var pilKriteria, pilAlgoritma, pilArah int
	var asc bool
	var temp tabBuku

	if isEmpty(n) {
		fmt.Println("Data kosong, sorting dibatalkan.")
	} else {
		fmt.Println("\n======================================")
		fmt.Println("|         PILIH KRITERIA URUT        |")
		fmt.Println("======================================")
		fmt.Println("| 1. ID Buku                         |")
		fmt.Println("| 2. Tahun Terbit                    |")
		fmt.Println("| 3. Stok Buku                       |")
		fmt.Println("======================================")
		fmt.Print("Pilihan Kriteria: ")
		fmt.Scan(&pilKriteria)

		if pilKriteria == 1 || pilKriteria == 2 || pilKriteria == 3 {
			fmt.Println("\n======================================")
			fmt.Println("|           PILIH ALGORITMA          |")
			fmt.Println("======================================")
			fmt.Println("| 1. Gunakan Selection Sort          |")
			fmt.Println("| 2. Gunakan Insertion Sort          |")
			fmt.Println("======================================")
			fmt.Print("Pilihan Algoritma: ")
			fmt.Scan(&pilAlgoritma)

			if pilAlgoritma == 1 || pilAlgoritma == 2 {
				fmt.Println("\n====================================")
				fmt.Println("|              ARAH URUTAN         |")
				fmt.Println("====================================")
				fmt.Println("| 1. Kecil ke Besar (Ascending)    |")
				fmt.Println("| 2. Besar ke Kecil (Descending)   |")
				fmt.Println("====================================")
				fmt.Print("Pilihan: ")
				fmt.Scan(&pilArah)

				asc = pilArah == 1
				temp = *A

				if pilAlgoritma == 1 {
					selectionSortBuku(&temp, n, pilKriteria, asc)
					fmt.Println("\nData berhasil diurutkan dengan Selection Sort.")
				} else {
					insertionSortBuku(&temp, n, pilKriteria, asc)
					fmt.Println("\nData berhasil diurutkan dengan Insertion Sort.")
				}
				cetakDataBuku(temp, n)
			} else {
				fmt.Println("Pilihan algoritma tidak valid.")
			}
		} else {
			fmt.Println("Pilihan kriteria tidak valid.")
		}
	}
}

func menuAnggota() {
	fmt.Println("\n===================================================")
	fmt.Println("|                MANAJEMEN ANGGOTA                |")
	fmt.Println("===================================================")
	fmt.Println("| 1. Tambah Anggota Perpustakaan                  |")
	fmt.Println("| 2. Tampilkan Semua Anggota                      |")
	fmt.Println("| 3. Ubah Informasi Data Anggota                  |")
	fmt.Println("| 4. Hapus Data Anggota Dari Sistem               |")
	fmt.Println("| 0. Kembali ke Menu Utama                        |")
	fmt.Println("===================================================")
}

func menuManajemenAnggota(A *tabAnggota, n *int) {
	var pilihan int
	pilihan = -1

	for pilihan != 0 {
		menuAnggota()
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahAnggota(A, n)
		case 2:
			cetakDataAnggota(*A, *n)
		case 3:
			ubahAnggota(A, *n)
		case 4:
			hapusAnggota(A, n)
		case 0:
			fmt.Println("Kembali ke menu utama...")
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func tambahAnggota(A *tabAnggota, n *int) {
	if isFull(*n) {
		fmt.Println("Kapasitas database anggota penuh.")
	} else {
		fmt.Println("\n===== TAMBAH DATA ANGGOTA BARU =====")
		A[*n].IDAnggota = idAnggotaTerakhir(*A, *n) + 1
		fmt.Println("ID Anggota (Otomatis) :", A[*n].IDAnggota)

		fmt.Print("Nama Lengkap          : ")
		fmt.Scan(&A[*n].Nama)

		fmt.Print("Kelas / Jurusan       : ")
		fmt.Scan(&A[*n].Kelas)

		*n++
		fmt.Println("\nAnggota baru sukses didaftarkan.")
	}
}

func cetakDataAnggota(A tabAnggota, n int) {
	var i int

	if isEmpty(n) {
		fmt.Println("Belum ada anggota terdaftar.")
	} else {
		fmt.Println("\n================== DAFTAR ANGGOTA PERPUSTAKAAN ===================")
		fmt.Printf("| %-3s | %-10s | %-25s | %-15s |\n", "No", "ID Anggota", "Nama Lengkap", "Kelas")
		fmt.Println("==================================================================")

		for i = 0; i < n; i++ {
			fmt.Printf("| %-3d | %-10d | %-25s | %-15s |\n", i+1, A[i].IDAnggota, A[i].Nama, A[i].Kelas)
		}
		fmt.Println("==================================================================")
	}
}

func ubahAnggota(A *tabAnggota, n int) {
	var id, idx int

	if isEmpty(n) {
		fmt.Println("Data kosong.")
	} else {
		fmt.Print("Masukkan ID Anggota yang ingin diubah: ")
		fmt.Scan(&id)

		idx = binarySearchAnggotaByID(*A, n, id)

		if idx == -1 {
			fmt.Println("Data anggota tidak ditemukan.")
		} else {
			fmt.Println("\n===== MASUKKAN PERUBAHAN DATA =====")
			fmt.Print("Nama Baru  : ")
			fmt.Scan(&A[idx].Nama)

			fmt.Print("Kelas Baru : ")
			fmt.Scan(&A[idx].Kelas)

			fmt.Println("\nInformasi data anggota berhasil diubah.")
		}
	}
}

func hapusAnggota(A *tabAnggota, n *int) {
	var id, idx, i int

	if isEmpty(*n) {
		fmt.Println("Tidak ada data untuk dihapus.")
	} else {
		fmt.Print("Masukkan ID Anggota yang akan dihapus: ")
		fmt.Scan(&id)

		idx = binarySearchAnggotaByID(*A, *n, id)

		if idx == -1 {
			fmt.Println("Anggota dengan ID tersebut tidak ditemukan.")
		} else {
			for i = idx; i < *n-1; i++ {
				A[i] = A[i+1]
			}
			*n--
			fmt.Println("\nAnggota sukses dikeluarkan dari sistem data.")
		}
	}
}

func menuPeminjaman() {
	fmt.Println("\n===================================================")
	fmt.Println("|                  PEMINJAMAN BUKU                |")
	fmt.Println("===================================================")
	fmt.Println("| 1. Pinjam Buku (Transaksi Baru)                 |")
	fmt.Println("| 2. Tampilkan Seluruh Log Peminjaman             |")
	fmt.Println("| 3. Kembalikan Buku (Selesaikan Transaksi)       |")
	fmt.Println("| 4. Perpanjang Durasi Pinjam                     |")
	fmt.Println("| 0. Kembali ke Menu Utama                        |")
	fmt.Println("===================================================")
}

func menuPeminjamanBuku(P *tabPeminjaman, nPinjam *int, B *tabBuku, nBuku *int, A *tabAnggota, nAnggota int) {
	var pilihan int
	pilihan = -1

	for pilihan != 0 {
		menuPeminjaman()
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			pinjamBuku(P, nPinjam, B, *nBuku, A, nAnggota)
		case 2:
			cetakPeminjaman(*P, *nPinjam)
		case 3:
			kembalikanBuku(P, nPinjam, B, *nBuku)
		case 4:
			perpanjangPinjam(P, *nPinjam)
		case 0:
			fmt.Println("Kembali ke menu utama...")
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func pinjamBuku(P *tabPeminjaman, nPinjam *int, B *tabBuku, nBuku int, A *tabAnggota, nAnggota int) {
	var idA, idB, idxBuku, idxAnggota int

	if isFull(*nPinjam) {
		fmt.Println("Antrean peminjaman penuh.")
	} else {
		fmt.Println("\n===== TRANSAKSI PEMINJAMAN BARU =====")
		fmt.Print("Masukkan ID Anggota : ")
		fmt.Scan(&idA)

		idxAnggota = binarySearchAnggotaByID(*A, nAnggota, idA)

		if idxAnggota == -1 {
			fmt.Println("Maaf ID Anggota tidak terdaftar di sistem!")
		} else {
			fmt.Print("Masukkan ID Buku    : ")
			fmt.Scan(&idB)

			idxBuku = binarySearchBukuByID(*B, nBuku, idB)

			if idxBuku == -1 {
				fmt.Println("Maaf ID Buku tidak ditemukan di katalog!")
			} else {
				if B[idxBuku].Stok <= 0 {
					fmt.Println("Maaf Stok buku ini sedang kosong / habis dipinjam!")
				} else {
					P[*nPinjam].IDPinjam = idPinjamTerakhir(*P, *nPinjam) + 1
					P[*nPinjam].IDAnggota = idA
					P[*nPinjam].IDBuku = idB

					fmt.Print("Lama Hari Pinjam    : ")
					fmt.Scan(&P[*nPinjam].LamaHari)

					B[idxBuku].Stok--
					*nPinjam++

					fmt.Println("\nTransaksi Berhasil! Stok buku otomatis dikurangi 1.")
				}
			}
		}
	}
}

func cetakPeminjaman(A tabPeminjaman, n int) {
	var i int

	if isEmpty(n) {
		fmt.Println("Tidak ada transaksi peminjaman aktif saat ini.")
	} else {
		fmt.Println("\n======================== LOG AKTIVITAS PINJAM ========================")
		fmt.Printf("| %-3s | %-12s | %-12s | %-12s | %-15s |\n", "No", "ID Transaksi", "ID Anggota", "ID Buku", "Durasi (Hari)")
		fmt.Println("======================================================================")

		for i = 0; i < n; i++ {
			fmt.Printf("| %-3d | %-12d | %-12d | %-12d | %-15d |\n", i+1, A[i].IDPinjam, A[i].IDAnggota, A[i].IDBuku, A[i].LamaHari)
		}
		fmt.Println("======================================================================")
	}
}

func kembalikanBuku(P *tabPeminjaman, nPinjam *int, B *tabBuku, nBuku int) {
	var idTrans, idxTrans, idxBuku, i int

	if isEmpty(*nPinjam) {
		fmt.Println("Tidak ada data peminjaman.")
	} else {
		fmt.Print("Masukkan ID Transaksi Peminjaman: ")
		fmt.Scan(&idTrans)

		idxTrans = binarySearchPinjamByID(*P, *nPinjam, idTrans)

		if idxTrans == -1 {
			fmt.Println("ID Transaksi tidak valid / tidak ditemukan.")
		} else {
			idxBuku = binarySearchBukuByID(*B, nBuku, P[idxTrans].IDBuku)

			if idxBuku != -1 {
				B[idxBuku].Stok++
			}

			for i = idxTrans; i < *nPinjam-1; i++ {
				P[i] = P[i+1]
			}
			*nPinjam--

			fmt.Println("\nPengembalian Berhasil Selesai! Stok buku dikembalikan.")
		}
	}
}

func perpanjangPinjam(P *tabPeminjaman, n int) {
	var idTrans, tambahan, idx int

	if isEmpty(n) {
		fmt.Println("Data transaksi kosong.")
	} else {
		fmt.Print("Masukkan ID Transaksi: ")
		fmt.Scan(&idTrans)

		idx = binarySearchPinjamByID(*P, n, idTrans)

		if idx == -1 {
			fmt.Println("Transaksi tidak ditemukan.")
		} else {
			fmt.Print("Masukkan jumlah hari perpanjangan: ")
			fmt.Scan(&tambahan)
			P[idx].LamaHari += tambahan
			fmt.Println("\nDurasi peminjaman berhasil diperpanjang.")
		}
	}
}

func lebihKecil(b1, b2 Buku, kriteria int, asc bool) bool {
	var kondisi bool

	if kriteria == 1 {
		if asc {
			kondisi = b1.ID < b2.ID
		} else {
			kondisi = b1.ID > b2.ID
		}
	} else if kriteria == 2 {
		if asc {
			kondisi = b1.Tahun < b2.Tahun
		} else {
			kondisi = b1.Tahun > b2.Tahun
		}
	} else if kriteria == 3 {
		if asc {
			kondisi = b1.Stok < b2.Stok
		} else {
			kondisi = b1.Stok > b2.Stok
		}
	}

	return kondisi
}

func selectionSortBuku(A *tabBuku, n int, kriteria int, asc bool) {
	var i, j, idx int
	var temp Buku

	for i = 0; i < n-1; i++ {
		idx = i
		for j = i + 1; j < n; j++ {
			if lebihKecil(A[j], A[idx], kriteria, asc) {
				idx = j
			}
		}
		temp = A[i]
		A[i] = A[idx]
		A[idx] = temp
	}
}

func insertionSortBuku(A *tabBuku, n int, kriteria int, asc bool) {
	var i, pass int
	var temp Buku

	for pass = 1; pass < n; pass++ {
		temp = A[pass]
		i = pass
		for i > 0 && lebihKecil(temp, A[i-1], kriteria, asc) {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
	}
}

func sequentialSearchJudul(A tabBuku, n int, judul string) int {
	var idx, i int
	idx = -1
	i = 0
	for i < n && idx == -1 {
		if A[i].Judul == judul {
			idx = i
		}
		i++
	}
	return idx
}

func tampilkanStatistik(A tabBuku, nBuku int, nPinjam int, nAnggota int) {
	var i, j, k, count int
	var kategoriUnik [NMAX]string
	var nKategori int
	var ada bool

	fmt.Println("\n========================================================")
	fmt.Println("|              RINGKASAN STATISTIK SISTEM              |")
	fmt.Println("========================================================")
	fmt.Printf("| %-38s : %-11d |\n", "Total Koleksi Buku (Stok Fisik)", totalStok(A, nBuku))
	fmt.Printf("| %-38s : %-11d |\n", "Jumlah Variasi Judul Buku", nBuku)
	fmt.Printf("| %-38s : %-11d |\n", "Jumlah Transaksi Peminjaman Aktif", nPinjam)
	fmt.Printf("| %-38s : %-11d |\n", "Jumlah Anggota Terdaftar", nAnggota)
	fmt.Println("========================================================")

	nKategori = 0
	for i = 0; i < nBuku; i++ {
		ada = false
		for j = 0; j < nKategori; j++ {
			if A[i].Kategori == kategoriUnik[j] {
				ada = true
			}
		}
		if !ada {
			kategoriUnik[nKategori] = A[i].Kategori
			nKategori++
		}
	}

	fmt.Println("\n========================================")
	fmt.Println("|       SEBARAN STOK PER KATEGORI      |")
	fmt.Println("========================================")
	fmt.Printf("| %-22s | %-11s |\n", "Nama Kategori", "Total Stok")
	fmt.Println("----------------------------------------")

	if nBuku == 0 {
		fmt.Printf("| %-36s |\n", "Belum ada data buku di dalam sistem")
	} else {
		for i = 0; i < nKategori; i++ {
			count = 0
			for k = 0; k < nBuku; k++ {
				if A[k].Kategori == kategoriUnik[i] {
					count += A[k].Stok
				}
			}
			fmt.Printf("| %-22s | %-11d |\n", kategoriUnik[i], count)
		}
	}
	fmt.Println("========================================")
}

func totalStok(A tabBuku, n int) int {
	if n == 0 {
		return 0
	}
	return A[n-1].Stok + totalStok(A, n-1)
}

func isEmpty(n int) bool {
	return n == 0
}

func isFull(n int) bool {
	return n >= NMAX
}
