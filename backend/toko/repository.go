package toko

import (
	"database/sql"
	"log"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/model"
)

type Repository interface{
	FindAllToko() []model.Toko
	FindOneToko(id int) model.Toko
	CreateToko(toko model.Toko)
	GetLastToko() model.Toko
	UpdateToko(t model.Toko, id_toko int)
	DeleteToko(id_toko int)
}

type repository struct{
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository{
	return &repository{db}
}

func (r *repository) CreateToko(t model.Toko){
	_, err := r.db.Query("INSERT INTO daftar_toko (nama_toko, alamat, foto_toko, deskripsi, jam_operasional) VALUES (?, ?, ?, ?, ?) ", t.Nama_toko, t.Alamat, t.Foto, t.Deskripsi, t.JamOperasional)
	if err != nil{
		log.Fatal(err)
	}
}

func (r *repository) GetLastToko() model.Toko{
	var t model.Toko

	rows, err := r.db.Query("SELECT id_toko, nama_toko, alamat, foto_toko, deskripsi, jam_operasional FROM daftar_toko ORDER BY id_toko DESC LIMIT 1")
	if err != nil{
		log.Fatal(err)
	}
	for rows.Next(){
		rows.Scan(&t.Id_toko, &t.Nama_toko, &t.Alamat, &t.Foto, &t.Deskripsi, &t.JamOperasional)
	}

	return t
}

func (r *repository) FindAllToko() []model.Toko{
	var tokos []model.Toko
	rows, err := r.db.Query(`SELECT nama_toko, id_toko, alamat, foto_toko, deskripsi, jam_operasional FROM daftar_toko`)
	if err != nil{
		log.Fatal(err)
	}

	for rows.Next(){
		var toko model.Toko
		rows.Scan(&toko.Nama_toko, &toko.Id_toko, &toko.Alamat, &toko.Foto, &toko.Deskripsi, &toko.JamOperasional)

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

func (r *repository) UpdateToko(t model.Toko, id_toko int){
	r.db.Query("UPDATE daftar_toko SET nama_toko = ?, alamat = ?, deskripsi = ?, jam_operasional = ? WHERE id_toko = ?", t.Nama_toko, t.Alamat, t.Deskripsi, t.JamOperasional, id_toko)
} 

func (r *repository) DeleteToko(id_toko int){
	r.db.Query("DELETE FROM daftar_toko WHERE id_toko = ?", id_toko)
}