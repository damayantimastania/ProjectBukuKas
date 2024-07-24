package Model

import (
	"ProjectBukuKas/Database"
	"ProjectBukuKas/Node"
)

func InputTransaksi(noTransaksi int, ket string, nominal int, ketPembayaran string) {
	newDataTransaksi := &Node.DaftarTransaksi{
		Data: Node.Transaksi{
			NoTransaksi:   noTransaksi,
			Ket:           ket,
			Nominal:       nominal,
			KetPembayaran: ketPembayaran,
		},
		Next: nil,
	}
	if Database.DBTrans == nil {
		Database.DBTrans = newDataTransaksi
		return
	}
	current := Database.DBTrans
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newDataTransaksi
}
func TampilDaftarTransaksi() *Node.DaftarTransaksi {
	if Database.DBTrans == nil {
		return nil
	}
	return Database.DBTrans
}
func HapusTransaksi(noTransaksi int) bool {
	if Database.DBTrans == nil {
		return false
	}

	prev := &Database.DBTrans
	curr := Database.DBTrans

	for curr != nil {
		if curr.Data.NoTransaksi == noTransaksi {
			*prev = curr.Next
			return true
		}
		prev = &curr.Next
		curr = curr.Next
	}
	return false
}
func CariTransaksi(noTransaksi int) *Node.DaftarTransaksi {
	curr := Database.DBTrans
	for curr != nil {
		if curr.Data.NoTransaksi == noTransaksi {
			return curr
		}
		curr = curr.Next
	}
	return nil
}
func UpdateTransaksi(transaksi Node.Transaksi) bool {
	addr := CariTransaksi(transaksi.NoTransaksi)
	if addr == nil {
		return false
	}
	addr.Data.Ket = transaksi.Ket
	addr.Data.Nominal = transaksi.Nominal
	addr.Data.KetPembayaran = transaksi.KetPembayaran
	return true
}
