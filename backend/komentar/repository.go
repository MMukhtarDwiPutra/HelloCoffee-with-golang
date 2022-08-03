package komentar

import(
	"database/sql"
	"log"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/model"
)

type Repository interface{
	FindAllKomentar(id_menu int, id_user int) []model.Komentar
	AddKomentar(komentar string, id_user int, id_menu int)
	DeleteKomentar(id_komentar int)
}

type repository struct{
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository{
	return &repository{db}
}

func (r *repository) FindAllKomentar(id_menu int, id_user int) []model.Komentar{
	var komentars []model.Komentar
	rows, err := r.db.Query(`SELECT id_komentar, id_user, nama_komentar, isi_komentar FROM komentar WHERE id_menu = ?`, id_menu)	
	if err != nil{
		log.Fatal(err)
	}

	for rows.Next(){
		var komentar model.Komentar
		rows.Scan(&komentar.Id_komentar, &komentar.IdUser, &komentar.NamaKomentar, &komentar.IsiKomentar)
		komentar.SessionIdUser = id_user
		komentar.IdMenu = id_menu
		komentars = append(komentars, komentar)
	}

	return komentars
}

func (r *repository) AddKomentar(komentar string, id_user int, id_menu int){
	rows, err := r.db.Query("SELECT username FROM user WHERE `id_user` = ?",id_user)
	if err != nil{
		log.Fatal(err)
	}
	var nama string
	for rows.Next(){
		rows.Scan(&nama)
	}

	r.db.Query("INSERT INTO komentar (nama_komentar, isi_komentar, id_user, id_menu) VALUES (?,?,?,?)",nama,komentar,id_user,id_menu)
}

func (r *repository) DeleteKomentar(id_komentar int){
	r.db.Query("DELETE FROM komentar WHERE id_komentar = ?",id_komentar)
}