package akun

import (
	"database/sql"
	"log"
)

type Repository interface{
	FindOneUser(email string, password string) (int, int)
	CreateAccount(username string, password string, email string, gender string)
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
		var id_user int
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