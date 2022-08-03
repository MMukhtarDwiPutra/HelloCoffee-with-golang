package handler

import (
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/backend/akun"
	"net/http"
	"github.com/gorilla/sessions"
	"os"
	"html/template"
	"log"
)

type akunHandler struct{
	akunService akun.Service
}

func NewAkunHandler(akunService akun.Service) *akunHandler{
	return &akunHandler{akunService}
}

func (h *akunHandler) LoginProcess(w http.ResponseWriter, r *http.Request){
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

func (h *akunHandler) RegisterNewAccount(w http.ResponseWriter, r *http.Request){
	username := r.FormValue("username")
	password := r.FormValue("password")
	gender := r.FormValue("gender")
	email := r.FormValue("email")

	h.akunService.CreateAccount(username, password, email, gender)
	
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *akunHandler) RegisterHandler(w http.ResponseWriter, r *http.Request){
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

func (h *akunHandler) LoginHandler(w http.ResponseWriter, r *http.Request){
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