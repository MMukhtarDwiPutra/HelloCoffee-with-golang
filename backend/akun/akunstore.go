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
	"time"
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

	rows, err := db.Query(`SELECT id_user, id_toko FROM user WHERE email = ? and password = ?`,email, password)
	if err != nil{
		panic(err.Error())
	}

	store := sessions.NewCookieStore([]byte("super-secret"))
	session, _ := store.Get(r, "session-name")
	
	id_toko := -1
	for rows.Next(){
		var id_user int
		rows.Scan(&id_user, &id_toko)
		session.Values["id_user"] = id_user
    	session.Save(r, w)
	}

	fmt.Println(id_toko)
	if(id_toko == 0){
		http.Redirect(w, r, "/home",http.StatusSeeOther)
	}else if(id_toko == -1){
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}else{
		session.Values["id_toko"] = id_toko
		session.Save(r, w)
		http.Redirect(w, r, "/home/toko", http.StatusSeeOther)
	}
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

func (as *AkunStore) DeleteMenu (w http.ResponseWriter, r *http.Request){
	id := r.URL.Query()["id"][0]

	db := connectDb()
	db.Query("DELETE FROM menu WHERE id_menu = ?",id)

	http.Redirect(w, r, "/home/toko", http.StatusSeeOther)
}

func (as *AkunStore) TambahMenuProcess (w http.ResponseWriter, r *http.Request){
	store := sessions.NewCookieStore([]byte("super-secret"))
	session, _ := store.Get(r, "session-name")
	id_toko := session.Values["id_toko"]
	nama := r.FormValue("nama_menu")
	harga := r.FormValue("harga")
	deskripsi := r.FormValue("deskripsi")
	jenis := r.FormValue("jenis")

	db := connectDb()
	db.Query("INSERT INTO menu (nama_menu, harga, deskripsi, jenis, id_toko) VALUES (?, ?, ?, ?, ?)",nama, harga, deskripsi, jenis, id_toko)

	http.Redirect(w, r, "/home/toko", http.StatusSeeOther)
}

func (as *AkunStore) EditMenuProcess (w http.ResponseWriter, r *http.Request){
	id := r.URL.Query()["id"][0]
	nama := r.FormValue("nama_menu")
	harga := r.FormValue("harga")
	deskripsi := r.FormValue("deskripsi")
	jenis := r.FormValue("jenis")

	db := connectDb()
	db.Query("UPDATE menu SET nama_menu = ?, harga = ?, deskripsi = ?, jenis = ? WHERE id_menu = ?",nama, harga, deskripsi, jenis, id)

	http.Redirect(w, r, "/home/toko", http.StatusSeeOther)
}

func (as *AkunStore) TambahMenuHandler (w http.ResponseWriter, r *http.Request){
	store := sessions.NewCookieStore([]byte("super-secret"))
	session, _ := store.Get(r, "session-name")
	data := map[string]interface{}{
		"Id_user" : session.Values["id_user"],
	}

	path, _ := os.Getwd()
	t, _ := template.ParseFiles(path+`\backend\views\layout.html`,path+`\backend\views\tambah_menu.html`)
	err := t.Execute(w, data)
	if err != nil{
		log.Fatal(err)
	}
}

func (as *AkunStore) EditMenuHandler (w http.ResponseWriter, r *http.Request){
	id := r.URL.Query()["id"][0]

	db := connectDb()
	rows, _ := db.Query("SELECT id_menu, nama_menu, harga, deskripsi, jenis FROM menu WHERE id_menu = ?",id)
	var m model.Menu
	for rows.Next(){
		rows.Scan(&m.Id_menu, &m.Nama_menu, &m.Harga, &m.Deskripsi, &m.Jenis)
	}

	data := map[string]interface{}{
		"Id_user" : id,
		"Menu" : m,
	}

	path, _ := os.Getwd()
	t, err := template.ParseFiles(path+`\backend\views\layout_toko.html`,path+`\backend\views\edit_menu.html`)
	if err != nil{
		log.Fatal(err)
	}
	err = t.Execute(w, data)
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

func (as *AkunStore) HomeTokoHandler (w http.ResponseWriter, r *http.Request){
	store := sessions.NewCookieStore([]byte("super-secret"))
	session, err := store.Get(r, "session-name")
	id_user := session.Values["id_user"].(int)
	id_toko := session.Values["id_toko"].(int)

	var data []model.Menu

	db := connectDb()
	rows, _ := db.Query("SELECT id_menu, nama_menu, harga, jenis FROM menu WHERE id_toko = ?",id_toko)
	for rows.Next(){
		var m model.Menu
		rows.Scan(&m.Id_menu, &m.Nama_menu, &m.Harga, &m.Jenis)

		data = append(data, m)
	}

	tmp := map[string]interface{}{
		"Menu" : data,
		"Id_user" : id_user,
		"Id_toko" : id_toko,
	}

	path, _ := os.Getwd()
	t, err := template.ParseFiles(path+`\backend\views\layout_toko.html`,path+`\backend\views\home_toko.html`)
	if err != nil{
		log.Fatal(err)
	}

	t.Execute(w, tmp)
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

	var transaksi []model.Transaksi
	rows, _ = db.Query("SELECT t.tanggal_transaksi, t.qty, m.harga, m.nama_menu, t.status_transaksi FROM transaksi t JOIN menu m ON t.id_menu = m.id_menu WHERE t.id_user = ?",id)
	for rows.Next(){
		var t model.Transaksi
		var harga int
		rows.Scan(&t.TanggalTransaksi, &t.Qty, &harga, &t.NamaMenu, &t.StatusTransaksi)
		t.Harga = harga * t.Qty

		transaksi = append(transaksi, t)
	}

	path, _ := os.Getwd()
	tmp := map[string]interface{}{
		"Akun" : data,
		"Transaksi" : transaksi,
	}
	if(data.Id_toko == 0){
		t, _ := template.ParseFiles(path+`\backend\views\layout.html`,path+`\backend\views\akun.html`)
		t.Execute(w, tmp)	
	}else{
		t, _ := template.ParseFiles(path+`\backend\views\layout_toko.html`,path+`\backend\views\akun.html`)
		t.Execute(w, tmp)
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

func (as *AkunStore) KeranjangHandler(w http.ResponseWriter, r *http.Request){
	id := r.URL.Query()["id"][0]

	db := connectDb()

	rows, err := db.Query("SELECT k.id_keranjang, k.qty, k.id_user, m.harga, m.nama_menu FROM keranjang k JOIN menu m ON k.id_menu = m.id_menu WHERE k.id_user = ?",id)
	if err != nil{
		log.Fatal(err)
	}

	var data []model.Keranjang
	i := 1
	for rows.Next(){
		var k model.Keranjang
		rows.Scan(&k.IdKeranjang, &k.Qty, &k.IdUser, &k.Harga, &k.NamaMenu)
		k.No = i
		i++
		data = append(data, k)
	}

	path, _ := os.Getwd()
	t, err := template.ParseFiles(path+`\backend\views\layout.html`,path+`\backend\views\keranjang.html`)
	if err != nil{
		log.Fatal(err)
	}

	tmp := map[string]interface{}{
		"Id_user" : id,
		"keranjang" : data,
	}

	err = t.Execute(w, tmp)
	if err != nil{
		log.Fatal(err)
	}
}

func (as *AkunStore) TambahKeranjang(w http.ResponseWriter, r *http.Request){
	id_menu := r.URL.Query()["id"][0]
	store := sessions.NewCookieStore([]byte("super-secret"))
	session, err := store.Get(r, "session-name")
	if err != nil{
		log.Fatal(err)
	}

	id_user := session.Values["id_user"].(int)
	qty := r.FormValue("jumlah")

	db := connectDb()
	rows, _ := db.Query("SELECT harga FROM menu WHERE id_menu = ?",id_menu)
	
	var harga int
	for rows.Next(){
		rows.Scan(&harga)
	}

	db.Query("INSERT INTO keranjang (qty, id_menu, id_user) VALUES (?, ?, ?)",qty, id_menu, id_user)

	http.Redirect(w, r, "/menu", http.StatusSeeOther)
}

func (as *AkunStore) HapusKeranjang(w http.ResponseWriter, r *http.Request){
	id := r.URL.Query()["id"][0]

	db := connectDb()

	db.Query("DELETE FROM keranjang WHERE id_user = ?",id)

	http.Redirect(w,r,"/menu", http.StatusSeeOther)
}

func (as *AkunStore) CheckoutHandler (w http.ResponseWriter, r *http.Request){
	path, _ := os.Getwd()
	t, err := template.ParseFiles(path+`\backend\views\layout.html`,path+`\backend\views\checkout.html`)
	if err != nil{
		log.Fatal(err)
	}

	id := r.URL.Query()["id"][0]
	db := connectDb()


	rows, err := db.Query("SELECT k.id_keranjang, k.qty, k.id_user, m.harga, m.nama_menu FROM keranjang k JOIN menu m ON k.id_menu = m.id_menu WHERE k.id_user = ?",id)
	totalAll := 0
	var data []model.Keranjang
	for rows.Next(){
		var k model.Keranjang
		rows.Scan(&k.IdKeranjang, &k.Qty, &k.IdUser, &k.Harga, &k.NamaMenu)
		k.Total = k.Qty * k.Harga
		data = append(data, k)
		totalAll = totalAll + k.Total
	}

	tmp := map[string]interface{}{
		"Id_user" : id,
		"Keranjang" : data,
		"TotalAll" : totalAll,
	}

	err = t.Execute(w, tmp)
	if err != nil{
		log.Fatal(err)
	}
}

func (as *AkunStore) CheckoutProcess (w http.ResponseWriter, r *http.Request){
	id := r.URL.Query()["id"][0]
	fullName := r.FormValue("fname")
	email := r.FormValue("email")
	address := r.FormValue("address")
	city := r.FormValue("city")
	zip := r.FormValue("zip")
	state := r.FormValue("state")

	db := connectDb()
	db.Query("INSERT INTO detail_transaksi (full_name, email, address, city, zip, state) VALUES (?, ?, ?, ?, ?, ?) ",fullName, email, address, city, zip, state)
	row, _ := db.Query("SELECT id_detail_transaksi FROM detail_transaksi ORDER BY id_detail_transaksi DESC LIMIT 1")

	var id_detail_transaksi int
	for row.Next(){
		row.Scan(&id_detail_transaksi)
		fmt.Println(id_detail_transaksi)
	}

	currentTime := time.Now()
	tanggal := currentTime.Format("2006-01-02")

	rows, _ := db.Query("SELECT id_keranjang, qty, id_user, id_menu FROM keranjang WHERE id_user = ?",id)
	for rows.Next(){
		var k model.Keranjang
		rows.Scan(&k.IdKeranjang, &k.Qty, &k.IdUser, &k.IdMenu)
		db.Query(`INSERT INTO transaksi (tanggal_transaksi, qty, id_user, id_menu, id_detail_transaksi, status_transaksi) VALUES (?, ?, ?, ?, ?, "Baru")`, tanggal, k.Qty, k.IdUser, k.IdMenu, id_detail_transaksi)
		db.Query("DELETE FROM keranjang WHERE id_keranjang = ?", k.IdKeranjang)
	}

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}