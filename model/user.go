package model

type User struct{
	Id_user int `json:"id_user" db:"id_user"`
	Password string `validate:"required" json:"password"`
	Username string `validate:"required" json:"username" db:"username"`
	Email string `validate:"required" json:"email" db:"email"`
	Nama string `validate:"required" json:"nama"`
	Gender string `validate:"required" json:"gender" db:"gender"`
	Id_toko int `validate:"required" json:"id_toko" db:"id_toko"`
	Foto string `validate:"required" json:"foto" db:"Foto"`
}