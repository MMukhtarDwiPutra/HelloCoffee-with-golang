package menu

import(
	"database/sql"
	"log"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/model"
)

type Repository interface{
	FindAllMenu() []model.Menu
	FindOneMenu(id_menu int) model.Menu
	CreateMenu(nama string, harga int, deskripsi string, jenis string, id_toko int)
	DeleteMenu(id_menu int)
	UpdateMenu(nama string, harga int, deskripsi string, jenis string, id int)
}

type repository struct{
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository{
	return &repository{db}
}

func (r *repository) FindAllMenu() []model.Menu{
	var menus []model.Menu
	rows, err := r.db.Query(`SELECT m.id_menu, m.nama_menu, m.harga, m.deskripsi, m.jenis, m.foto_kopi, dt.nama_toko 
		FROM menu as m 
		JOIN daftar_toko as dt ON dt.id_toko = m.id_toko`)
	if err != nil{
		log.Fatal(err)
	}

	for rows.Next(){
		var menu model.Menu
		rows.Scan(&menu.Id_menu, &menu.Nama_menu, &menu.Harga, &menu.Deskripsi, &menu.Jenis, &menu.Foto, &menu.Nama_toko)

		menus = append(menus, menu)
	}

	return menus
}

func (r *repository) FindOneMenu(id_menu int) model.Menu{
	rows, err := r.db.Query(`SELECT m.id_menu, m.nama_menu, m.harga, m.deskripsi, m.jenis, m.foto_kopi, dt.nama_toko 
		FROM menu as m 
		JOIN daftar_toko as dt ON dt.id_toko = m.id_toko 
		WHERE m.id_menu = ?`,id_menu)
	if err != nil{
		log.Fatal(err)
	}

	var menu model.Menu

	for rows.Next(){
		rows.Scan(&menu.Id_menu, &menu.Nama_menu, &menu.Harga, &menu.Deskripsi, &menu.Jenis, &menu.Foto, &menu.Nama_toko)
	}

	return menu
}

func (r *repository) CreateMenu(nama string, harga int, deskripsi string, jenis string, id_toko int){
	r.db.Query("INSERT INTO menu (nama_menu, harga, deskripsi, jenis, id_toko) VALUES (?, ?, ?, ?, ?)",nama, harga, deskripsi, jenis, id_toko)
}

func (r *repository) DeleteMenu(id_menu int) {
	r.db.Query("DELETE FROM menu WHERE id_menu = ?",id_menu)
}

func (r *repository) UpdateMenu(nama string, harga int, deskripsi string, jenis string, id_menu int){
	r.db.Query("UPDATE menu SET nama_menu = ?, harga = ?, deskripsi = ?, jenis = ? WHERE id_menu = ?",nama, harga, deskripsi, jenis, id_menu)
}