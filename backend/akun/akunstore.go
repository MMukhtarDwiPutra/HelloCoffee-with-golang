package akun

import (
	// "fmt"
	"log"
	"net/http"
	"database/sql"
	"html/template"
	// "strconv"
	// "github.com/gorilla/mux"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/model"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

type AkunStore struct{
	data []model.Akun
}

func NewDataAkun() *AkunStore{
	newData := make([]model.Akun, 0)

	return &AkunStore{
		data:newData,
	}
}

func connectDb() *sql.DB{
	db, err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/hellocoffee")
	if(err != nil){
		log.Fatal(err)
	}
	return db
}

func (as *AkunStore) LoginHandler(w http.ResponseWriter, r *http.Request){
	email := r.FormValue("email")
	password := r.FormValue("password")

	db := connectDb()

	rows, err := db.Query(`SELECT email FROM user WHERE email = ? and password = ?`,email, password)
	if err != nil{
		panic(err.Error())
	}

	i := 0
	for rows.Next(){
		i++
	}
	if(i == 0){
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}else{
		http.Redirect(w, r, "/home",http.StatusSeeOther)
	}
}

func (as *AkunStore) RegisterHandler(w http.ResponseWriter, r *http.Request){}

func (as *AkunStore) HomeHandler(w http.ResponseWriter, r *http.Request){
	var data []model.Toko

	db := connectDb()
	rows, _ := db.Query(`SELECT nama_toko, id_toko, alamat, foto, id_user FROM daftar_toko`)
	
	for rows.Next(){
		var toko model.Toko
		rows.Scan(&toko.Nama_toko, &toko.Id_toko, &toko.Alamat, &toko.Foto, &toko.Id_user)

		data = append(data, toko)
	}

	path, _ := os.Getwd()
	t, err := template.ParseFiles(path+`\backend\views\layout.html`,path+`\backend\views\home.html`)
	if err != nil{
		log.Fatal(err)
	}
	err = t.Execute(w,data)
	if err != nil {
		log.Fatal(err)
	}
}