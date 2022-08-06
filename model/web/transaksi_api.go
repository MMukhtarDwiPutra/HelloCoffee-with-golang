package web

type TransaksiResponse struct{
	IdTransaksi int `json:"id_transaksi" db:"id_transaksi"`
	TanggalTransaksi string `json:"tanggal_transaksi" db:"tanggal_transaksi"`
	Qty int `json:"qty" db:"qty"`
	NamaMenu string `json:"nama_menu" db:"nama_menu"`
	Harga int `json:"harga"`
	StatusTransaksi string `json:"status_transaksi" db:"status_transaksi"`
	DetailTransaksi DetailTransaksiResponse `json:"detail_transaksi"`
}

type DetailTransaksiResponse struct{
	Nama string `json:"nama"`
	Email string `json:"email"`
	Address string `json:"address"`
	State string `json:"state"`
	Zip string `json:"zip"`
	City string `json:"city"`
}