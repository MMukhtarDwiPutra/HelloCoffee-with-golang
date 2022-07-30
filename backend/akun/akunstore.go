package akun

import (
	"fmt"
	"log"
	"net/http"
	"database/sql"
	"html/template"
	"strconv"
	// "strconv"
	// "github.com/gorilla/mux"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
	"os"
)

type DataTokoStore struct{
	Id_user int
	datastore []model.Toko
}

type MenuStore struct{
	Id_user int
	datastore []model.Menu
	komentar []model.Komentar
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

func (as *AkunStore) DetailMenu(w http.ResponseWriter, r *http.Request){
	store := sessions.NewCookieStore([]byte("super-secret"))
	session, err := store.Get(r, "session-name")
	data := MenuStore{}

	id_menu := r.URL.Query()["id"][0]
	data.Id_user = session.Values["id_user"].(int)
	db := connectDb()


	rows, _ := db.Query(`SELECT m.id_menu, m.nama_menu, m.harga, m.deskripsi, m.jenis, m.foto_kopi, dt.nama_toko 
		FROM menu as m 
		JOIN daftar_toko as dt ON dt.id_toko = m.id_toko 
		WHERE m.id_menu = ?`,id_menu)
	for rows.Next(){
		var menu model.Menu
		rows.Scan(&menu.Id_menu, &menu.Nama_menu, &menu.Harga, &menu.Deskripsi, &menu.Jenis, &menu.Foto, &menu.Nama_toko)

		data.datastore = append(data.datastore, menu)
	}

	rows, _ = db.Query(`SELECT id_komentar, id_user, nama_komentar, isi_komentar FROM komentar WHERE id_menu = ?`, id_menu)	

	for rows.Next(){
		var komentar model.Komentar
		rows.Scan(&komentar.Id_komentar, &komentar.IdUser, &komentar.NamaKomentar, &komentar.IsiKomentar)
		komentar.SessionIdUser = session.Values["id_user"].(int)
		komentar.IdMenu, _ = strconv.Atoi(id_menu)
		data.komentar = append(data.komentar, komentar)
	}

	path, _ := os.Getwd()
	t, err := template.ParseFiles(path+`\backend\views\layout.html`,path+`\backend\views\detail_menu.html`)
	if err != nil{
		log.Fatal(err)
	}

	tmp := map[string]interface{}{
    	"Id_user" : data.Id_user,
    	"datastore" : data.datastore,
    	"komentar" : data.komentar,
    	"id_menu" : id_menu,
    }
	err = t.Execute(w, tmp)
	if err != nil{
		log.Fatal(err)
	}
}

func (as *AkunStore) MenuHandler(w http.ResponseWriter, r *http.Request){
	store := sessions.NewCookieStore([]byte("super-secret"))
	session, err := store.Get(r, "session-name")
	data := MenuStore{}

	data.Id_user = session.Values["id_user"].(int)
	db := connectDb()

	rows, _ := db.Query(`SELECT m.id_menu, m.nama_menu, m.harga, m.deskripsi, m.jenis, m.foto_kopi, dt.nama_toko 
		FROM menu as m 
		JOIN daftar_toko as dt ON dt.id_toko = m.id_toko`)
	for rows.Next(){
		var menu model.Menu
		rows.Scan(&menu.Id_menu, &menu.Nama_menu, &menu.Harga, &menu.Deskripsi, &menu.Jenis, &menu.Foto, &menu.Nama_toko)

		data.datastore = append(data.datastore, menu)
	}

	path, _ := os.Getwd()
	t, err := template.ParseFiles(path+`\backend\views\layout.html`,path+`\backend\views\menu.html`)
	if err != nil{
		log.Fatal(err)
	}

	tmp := map[string]interface{}{
    	"Id_user" : data.Id_user,
    	"datastore" : data.datastore,
    }
	t.Execute(w, tmp)
}

func (as *AkunStore) DetailTokoHandler(w http.ResponseWriter, r *http.Request){
	db := connectDb()
	id := r.URL.Query()["id"][0]

	rows, err := db.Query(`SELECT id_toko, nama_toko, alamat, foto_toko, deskripsi, jam_operasional FROM daftar_toko WHERE id_toko = ?`,id)
	if err != nil{
		log.Fatal(err)
	}

	var toko model.Toko
	for rows.Next(){
		rows.Scan(&toko.Id_toko, &toko.Nama_toko, &toko.Alamat, &toko.Foto, &toko.Deskripsi, &toko.JamOperasional)
		fmt.Println(toko.Nama_toko)
	}

	path, _ := os.Getwd()
	t, err := template.ParseFiles(path+`\backend\views\layout.html`,path+`\backend\views\detail_toko.html`)
	if err != nil{
		log.Fatal(err)
	}

	t.Execute(w, toko)
}

func (as *AkunStore) HomeHandler(w http.ResponseWriter, r *http.Request){
	store := sessions.NewCookieStore([]byte("super-secret"))
	session, err := store.Get(r, "session-name")
	data := DataTokoStore{}

	data.Id_user = session.Values["id_user"].(int)

	db := connectDb()
	rows, _ := db.Query(`SELECT nama_toko, id_toko, alamat, foto_toko, id_user FROM daftar_toko`)
	
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

func (as *AkunStore) TambahKomentar(w http.ResponseWriter, r *http.Request){
	store := sessions.NewCookieStore([]byte("super-secret"))
	session, err := store.Get(r, "session-name")
	if err != nil{
		log.Fatal(err)
	}

	id_user := session.Values["id_user"].(int)
	id_menu := r.URL.Query()["id"][0]
	komentar := r.FormValue("isi_komentar")

	db := connectDb()
	rows, err := db.Query("SELECT username FROM user WHERE `id_user` = ?",id_user)
	var nama string
	for rows.Next(){
		rows.Scan(&nama)
	}

	db.Query("INSERT INTO komentar (nama_komentar, isi_komentar, id_user, id_menu) VALUES (?,?,?,?)",nama,komentar,id_user,id_menu)

	http.Redirect(w, r, "/menu/detail/?id="+id_menu, http.StatusSeeOther)
}
func (as *AkunStore) HapusKomentar(w http.ResponseWriter, r *http.Request){
	id_komentar := r.URL.Query()["id_komentar"][0]
	id_menu := r.URL.Query()["id_menu"][0]

	db := connectDb()
	db.Query("DELETE FROM komentar WHERE id_komentar = ?",id_komentar)

	http.Redirect(w, r, "/menu/detail/?id="+id_menu, http.StatusSeeOther)
}