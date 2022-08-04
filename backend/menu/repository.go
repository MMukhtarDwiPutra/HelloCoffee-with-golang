package menu

import(
	"database/sql"
	"log"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/model"
)

type Repository interface{
	FindAllMenu() []model.Menu
	FindOneMenu(id_menu int) model.Menu
	CreateMenu(nama string, harga int, deskripsi string, jenis string, id_toko int, foto_kopi string)
	DeleteMenu(id_menu int)
	UpdateMenu(nama string, harga int, deskripsi string, jenis string, id int)
	FindAllMenuFromToko(id_toko int) []model.Menu
	GetLastMenu() model.Menu
}

type repository struct{
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository{
	return &repository{db}
}

func (r *repository) FindAllMenuFromToko(id_toko int) []model.Menu{
	rows, err := r.db.Query("SELECT id_menu, nama_menu, harga, jenis, foto_kopi, id_toko FROM menu WHERE id_toko = ?",id_toko)
	if err != nil{
		log.Fatal(err)
	}

	var data []model.Menu
	for rows.Next(){
		var m model.Menu
		rows.Scan(&m.Id_menu, &m.Nama_menu, &m.Harga, &m.Jenis, &m.Foto, &m.Id_toko)

		data = append(data, m)
	}

	return data
}

func (r *repository) FindAllMenu() []model.Menu{
	var menus []model.Menu
	rows, err := r.db.Query(`SELECT m.id_menu, m.nama_menu, m.harga, m.deskripsi, m.jenis, m.foto_kopi, dt.nama_toko, m.id_toko 
		FROM menu as m 
		JOIN daftar_toko as dt ON dt.id_toko = m.id_toko`)
	if err != nil{
		log.Fatal(err)
	}

	for rows.Next(){
		var menu model.Menu
		rows.Scan(&menu.Id_menu, &menu.Nama_menu, &menu.Harga, &menu.Deskripsi, &menu.Jenis, &menu.Foto, &menu.Nama_toko, &menu.Id_toko)

		menus = append(menus, menu)
	}

	return menus
}

func (r *repository) FindOneMenu(id_menu int) model.Menu{
	rows, err := r.db.Query(`SELECT m.id_menu, m.nama_menu, m.harga, m.deskripsi, m.jenis, m.foto_kopi, dt.nama_toko, m.id_toko 
		FROM menu as m 
		JOIN daftar_toko as dt ON dt.id_toko = m.id_toko 
		WHERE m.id_menu = ?`,id_menu)
	if err != nil{
		log.Fatal(err)
	}

	var menu model.Menu

	for rows.Next(){
		rows.Scan(&menu.Id_menu, &menu.Nama_menu, &menu.Harga, &menu.Deskripsi, &menu.Jenis, &menu.Foto, &menu.Nama_toko, &menu.Id_toko)
	}

	return menu
}

func (r *repository) CreateMenu(nama string, harga int, deskripsi string, jenis string, id_toko int, foto_kopi string){
	r.db.Query("INSERT INTO menu (nama_menu, harga, deskripsi, jenis, id_toko, foto_kopi) VALUES (?, ?, ?, ?, ?, ?)",nama, harga, deskripsi, jenis, id_toko, foto_kopi)
}

func (r *repository) DeleteMenu(id_menu int) {
	r.db.Query("DELETE FROM menu WHERE id_menu = ?",id_menu)
}

func (r *repository) UpdateMenu(nama string, harga int, deskripsi string, jenis string, id_menu int){
	r.db.Query("UPDATE menu SET nama_menu = ?, harga = ?, deskripsi = ?, jenis = ? WHERE id_menu = ?",nama, harga, deskripsi, jenis, id_menu)
}

func (r *repository) GetLastMenu() model.Menu{
	var menu model.Menu

	rows, err := r.db.Query("SELECT id_menu, nama_menu, harga, deskripsi, jenis, foto_kopi, id_toko FROM menu ORDER BY id_menu DESC LIMIT 1")
	if err != nil{
		log.Fatal(err)
	}
	for rows.Next(){
		rows.Scan(&menu.Id_menu, &menu.Nama_menu, &menu.Harga, &menu.Deskripsi, &menu.Jenis, &menu.Foto, &menu.Id_toko)
	}

	return menu
}