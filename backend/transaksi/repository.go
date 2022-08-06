package transaksi

import(
	"database/sql"
	"log"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/model"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/model/web"
)

type Repository interface{
	GetAllTransaksi(id_user int) []model.Transaksi
	GetTransaksiToko(id_toko int) []model.Transaksi
	UpdateStatusTransaksi(status string, id_transaksi int)
	AddTransaksi(fullname string, email string, address string, city string, zip string, state string)
	GetTransaksiAPI(id_transaksi int) web.TransaksiResponse
	GetAllTransaksiAPI() []web.TransaksiResponse
}

type repository struct{
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository{
	return &repository{db}
}

func (r *repository) GetAllTransaksi(id_user int) []model.Transaksi{
	var transaksi []model.Transaksi
	rows, err := r.db.Query("SELECT t.tanggal_transaksi, t.qty, m.harga, m.nama_menu, t.status_transaksi FROM transaksi t JOIN menu m ON t.id_menu = m.id_menu WHERE t.id_user = ?",id_user)
	if err != nil{
		log.Fatal(err)
	}

	for rows.Next(){
		var t model.Transaksi
		var harga int
		rows.Scan(&t.TanggalTransaksi, &t.Qty, &harga, &t.NamaMenu, &t.StatusTransaksi)
		t.Harga = harga * t.Qty

		transaksi = append(transaksi, t)
	}

	return transaksi
}

func (r *repository) GetTransaksiToko(id_toko int) []model.Transaksi{
	rows, err := r.db.Query(`SELECT t.id_transaksi, t.tanggal_transaksi, dt.full_name, dt.email, dt.address, dt.city, dt.zip, dt.state, t.Qty, m.Harga, m.foto_kopi, m.nama_menu FROM detail_transaksi dt JOIN transaksi t ON t.id_detail_transaksi = dt.id_detail_transaksi JOIN menu m ON m.id_menu = t.id_menu WHERE id_toko = ? and t.status_transaksi = "Baru"`,id_toko)
	if err != nil{
		log.Fatal(err)
	}

	var data []model.Transaksi
	for rows.Next(){
		var t model.Transaksi
		rows.Scan(&t.IdTransaksi, &t.TanggalTransaksi, &t.Nama, &t.Email, &t.Address, &t.City, &t.Zip, &t.State, &t.Qty, &t.Harga, &t.Foto, &t.NamaMenu)
		t.Harga = t.Qty * t.Harga
		data = append(data, t)
	}
	return data
}

func (r *repository) UpdateStatusTransaksi(status string, id_transaksi int){
	r.db.Query("UPDATE transaksi SET status_transaksi = ? WHERE id_transaksi = ?",status,id_transaksi)
}

func (r *repository) AddTransaksi(fullname string, email string, address string, city string, zip string, state string){
	r.db.Query("INSERT INTO detail_transaksi (full_name, email, address, city, zip, state) VALUES (?, ?, ?, ?, ?, ?) ",fullname, email, address, city, zip, state)
}

func (r *repository) GetTransaksiAPI(id_transaksi int) web.TransaksiResponse{
	var dt web.DetailTransaksiResponse
	var t web.TransaksiResponse

	rows, err := r.db.Query("SELECT t.id_transaksi, t.tanggal_transaksi, t.qty, m.nama_menu, m.harga, t.status_transaksi, dt.full_name, dt.email, dt.address, dt.state, dt.zip, dt.city FROM transaksi t JOIN menu m ON t.id_menu = m.id_menu JOIN detail_transaksi dt ON dt.id_detail_transaksi = t.id_detail_transaksi WHERE t.id_transaksi = ?", id_transaksi)
	if err != nil{
		log.Fatal(err)
	}

	for rows.Next(){
		rows.Scan(&t.IdTransaksi, &t.TanggalTransaksi, &t.Qty, &t.NamaMenu, &t.Harga, &t.StatusTransaksi, &dt.Nama, &dt.Email, &dt.Address, &dt.State, &dt.Zip, &dt.City)
		t.DetailTransaksi = dt
	}

	return t
}

func (r *repository) GetAllTransaksiAPI() []web.TransaksiResponse{
	var transaksi []web.TransaksiResponse

	rows, err := r.db.Query("SELECT t.id_transaksi, t.tanggal_transaksi, t.qty, m.nama_menu, m.harga, t.status_transaksi, dt.full_name, dt.email, dt.address, dt.state, dt.zip, dt.city FROM transaksi t JOIN menu m ON t.id_menu = m.id_menu JOIN detail_transaksi dt ON dt.id_detail_transaksi = t.id_detail_transaksi")
	if err != nil{
		log.Fatal(err)
	}

	for rows.Next(){
		var dt web.DetailTransaksiResponse
		var t web.TransaksiResponse
		rows.Scan(&t.IdTransaksi, &t.TanggalTransaksi, &t.Qty, &t.NamaMenu, &t.Harga, &t.StatusTransaksi, &dt.Nama, &dt.Email, &dt.Address, &dt.State, &dt.Zip, &dt.City)
		t.DetailTransaksi = dt
		transaksi = append(transaksi, t)
	}

	return transaksi
}