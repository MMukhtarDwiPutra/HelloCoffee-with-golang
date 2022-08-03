package handler

import (
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/backend/toko"
	"net/http"
	"log"
	"os"
	"github.com/gorilla/sessions"
	"html/template"
	"strconv"
)

type tokoHandler struct{
	tokoService toko.Service
}

func NewTokoHandler(tokoService toko.Service) *tokoHandler{
	return &tokoHandler{tokoService}
}

func (h *tokoHandler) HomeHandler(w http.ResponseWriter, r *http.Request){
	store := sessions.NewCookieStore([]byte("super-secret"))
	session, err := store.Get(r, "session-name")

	id_user := session.Values["id_user"].(int)

	datastore := h.tokoService.FindAllToko()

	path, _ := os.Getwd()
	t, err := template.ParseFiles(path+`\backend\views\layout.html`,path+`\backend\views\home.html`)
	if err != nil{
		log.Fatal(err)
	}

    tmp := map[string]interface{}{
    	"Id_user" : id_user,
    	"datastore" : datastore,
    }
	err = t.Execute(w,tmp)
	if err != nil {
		log.Fatal(err)
	}
}

func (h *tokoHandler) DetailTokoHandler(w http.ResponseWriter, r *http.Request){
	idString := r.URL.Query()["id"][0]
	id, _ := strconv.Atoi(idString)

	toko := h.tokoService.FindOneToko(id)

	path, _ := os.Getwd()
	t, err := template.ParseFiles(path+`\backend\views\layout.html`,path+`\backend\views\detail_toko.html`)
	if err != nil{
		log.Fatal(err)
	}

	t.Execute(w, toko)
}