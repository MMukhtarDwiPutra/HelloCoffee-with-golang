package model

type Menu struct{
	Id_menu int `db:"id_menu"`
	Nama_menu string `db:"nama_menu"`
	Harga int `db:"harga"`
	Deskripsi string `db:"deskripsi"`
	Jenis string `db:"jenis"`
	Foto string `db:"foto_kopi"`
	Nama_toko string `db:"nama_toko"`
	Alamat string `db:"alamat"`
}