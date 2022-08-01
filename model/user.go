package model

type User struct{
	Id_user int `json:"id_user" db:"id_user"`
	Username string `json:"username" db:"username"`
	Email string `json:"email" db:"email"`
	Gender string `json:"gender" db:"gender"`
	Id_toko int `json:"id_toko" db:"id_toko"`
	Foto string `db:"foto_kopi"`
}