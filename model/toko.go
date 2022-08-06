package model

type Toko struct {
	Id_toko int `json:"id_toko" db:"id_toko"`
	Nama_toko string `validate:"required" json:"nama_toko" db:"nama_toko"`
	Alamat string `validate:"required" json:"alamat" db:"alamat"`
	Foto string `json:"foto" db:"foto_toko"`
	Deskripsi string `validate:"required" json:"deskripsi" db:"deskripsi"`
	JamOperasional string `validate:"required" json:"jam_operasional" db:"jam_operasional"`
}