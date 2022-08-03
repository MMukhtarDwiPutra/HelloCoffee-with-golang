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

func (as *AkunStore) ProcessTransaksi(w http.ResponseWriter, r *http.Request){
	id_toko := r.URL.Query()["id_toko"][0]
	id_transaksi := r.URL.Query()["id_transaksi"][0]
	pesanan := r.URL.Query()["pesanan"][0]

	db := connectDb()

	status := "Baru"
	if pesanan == "diterima"{
		status = "Success"
	}else if pesanan == "ditolak"{
		status = "Pesanan ditolak"
	}

	db.Query("UPDATE transaksi SET status_transaksi = ? WHERE id_transaksi = ?",status,id_transaksi)

	http.Redirect(w, r, "/transaksi/?id="+id_toko, http.StatusSeeOther)
}

func (as *AkunStore) CheckoutNowHandler(w http.ResponseWriter, r *http.Request){
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

	http.Redirect(w, r, "/keranjang/checkout/?id="+strconv.Itoa(id_user), http.StatusSeeOther)
}

func (as *AkunStore) TransaksiHandler (w http.ResponseWriter, r *http.Request){
	id_toko := r.URL.Query()["id"][0]
	store := sessions.NewCookieStore([]byte("super-secret"))
	session, _ := store.Get(r, "session-name")
	id_user := session.Values["id_user"]

	db := connectDb()
	rows, err := db.Query(`SELECT t.id_transaksi, t.tanggal_transaksi, dt.full_name, dt.email, dt.address, dt.city, dt.zip, dt.state, t.Qty, m.Harga, m.foto_kopi, m.nama_menu FROM detail_transaksi dt JOIN transaksi t ON t.id_detail_transaksi = dt.id_detail_transaksi JOIN menu m ON m.id_menu = t.id_menu WHERE id_toko = ? and t.status_transaksi = "Baru"`,id_toko)
	if err != nil{
		log.Fatal(err)
	}

	var data []model.Transaksi
	for rows.Next(){
		var t model.Transaksi
		rows.Scan(&t.IdTransaksi, &t.TanggalTransaksi, &t.Nama, &t.Email, &t.Address, &t.City, &t.Zip, &t.State, &t.Qty, &t.Harga, &t.Foto, &t.NamaMenu)
		t.Harga = t.Qty * t.Harga
		fmt.Println(t.IdTransaksi)
		data = append(data, t)
	}

	path, _ := os.Getwd()
	t, err := template.ParseFiles(path+`\backend\views\layout_toko.html`,path+`\backend\views\transaksi.html`)
	if err != nil{
		log.Fatal(err)
	}

	tmp := map[string]interface{}{
		"Id_user" : id_user,
		"Id_toko" : id_toko,
		"Transaksi" : data,
	}
	err = t.Execute(w, tmp)
	if err != nil{
		log.Fatal(err)
	}
}

func (as *AkunStore) HomeTokoHandler (w http.ResponseWriter, r *http.Request){
	store := sessions.NewCookieStore([]byte("super-secret"))
	session, err := store.Get(r, "session-name")
	id_user := session.Values["id_user"].(int)
	id_toko := session.Values["id_toko"].(int)

	var data []model.Menu

	db := connectDb()
	rows, _ := db.Query("SELECT id_menu, nama_menu, harga, jenis, foto_kopi FROM menu WHERE id_toko = ?",id_toko)
	for rows.Next(){
		var m model.Menu
		rows.Scan(&m.Id_menu, &m.Nama_menu, &m.Harga, &m.Jenis, &m.Foto)

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

func (as *AkunStore) LogoutHandler(w http.ResponseWriter, r *http.Request){
	store := sessions.NewCookieStore([]byte("super-secret"))
	session, _ := store.Get(r, "session-name")
	session.Values["id_user"] = -1
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (as *AkunStore) SettingHandler(w http.ResponseWriter, r *http.Request){
	id := r.URL.Query()["id"][0]

	db := connectDb()

	rows, err:= db.Query(`SELECT id_user, username, email, gender, id_toko, Foto FROM user WHERE id_user = ?`,id)
	if err != nil{
		log.Fatal(err)
	}

	var data model.User
	for rows.Next(){
		rows.Scan(&data.Id_user, &data.Username, &data.Email, &data.Gender, &data.Id_toko, &data.Foto)
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
		"Id_user" : id,
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

	rows, err := db.Query("SELECT k.id_keranjang, k.qty, k.id_user, m.harga, m.nama_menu, m.foto_kopi FROM keranjang k JOIN menu m ON k.id_menu = m.id_menu WHERE k.id_user = ?",id)
	if err != nil{
		log.Fatal(err)
	}

	var data []model.Keranjang
	i := 1
	for rows.Next(){
		var k model.Keranjang
		rows.Scan(&k.IdKeranjang, &k.Qty, &k.IdUser, &k.Harga, &k.NamaMenu, &k.Foto)
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