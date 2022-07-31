package model

type Keranjang struct{
	No int
	IdKeranjang int `db:"id_keranjang"`
	NamaMenu string `db:"nama_menu"`
	Qty int `db:"qty"`
	IdUser int `db:"id_user"`
	Harga int `db:"harga"`
	Total int
	IdMenu int `db:"id_menu"`
}