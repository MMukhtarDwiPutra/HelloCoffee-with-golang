package toko

import (
	"database/sql"
	"log"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/model"
)

type Repository interface{
	FindAllToko() []model.Toko
	FindOneToko(id int) model.Toko
}

type repository struct{
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository{
	return &repository{db}
}

func (r *repository) FindAllToko() []model.Toko{
	var tokos []model.Toko
	rows, err := r.db.Query(`SELECT nama_toko, id_toko, alamat, foto_toko, id_user FROM daftar_toko`)
	if err != nil{
		log.Fatal(err)
	}

	for rows.Next(){
		var toko model.Toko
		rows.Scan(&toko.Nama_toko, &toko.Id_toko, &toko.Alamat, &toko.Foto, &toko.Id_user)

		tokos = append(tokos, toko)
	}
	return tokos
}

func (r *repository) FindOneToko(id int) model.Toko{
	rows, err := r.db.Query(`SELECT id_toko, nama_toko, alamat, foto_toko, deskripsi, jam_operasional FROM daftar_toko WHERE id_toko = ?`,id)
	if err != nil{
		log.Fatal(err)
	}

	var toko model.Toko
	for rows.Next(){
		rows.Scan(&toko.Id_toko, &toko.Nama_toko, &toko.Alamat, &toko.Foto, &toko.Deskripsi, &toko.JamOperasional)
	}
	return toko
}