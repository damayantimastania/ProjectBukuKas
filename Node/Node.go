package Node

type Transaksi struct {
	NoTransaksi   int
	Ket           string
	Nominal       int
	KetPembayaran string
}
type DaftarTransaksi struct {
	Data Transaksi
	Next *DaftarTransaksi
}
