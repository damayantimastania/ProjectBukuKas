package main

import (
	"ProjectBukuKas/View"
	"fmt"
)

func main() {
	for {
		View.TampilkanMenu()
		var choice int
		fmt.Print("Pilih Menu: ")
		fmt.Scan(&choice)
		fmt.Println("------------------")

		switch choice {
		case 1:
			View.InputTransaksi()
		case 2:
			View.TampilkanDaftarTransaksi()
		case 3:
			View.HapusTransaksi()
		case 4:
			View.UpdateTransaksi()
		case 5:
			View.HitungSaldo()
		case 6:
			fmt.Println("Keluar dari aplikasi.")
			fmt.Println("...")
			fmt.Println("..")
			fmt.Println(".")
			return
		default:
			fmt.Println("Pilihan tidak valid, coba lagi.")
			fmt.Println()
		}
	}
}
