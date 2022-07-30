package model

type Komentar struct{
	NamaKomentar string `db:"nama_komentar"`
	IdUser int `db:"id_user"`
	IsiKomentar string `db:"isi_komentar"`
	SessionIdUser int
	Id_komentar int `db:"id_komentar"`
	IdMenu int
}