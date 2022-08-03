package keranjang

import(
	"database/sql"
	"log"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/model"
	"time"
)

type Repository interface{
	GetAllKeranjang(id_user int) ([]model.Keranjang, int)
	AddKeranjang(id_menu int, qty int, id_user int)
	DeleteAllKeranjang(id_user int)
	GetLastRowIdDetailTransaksi() int
	AddToTransaksi(tanggal string, k model.Keranjang, id_detail_transaksi int)
	DeleteKeranjang(k model.Keranjang)
	AddTransaksi(id_user int)
}

type repository struct{
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository{
	return &repository{db}
}

func (r *repository) GetAllKeranjang(id_user int) ([]model.Keranjang, int){
	rows, err := r.db.Query("SELECT k.id_keranjang, k.qty, k.id_user, m.harga, m.nama_menu, m.foto_kopi FROM keranjang k JOIN menu m ON k.id_menu = m.id_menu WHERE k.id_user = ?",id_user)
	if err != nil{
		log.Fatal(err)
	}

	var data []model.Keranjang
	i := 1
	totalAll := 0
	for rows.Next(){
		var k model.Keranjang
		rows.Scan(&k.IdKeranjang, &k.Qty, &k.IdUser, &k.Harga, &k.NamaMenu, &k.Foto)
		k.Total = k.Qty * k.Harga
		k.No = i
		data = append(data, k)
		totalAll = totalAll + k.Total
		i++
	}
	return data, totalAll
}

func (r *repository) AddKeranjang(id_menu int, qty int, id_user int){
	r.db.Query("INSERT INTO keranjang (qty, id_menu, id_user) VALUES (?, ?, ?)",qty, id_menu, id_user)
}

func (r *repository) DeleteAllKeranjang(id_user int){
	r.db.Query("DELETE FROM keranjang WHERE id_user = ?",id_user)
}

func (r *repository) GetLastRowIdDetailTransaksi() int{
	row, _ := r.db.Query("SELECT id_detail_transaksi FROM detail_transaksi ORDER BY id_detail_transaksi DESC LIMIT 1")

	var id_detail_transaksi int
	for row.Next(){
		row.Scan(&id_detail_transaksi)
	}
	return id_detail_transaksi
}

func (r *repository) AddToTransaksi(tanggal string, k model.Keranjang, id_detail_transaksi int){
	r.db.Query(`INSERT INTO transaksi (tanggal_transaksi, qty, id_user, id_menu, id_detail_transaksi, status_transaksi) VALUES (?, ?, ?, ?, ?, "Baru")`, tanggal, k.Qty, k.IdUser, k.IdMenu, id_detail_transaksi)
}

func (r *repository) DeleteKeranjang(k model.Keranjang){
	r.db.Query("DELETE FROM keranjang WHERE id_keranjang = ?", k.IdKeranjang)
}

func (r *repository) AddTransaksi(id_user int){
	currentTime := time.Now()
	tanggal := currentTime.Format("2006-01-02")

	id_detail_transaksi := r.GetLastRowIdDetailTransaksi()
	rows, _ := r.db.Query("SELECT id_keranjang, qty, id_user, id_menu FROM keranjang WHERE id_user = ?",id_user)
	for rows.Next(){
		var k model.Keranjang
		rows.Scan(&k.IdKeranjang, &k.Qty, &k.IdUser, &k.IdMenu)
		r.AddToTransaksi(tanggal, k, id_detail_transaksi)
		r.DeleteKeranjang(k)
	}
}