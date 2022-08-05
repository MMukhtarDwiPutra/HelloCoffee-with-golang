package model

type Akun struct{
	Nama string `validate:"required" json:"nama" db:"nama"`
	Password string `validate:"required" json:"password" db:"password"`
	ID_toko int `validate:"required" json"id_toko" db:"id_toko"`
	Email string `validate:"required" json:"email" db:"email"`
	Gender string `validate:"required" json:"gender" db:"gender"`
	Foto string `validate:"required" json:"foto_kopi" db:"foto_kopi"`
}