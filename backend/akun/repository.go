package akun

import (
	"database/sql"
	"log"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/model"
)

type Repository interface{
	FindOneUser(email string, password string) (int, int)
	CreateAccount(username string, password string, email string, gender string)
	GetOneAkun(id_user int) model.User
	DeleteAkun(id_user int)
	UpdateAkun(nama string, email string, gender string, id_user int)
	UpdatePassword(pwLama string, pwBaru string, id_user int)
} 

type repository struct{
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository{
	return &repository{db}
}

func (r *repository) FindOneUser(email string, password string) (int, int){
	rows, err := r.db.Query(`SELECT id_user, id_toko FROM user WHERE email = ? and password = ?`,email, password)
	if err != nil{
		log.Fatal(err)
	}

	var id_user int
	id_toko := -1
	for rows.Next(){
		rows.Scan(&id_user, &id_toko)
	}
	return id_user, id_toko
}

func (r *repository) CreateAccount(username string, password string, email string, gender string){
	insert, err := r.db.Query(`INSERT INTO user (username, password, email, gender, Foto) VALUES (?, ?, ?, ?, "kopi.png") `, username, password, email, gender)
	if err != nil{
		log.Fatal(err)
	}
	defer insert.Close()
}

func (r *repository) GetOneAkun(id_user int) model.User{
	rows, err:= r.db.Query(`SELECT id_user, username, email, gender, id_toko, Foto FROM user WHERE id_user = ?`,id_user)
	if err != nil{
		log.Fatal(err)
	}

	var data model.User
	for rows.Next(){
		rows.Scan(&data.Id_user, &data.Username, &data.Email, &data.Gender, &data.Id_toko, &data.Foto)
	}

	return data
}

func (r *repository) DeleteAkun(id_user int){
	_, err := r.db.Query("DELETE FROM user WHERE id_user = ?", id_user)
	if err != nil{
		log.Fatal(err)
	}
}

func (r *repository) UpdateAkun(nama string, email string, gender string, id_user int){
	_, err := r.db.Query("UPDATE user SET username = ?, email = ?, gender = ? WHERE id_user = ?",nama, email, gender, id_user)
	if err != nil{
		log.Fatal(err)
	}
}

func (r *repository) UpdatePassword(pwLama string, pwBaru string, id_user int){
	rows, err := r.db.Query(`SELECT * FROM user WHERE password = ? and id_user = ?`,pwLama, id_user)
	if err != nil{
		log.Fatal(err)
	}
	if rows.Next(){
		r.db.Query("UPDATE user SET password = ? WHERE id_user = ?", pwBaru, id_user)
	}
}