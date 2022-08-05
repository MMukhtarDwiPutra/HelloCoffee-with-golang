package akun

import (
	"database/sql"
	"log"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/model"
)

type Repository interface{
	FindOneUser(email string, password string) (int, int)
	CreateAccount(username string, password string, email string, nama string, gender string, foto string, id_toko int)
	GetOneAkun(id_user int) model.User
	DeleteAkun(id_user int)
	UpdateAkun(nama string, email string, gender string, id_user int)
	UpdatePassword(pwLama string, pwBaru string, id_user int)
	GetLastUser() model.User
	FindAllUser() []model.User
	UpdateAkunAPI(username string, password string, email string, gender string, nama string, foto string, id_user int)
} 

type repository struct{
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository{
	return &repository{db}
}

func (r *repository) FindAllUser() []model.User{
	var u []model.User

	rows, err := r.db.Query("SELECT id_user, username, password, nama, email, gender, Foto, id_toko FROM user")
	if err != nil{
		log.Fatal(err)
	}

	for rows.Next(){
		var user model.User
		rows.Scan(&user.Id_user, &user.Username, &user.Password, &user.Nama, &user.Email, &user.Gender, &user.Foto, &user.Id_toko)
		u = append(u, user)
	}

	return u
}

func (r *repository) FindOneUser(email string, password string) (int, int){
	rows, err := r.db.Query(`SELECT id_user, id_toko FROM user WHERE email = ? and password = ?`,email, password)
	if err != nil{
		log.Fatal(err)
	}

	id_user := -1
	id_toko := 9999
	for rows.Next(){
		rows.Scan(&id_user, &id_toko)
	}
	return id_user, id_toko
}

func (r *repository) CreateAccount(username string, password string, email string, nama string, gender string, foto string, id_toko int){
	insert, err := r.db.Query(`INSERT INTO user (username, password, email, nama, gender, Foto, id_toko) VALUES (?, ?, ?, ?, ?, ?, ?) `, username, password, email, nama, gender, foto, id_toko)
	if err != nil{
		log.Fatal(err)
	}
	defer insert.Close()
}

func (r *repository) GetOneAkun(id_user int) model.User{
	rows, err:= r.db.Query(`SELECT id_user, password, username, email, gender, nama, Foto, id_toko FROM user WHERE id_user = ?`,id_user)
	if err != nil{
		log.Fatal(err)
	}

	var data model.User
	for rows.Next(){
		rows.Scan(&data.Id_user, &data.Password, &data.Username, &data.Email, &data.Gender, &data.Nama, &data.Foto, &data.Id_toko)
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
	_, err := r.db.Query("UPDATE user SET nama = ?, email = ?, gender = ? WHERE id_user = ?",nama, email, gender, id_user)
	if err != nil{
		log.Fatal(err)
	}
}

func (r *repository) UpdateAkunAPI(username string, password string, email string, gender string, nama string, foto string, id_user int){
	_, err := r.db.Query("UPDATE user SET username = ?, password = ?, email = ?, gender = ?, nama = ?, foto = ? WHERE id_user = ?",nama, password, email, gender, nama, foto, id_user)
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

func (r *repository) GetLastUser() model.User{
	var user model.User

	rows, err := r.db.Query("SELECT id_user, username, password, nama, email, gender, Foto, id_toko FROM user ORDER BY id_user DESC LIMIT 1")
	if err != nil{
		log.Fatal(err)
	}

	for rows.Next(){
		rows.Scan(&user.Id_user, &user.Username, &user.Password, &user.Nama, &user.Email, &user.Gender, &user.Foto, &user.Id_toko)
	}

	return user
}