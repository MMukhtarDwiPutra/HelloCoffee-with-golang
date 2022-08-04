package model

type Menu struct{
	Id_menu int `json:"id_menu" db:"id_menu"`
	Nama_menu string `json:"nama_menu" db:"nama_menu"`
	Harga int `json:"harga" db:"harga"`
	Deskripsi string `json:"deskripsi" db:"deskripsi"`
	Jenis string `json:"jenis" db:"jenis"`
	Foto string `json:"foto_kopi" db:"foto_kopi"`
	Nama_toko string `json:"nama_toko" db:"nama_toko"`
	Id_toko int `json:"id_toko" db:"id_toko"`
}