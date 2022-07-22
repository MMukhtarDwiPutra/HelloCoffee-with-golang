package model

type akun struct{
	nama string `json:"nama" db:"nama"`
	password string `json:"password" db:"password"`
}