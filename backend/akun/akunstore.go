package akun

type AkunStore struct{
	nama string
	password string
}

func NewDataAkun() *AkunStore{
	namaDefault := ""
	passwordDefault := ""

	return &AkunStore{
		nama := namaDefault,
		password := passwordDefault,
	}
}