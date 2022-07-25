package model

type Toko struct {
	id_toko int `json:"id_toko" db:"id_toko"`
	nama_toko string `json:"nama_toko" db:"nama_toko"`
	alamat string `json:"alamat" db:"alamat"`
	id_user int `json:"id_user" db:"id_user"`
	foto string `json:"foto" db:"foto"`
}