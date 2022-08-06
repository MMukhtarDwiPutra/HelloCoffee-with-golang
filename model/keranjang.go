package model

type Keranjang struct{
	No int
	IdKeranjang int `json:"id_keranjang" db:"id_keranjang"`
	NamaMenu string `validate:"required" json:"nama_menu" db:"nama_menu"`
	Qty int `validate:"required" json:"qty" db:"qty"`
	IdUser int `validate:"required" json:"id_user" db:"id_user"`
	Harga int `validate:"required" json:"harga" db:"harga"`
	Total int
	IdMenu int `validate:"required" json:"id_menu" db:"id_menu"`
	Foto string `validate:"required" json:"foto_kopi" db:"foto_kopi"`
}