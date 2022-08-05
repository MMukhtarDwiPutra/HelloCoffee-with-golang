package model

type Menu struct{
	Id_menu int `json:"id_menu" db:"id_menu"`
	Nama_menu string `validate:"required" json:"nama_menu" db:"nama_menu"`
	Harga int `validate:"required" json:"harga" db:"harga"`
	Deskripsi string `validate:"required" json:"deskripsi" db:"deskripsi"`
	Jenis string `validate:"required" json:"jenis" db:"jenis"`
	Foto string `json:"foto_kopi" db:"foto_kopi"`
	Nama_toko string `json:"nama_toko" db:"nama_toko"`
	Id_toko int `validate:"required" json:"id_toko" db:"id_toko"`
}