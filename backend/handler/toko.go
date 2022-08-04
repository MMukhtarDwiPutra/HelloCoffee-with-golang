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

type TokoHandler struct{
	tokoService toko.Service
}

func NewTokoHandler(tokoService toko.Service) *TokoHandler{
	return &TokoHandler{tokoService}
}

func (h *TokoHandler) HomeTokoHandler (w http.ResponseWriter, r *http.Request){
	store := sessions.NewCookieStore([]byte("super-secret"))
	session, err := store.Get(r, "session-name")
	id_user := session.Values["id_user"].(int)
	id_toko := session.Values["id_toko"].(int)

	data := h.tokoService.FindAllMenuFromToko(id_toko)

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



func (h *TokoHandler) HomeHandler(w http.ResponseWriter, r *http.Request){
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

func (h *TokoHandler) DetailTokoHandler(w http.ResponseWriter, r *http.Request){
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