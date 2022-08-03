package handler

import(
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/backend/keranjang"
	"net/http"
	"github.com/gorilla/sessions"
	"os"
	"html/template"
	"log"
	"strconv"
)

type keranjangHandler struct{
	keranjangService keranjang.Service
}

func NewKeranjangHandler(keranjangService keranjang.Service) *keranjangHandler{
	return &keranjangHandler{keranjangService}
}

func (h *keranjangHandler) KeranjangHandler(w http.ResponseWriter, r *http.Request){
	id := r.URL.Query()["id"][0]
	id_user, _ := strconv.Atoi(id)

	data, _ := h.keranjangService.GetAllKeranjang(id_user)

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

func (h *keranjangHandler) TambahKeranjang(w http.ResponseWriter, r *http.Request){
	id_menu_string := r.URL.Query()["id"][0]
	id_menu, _ := strconv.Atoi(id_menu_string)
	store := sessions.NewCookieStore([]byte("super-secret"))
	session, err := store.Get(r, "session-name")
	if err != nil{
		log.Fatal(err)
	}

	id_user := session.Values["id_user"].(int)
	qtyString := r.FormValue("jumlah")
	qty, _ := strconv.Atoi(qtyString)

	h.keranjangService.AddKeranjang(id_menu, qty, id_user)

	http.Redirect(w, r, "/menu", http.StatusSeeOther)
}

func (h *keranjangHandler) HapusKeranjang(w http.ResponseWriter, r *http.Request){
	id := r.URL.Query()["id"][0]
	id_user, _ := strconv.Atoi(id)
	
	h.keranjangService.DeleteKeranjang(id_user)

	http.Redirect(w,r,"/menu", http.StatusSeeOther)
}

func (h *keranjangHandler) CheckoutHandler (w http.ResponseWriter, r *http.Request){
	path, _ := os.Getwd()
	t, err := template.ParseFiles(path+`\backend\views\layout.html`,path+`\backend\views\checkout.html`)
	if err != nil{
		log.Fatal(err)
	}

	id := r.URL.Query()["id"][0]
	id_user, _ := strconv.Atoi(id)

	data, totalAll := h.keranjangService.GetAllKeranjang(id_user)

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

func (h *keranjangHandler) CheckoutProcess (w http.ResponseWriter, r *http.Request){
	id := r.URL.Query()["id"][0]
	id_user, _ := strconv.Atoi(id)
	fullName := r.FormValue("fname")
	email := r.FormValue("email")
	address := r.FormValue("address")
	city := r.FormValue("city")
	zip := r.FormValue("zip")
	state := r.FormValue("state")

	h.keranjangService.CheckoutProcess(id_user, fullName, email, address, city, zip, state)

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func (h *keranjangHandler) CheckoutNowHandler(w http.ResponseWriter, r *http.Request){
	id_menu_string := r.URL.Query()["id"][0]
	id_menu, _ := strconv.Atoi(id_menu_string)
	store := sessions.NewCookieStore([]byte("super-secret"))
	session, err := store.Get(r, "session-name")
	if err != nil{
		log.Fatal(err)
	}

	id_user := session.Values["id_user"].(int)
	qtyString := r.FormValue("jumlah")
	qty, _ := strconv.Atoi(qtyString)

	h.keranjangService.AddKeranjang(id_menu, qty, id_user)

	http.Redirect(w, r, "/keranjang/checkout/?id="+strconv.Itoa(id_user), http.StatusSeeOther)
}