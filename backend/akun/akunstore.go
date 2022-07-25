package akun

import (
	"fmt"
	"log"
	"net/http"
	"database/sql"
	"html/template"
	// "strconv"
	// "github.com/gorilla/mux"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/model"
	_ "github.com/go-sql-driver/mysql"
)

type AkunStore struct{
	data []model.Akun
}

type TokoStore struct{
	data []model.Toko
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
	data := TokoStore{}
	path, _ := os.Getwd()
	t, err := template.ParseFiles(path+`\views\home.html`)

	db := connectDb()
	rows, err := db.Query(`SELECT * FROM daftar_toko`)
	for rows.Next(){
		data = append(data, rows)
	}

	fmt.Println(err)
	fmt.Println(t.Execute(w,data))
}