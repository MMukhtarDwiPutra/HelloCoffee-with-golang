package model

type Akun struct{
	nama string `json:"nama" db:"nama"`
	password string `json:"password" db:"password"`
}