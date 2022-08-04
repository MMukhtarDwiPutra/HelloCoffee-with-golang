package handler

import (
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/backend/akun"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/backend/transaksi"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/backend/komentar"
	"net/http"
	"github.com/gorilla/sessions"
	"os"
	"html/template"
	"log"
	"strconv"
)

type AkunHandler struct{
	akunService akun.Service
	transaksiService transaksi.Service
	komentarService komentar.Service
}

func NewAkunHandler(akunService akun.Service, transaksiService transaksi.Service, komentarService komentar.Service) *AkunHandler{
	return &AkunHandler{akunService, transaksiService, komentarService}
}

func (h *AkunHandler) LoginProcess(w http.ResponseWriter, r *http.Request){
	email := r.FormValue("email")
	password := r.FormValue("password")

	store := sessions.NewCookieStore([]byte("super-secret"))
	session, _ := store.Get(r, "session-name")
	
	id_user, id_toko := h.akunService.FindOneUser(email, password)

	session.Values["id_user"] = id_user
	session.Save(r, w)
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

func (h *AkunHandler) RegisterNewAccount(w http.ResponseWriter, r *http.Request){
	username := r.FormValue("username")
	password := r.FormValue("password")
	gender := r.FormValue("gender")
	email := r.FormValue("email")

	h.akunService.CreateAccount(username, password, email, gender)
	
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *AkunHandler) RegisterHandler(w http.ResponseWriter, r *http.Request){
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

func (h *AkunHandler) LoginHandler(w http.ResponseWriter, r *http.Request){
	path, _ := os.Getwd()
	tmplt, err := template.ParseFiles(path+`\backend\views\login.html`)
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

func (h *AkunHandler) SettingHandler(w http.ResponseWriter, r *http.Request){
	store := sessions.NewCookieStore([]byte("super-secret"))
	session, _ := store.Get(r, "session-name")

	id := r.URL.Query()["id"][0]
	id_user, _ := strconv.Atoi(id)

	data := h.akunService.GetOneAkun(id_user)

	transaksi := h.transaksiService.GetAllTransaksi(id_user)
	
	path, _ := os.Getwd()
	tmp := map[string]interface{}{
		"Id_toko" : session.Values["id_toko"],
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

func (h *AkunHandler) LogoutHandler(w http.ResponseWriter, r *http.Request){
	store := sessions.NewCookieStore([]byte("super-secret"))
	session, _ := store.Get(r, "session-name")
	session.Values["id_user"] = -1
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *AkunHandler) DeleteAkun (w http.ResponseWriter, r *http.Request){
	id := r.URL.Query()["id"][0]
	id_user, _ := strconv.Atoi(id)

	h.akunService.DeleteAkun(id_user)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *AkunHandler) EditAkunHandler(w http.ResponseWriter, r *http.Request){
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

func (h *AkunHandler) EditAkun(w http.ResponseWriter, r *http.Request){
	id := r.URL.Query()["id"][0]
	id_user, _ := strconv.Atoi(id)
	nama := r.FormValue("nama")
	email := r.FormValue("email")
	gender := r.FormValue("gender")

	h.akunService.UpdateAkun(nama, email, gender, id_user)

	http.Redirect(w, r, "/akun/pengaturan/?id="+id , http.StatusSeeOther)
}

func (h *AkunHandler) EditPasswordHandler(w http.ResponseWriter, r *http.Request){
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

func (h *AkunHandler) EditPassword(w http.ResponseWriter, r *http.Request){
	id := r.URL.Query()["id"][0]
	id_user, _ := strconv.Atoi(id)

	pwBaru := r.FormValue("pwconf")
	pwLama := r.FormValue("pwlama")

	h.akunService.UpdatePassword(pwLama, pwBaru, id_user)
	
	http.Redirect(w, r, "/akun/pengaturan/edit/?id="+id , http.StatusSeeOther)
}

func (h *AkunHandler) TambahKomentar(w http.ResponseWriter, r *http.Request){
	store := sessions.NewCookieStore([]byte("super-secret"))
	session, err := store.Get(r, "session-name")
	if err != nil{
		log.Fatal(err)
	}

	id_user := session.Values["id_user"].(int)
	id_menu_string := r.URL.Query()["id"][0]
	id_menu, _ := strconv.Atoi(id_menu_string)
	komentar := r.FormValue("isi_komentar")

	h.komentarService.AddKomentar(komentar, id_user, id_menu)

	http.Redirect(w, r, "/menu/detail/?id="+id_menu_string, http.StatusSeeOther)
}

func (h *AkunHandler) HapusKomentar(w http.ResponseWriter, r *http.Request){
	id_komentar_string := r.URL.Query()["id_komentar"][0]
	id_komentar, _ := strconv.Atoi(id_komentar_string)
	id_menu := r.URL.Query()["id_menu"][0]

	h.komentarService.DeleteKomentar(id_komentar)

	http.Redirect(w, r, "/menu/detail/?id="+id_menu, http.StatusSeeOther)
}

func (h *AkunHandler) TransaksiHandler (w http.ResponseWriter, r *http.Request){
	id_toko_string := r.URL.Query()["id"][0]
	id_toko, _ := strconv.Atoi(id_toko_string)
	
	store := sessions.NewCookieStore([]byte("super-secret"))
	session, _ := store.Get(r, "session-name")
	id_user := session.Values["id_user"]

	data := h.transaksiService.GetTransaksiToko(id_toko)

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

func (h *AkunHandler) ProcessTransaksi(w http.ResponseWriter, r *http.Request){
	id_toko := r.URL.Query()["id_toko"][0]
	id_transaksi_string := r.URL.Query()["id_transaksi"][0]
	id_transaksi, _ := strconv.Atoi(id_transaksi_string)
	pesanan := r.URL.Query()["pesanan"][0]

	status := "Baru"
	if pesanan == "diterima"{
		status = "Success"
	}else if pesanan == "ditolak"{
		status = "Pesanan ditolak"
	}

	h.transaksiService.UpdateStatusTransaksi(status, id_transaksi)

	http.Redirect(w, r, "/transaksi/?id="+id_toko, http.StatusSeeOther)
}