package View

import (
	"ProjectBukuKas/Controller"
	"ProjectBukuKas/Node"
	"fmt"
)

func TampilkanMenu() {
	fmt.Println()
	fmt.Println("[ BUKU KAS HARIAN ]")
	fmt.Println("1. Input Transaksi")
	fmt.Println("2. Tampilkan Daftar Transaksi")
	fmt.Println("3. Hapus Transaksi")
	fmt.Println("4. Update Transaksi")
	fmt.Println("5. Hitung Saldo")
	fmt.Println("6. Keluar")
	fmt.Println()
}

func InputTransaksi() {
	var noTransaksi, nominal int
	var ket, ketPembayaran string

	fmt.Print("Masukkan No Transaksi: ")
	fmt.Scan(&noTransaksi)
	fmt.Print("Masukkan Keterangan: ")
	fmt.Scan(&ket)
	fmt.Print("Masukkan Nominal: ")
	fmt.Scan(&nominal)
	fmt.Print("Masukkan Metode Transaksi: ")
	fmt.Scan(&ketPembayaran)

	Controller.InputTransaksi(noTransaksi, ket, nominal, ketPembayaran)
	fmt.Println()
	fmt.Println("... Transaksi berhasil dibuat ...")
}

func TampilkanDaftarTransaksi() {
	transaksi := Controller.TampilDaftarTransaksi()
	fmt.Println("<< DAFTAR TRANSAKSI >>")
	fmt.Print("------------------")
	fmt.Println()
	if transaksi == nil {
		fmt.Println("Tidak ada transaksi.")
		return
	}
	for transaksi != nil {
		fmt.Printf("No : %d, Ket : %s, Nominal : %d, Metode Transaksi : %s\n",
			transaksi.Data.NoTransaksi, transaksi.Data.Ket, transaksi.Data.Nominal, transaksi.Data.KetPembayaran)
		transaksi = transaksi.Next
	}
	fmt.Println()
	fmt.Print("------------------")
}
func HapusTransaksi() {
	var noTransaksi int
	fmt.Print("Masukkan No Transaksi yang akan dihapus: ")
	fmt.Scan(&noTransaksi)
	fmt.Println()

	if Controller.HapusTransaksi(noTransaksi) {
		fmt.Println()
		fmt.Println("... Transaksi berhasil dihapus ...")
	} else {
		fmt.Println()
		fmt.Println("!! Transaksi tidak ditemukan !!")
	}
}

func UpdateTransaksi() {
	var noTransaksi, nominal int
	var ket, ketPembayaran string

	fmt.Print("Masukkan No Transaksi yang akan diupdate: ")
	fmt.Scan(&noTransaksi)

	transaksi := Node.Transaksi{
		NoTransaksi:   noTransaksi,
		Ket:           ket,
		Nominal:       nominal,
		KetPembayaran: ketPembayaran,
	}

	if Controller.UpdateTransaksi(transaksi) {
		fmt.Print("Masukkan Keterangan Baru: ")
		fmt.Scan(&ket)
		fmt.Print("Masukkan Nominal Baru: ")
		fmt.Scan(&nominal)
		fmt.Print("Masukkan Metode Transaksi Baru: ")
		fmt.Scan(&ketPembayaran)
		fmt.Println()
		fmt.Println("... Transaksi berhasil diupdate ...")
	} else {
		fmt.Println()
		fmt.Println("!! Transaksi tidak ditemukan !!")
	}
}

func HitungSaldo() {
	saldo := Controller.HitungSaldo()
	fmt.Printf("SALDO AKHIR : %d\n", saldo)
	fmt.Println("------------------")
}
