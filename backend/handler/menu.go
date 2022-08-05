package handler

import(
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/backend/menu"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/backend/komentar"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/model"
	"net/http"
	"log"
	"os"
	"github.com/gorilla/sessions"
	"html/template"
	"strconv"
)

type MenuHandler struct{
	menuService menu.Service
	komentarService komentar.Service
}

func NewMenuHandler (menuService menu.Service, komentarService komentar.Service) *MenuHandler{
	return &MenuHandler{menuService, komentarService}
}

func (h *MenuHandler) MenuHandler(w http.ResponseWriter, r *http.Request){
	store := sessions.NewCookieStore([]byte("super-secret"))
	session, err := store.Get(r, "session-name")
	
	id_user := session.Values["id_user"].(int)

	datastore := h.menuService.FindAllMenu()

	path, _ := os.Getwd()
	t, err := template.ParseFiles(path+`\backend\views\layout.html`,path+`\backend\views\menu.html`)
	if err != nil{
		log.Fatal(err)
	}

	tmp := map[string]interface{}{
    	"Id_user" : id_user,
    	"datastore" : datastore,
    }
	t.Execute(w, tmp)
}

func (h *MenuHandler) DetailMenu(w http.ResponseWriter, r *http.Request){
	store := sessions.NewCookieStore([]byte("super-secret"))
	session, err := store.Get(r, "session-name")

	id_menu_string := r.URL.Query()["id"][0]
	id_menu, _ := strconv.Atoi(id_menu_string)

	id_user := session.Values["id_user"].(int)

	menu := h.menuService.FindOneMenu(id_menu)

	komentar := h.komentarService.FindAllKomentar(id_menu, id_user)

	path, _ := os.Getwd()
	t, err := template.ParseFiles(path+`\backend\views\layout.html`,path+`\backend\views\detail_menu.html`)
	if err != nil{
		log.Fatal(err)
	}

	tmp := map[string]interface{}{
    	"Id_user" : id_user,
    	"Menu" : menu,
    	"komentar" : komentar,
    	"id_menu" : id_menu,
    }
	err = t.Execute(w, tmp)
	if err != nil{
		log.Fatal(err)
	}
}

func (h *MenuHandler) EditMenuHandler (w http.ResponseWriter, r *http.Request){
	idString := r.URL.Query()["id"][0]
	id_menu, _ := strconv.Atoi(idString)

	store := sessions.NewCookieStore([]byte("super-secret"))
	session, err := store.Get(r, "session-name")
	id_user := session.Values["id_user"].(int)

	m := h.menuService.FindOneMenu(id_menu)

	data := map[string]interface{}{
		"Id_user" : id_user,
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

func (h *MenuHandler) TambahMenuHandler (w http.ResponseWriter, r *http.Request){
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

func (h *MenuHandler) TambahMenuProcess (w http.ResponseWriter, r *http.Request){
	store := sessions.NewCookieStore([]byte("super-secret"))
	session, _ := store.Get(r, "session-name")
	id_toko := session.Values["id_toko"].(int)

	var menu model.Menu
	menu.Nama_menu = r.FormValue("nama_menu")
	harga, _ := strconv.Atoi(r.FormValue("harga"))
	menu.Harga = harga
	menu.Deskripsi = r.FormValue("deskripsi")
	menu.Jenis = r.FormValue("jenis")
	menu.Foto = ""
	menu.Id_toko = id_toko

	h.menuService.CreateMenu(menu)

	http.Redirect(w, r, "/home/toko", http.StatusSeeOther)
}

func (h *MenuHandler) DeleteMenu (w http.ResponseWriter, r *http.Request){
	idString := r.URL.Query()["id"][0]
	id, _ := strconv.Atoi(idString) 

	h.menuService.DeleteMenu(id)

	http.Redirect(w, r, "/home/toko", http.StatusSeeOther)
}

func (h *MenuHandler) EditMenuProcess(w http.ResponseWriter, r *http.Request){
	var menu model.Menu

	idString := r.URL.Query()["id"][0]
	id_menu, _ := strconv.Atoi(idString)
	menu.Id_menu = id_menu
	menu.Nama_menu = r.FormValue("nama_menu")
	hargaString := r.FormValue("harga")
	harga, _ := strconv.Atoi(hargaString)
	menu.Harga = harga
	menu.Deskripsi = r.FormValue("deskripsi")
	menu.Jenis = r.FormValue("jenis")

	h.menuService.UpdateMenu(menu)

	http.Redirect(w, r, "/home/toko", http.StatusSeeOther)
}