package model

type Transaksi struct{
	IdTransaksi int `db:"id_transaksi"`
	TanggalTransaksi string `db:"tanggal_transaksi"`
	Qty int `db:"qty"`
	NamaMenu string `db:"nama_menu"`
	Harga int
	StatusTransaksi string `db:"status_transaksi"`
	Nama string
	Email string
	Address string
	State string
	Zip string
	City string
	Foto string
}