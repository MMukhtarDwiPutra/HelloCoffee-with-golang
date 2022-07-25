package model

type Toko struct {
	Id_toko int `json:"id_toko" db:"id_toko"`
	Nama_toko string `json:"nama_toko" db:"nama_toko"`
	Alamat string `json:"alamat" db:"alamat"`
	Id_user int `json:"id_user" db:"id_user"`
	Foto string `json:"foto" db:"foto"`
}