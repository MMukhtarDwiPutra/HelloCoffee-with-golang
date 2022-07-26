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
	"github.com/gorilla/sessions"
	"os"
)

type DataStore struct{
	Id_user int
	datastore []model.Toko
}

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

	rows, err := db.Query(`SELECT id_user FROM user WHERE email = ? and password = ?`,email, password)
	if err != nil{
		panic(err.Error())
	}

	store := sessions.NewCookieStore([]byte("super-secret"))
	session, _ := store.Get(r, "session-name")
	
	for rows.Next(){
		var id_user int
		rows.Scan(&id_user)
		session.Values["id_user"] = id_user
    	session.Save(r, w)
    	fmt.Println(id_user)
		http.Redirect(w, r, "/home",http.StatusSeeOther)
	}		

	http.Redirect(w, r, "/", http.StatusSeeOther)
	
}

func (as *AkunStore) RegisterNewAccount(w http.ResponseWriter, r *http.Request){
	username := r.FormValue("username")
	password := r.FormValue("password")
	gender := r.FormValue("gender")
	email := r.FormValue("email")

	db := connectDb()

	insert, err := db.Query("INSERT INTO user (username, password, email, gender) VALUES (? , ?, ?, ?) ", username, password, email, gender)
	if err != nil{
		log.Fatal(err)
	}
	defer insert.Close()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (as *AkunStore) RegisterHandler(w http.ResponseWriter, r *http.Request){
	path, _ := os.Getwd()
	tmplt, err := template.ParseFiles(path+`\backend\views\registration.html`)
	if err != nil{
		log.Println(err)
		http.Error(w, "Error is happening", http.StatusInternalServerError)
		return
	}

	err = tmplt.Execute(w, nil)
	if err != nil{
		log.Println(err)
		http.Error(w, "Error is happening", http.StatusInternalServerError)
		return
	}	
}

func (as *AkunStore) HomeHandler(w http.ResponseWriter, r *http.Request){
	store := sessions.NewCookieStore([]byte("super-secret"))
	session, err := store.Get(r, "session-name")
	data := DataStore{}

	data.Id_user = session.Values["id_user"].(int)

	db := connectDb()
	rows, _ := db.Query(`SELECT nama_toko, id_toko, alamat, foto, id_user FROM daftar_toko`)
	
	for rows.Next(){
		var toko model.Toko
		rows.Scan(&toko.Nama_toko, &toko.Id_toko, &toko.Alamat, &toko.Foto, &toko.Id_user)

		data.datastore = append(data.datastore, toko)
	}

	path, _ := os.Getwd()
	t, err := template.ParseFiles(path+`\backend\views\layout.html`,path+`\backend\views\home.html`)
	if err != nil{
		log.Fatal(err)
	}
    fmt.Println(data.Id_user)
    tmp := map[string]interface{}{
    	"Id_user" : data.Id_user,
    	"datastore" : data.datastore,
    }
	err = t.Execute(w,tmp)
	if err != nil {
		log.Fatal(err)
	}
}

func (as *AkunStore) LogoutHandler(w http.ResponseWriter, r *http.Request){
	store := sessions.NewCookieStore([]byte("super-secret"))
	session, _ := store.Get(r, "session-name")
	session.Values["id_user"] = -1
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (as *AkunStore) SettingHandler(w http.ResponseWriter, r *http.Request){
	id := r.URL.Query()["id"][0]

	db := connectDb()

	rows, err:= db.Query(`SELECT id_user, username, email, gender, id_toko FROM user WHERE id_user = ?`,id)
	if err != nil{
		log.Fatal(err)
	}

	var data model.User
	for rows.Next(){
		rows.Scan(&data.Id_user, &data.Username, &data.Email, &data.Gender, &data.Id_toko)
	}

	path, _ := os.Getwd()
	t, err := template.ParseFiles(path+`\backend\views\layout.html`,path+`\backend\views\akun.html`)
	if err != nil{
		log.Fatal(err)
	}

	err = t.Execute(w, data)
	if err != nil{
		log.Fatal(err)
	}
}

func (as *AkunStore) DeleteAkun (w http.ResponseWriter, r *http.Request){
	id := r.URL.Query()["id"][0]
	db := connectDb()

	_, err := db.Query("DELETE FROM user WHERE id_user = ?", id)
	if err != nil{
		log.Fatal(err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (as *AkunStore) EditAkunHandler(w http.ResponseWriter, r *http.Request){
	path, _:= os.Getwd()
	t, err := template.ParseFiles(path+`\backend\views\layout.html`,path+`\backend\views\edit_akun.html`)
	if err != nil{
		log.Fatal(err)
	}

	store := sessions.NewCookieStore([]byte("super-secret"))
	session, err := store.Get(r, "session-name")

	data := map[string]interface{}{
		"Id_user" : session.Values["id_user"].(int),
	}
	err = t.Execute(w, data)
	if err != nil{
		log.Fatal(err)
	}
}

func (as *AkunStore) EditAkun(w http.ResponseWriter, r *http.Request){

	id := r.URL.Query()["id"][0]

	nama := r.FormValue("nama")
	email := r.FormValue("email")
	gender := r.FormValue("gender")
	
	db := connectDb()

	_, err := db.Query("UPDATE user SET username = ?, email = ?, gender = ? WHERE id_user = ?",nama, email, gender, id)
	if err != nil{
		log.Fatal(err)
	}

	http.Redirect(w, r, "/akun/pengaturan/?id="+id , http.StatusSeeOther)
}

func (as *AkunStore) EditPasswordHandler(w http.ResponseWriter, r *http.Request){
	path, _:= os.Getwd()
	t, err := template.ParseFiles(path+`\backend\views\layout.html`,path+`\backend\views\edit_password.html`)
	if err != nil{
		log.Fatal(err)
	}

	store := sessions.NewCookieStore([]byte("super-secret"))
	session, err := store.Get(r, "session-name")

	data := map[string]interface{}{
		"Id_user" : session.Values["id_user"].(int),
	}
	err = t.Execute(w, data)
	if err != nil{
		log.Fatal(err)
	}
}

func (as *AkunStore) EditPassword(w http.ResponseWriter, r *http.Request){
	id := r.URL.Query()["id"][0]

	pwBaru := r.FormValue("pwconf")
	pwLama := r.FormValue("pwlama")
	db := connectDb()

	rows, err := db.Query(`SELECT * FROM user WHERE password = ? and id_user = ?`,pwLama, id)
	if err != nil{
		log.Fatal(err)
	}
	if rows.Next(){
		db.Query("UPDATE user SET password = ? WHERE id_user = ?", pwBaru, id)
	}
	http.Redirect(w, r, "/akun/pengaturan/edit/?id="+id , http.StatusSeeOther)
}