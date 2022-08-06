package web

type CreateKeranjangRequest struct{
	Qty int `validate:"required" json:"qty" db:"qty"`
	IdMenu int `validate:"required" json:"id_menu" db:"id_menu"`
	IdUser int `validate:"required" json:"id_user" db:"id_user"`
}

type KeranjangResponse struct{
	IdKeranjang int `json:"id_keranjang" db:"id_keranjang"`
	NamaMenu string `json:"nama_menu" db:"nama_menu"`
	Qty int `json:"qty" db:"qty"`
	IdUser int `json:"id_user" db:"id_user"`
	Harga int `json:"harga" db:"harga"`
	Total int
	IdMenu int `json:"id_menu" db:"id_menu"`
	Foto string `json:"foto_kopi" db:"foto_kopi"`
}