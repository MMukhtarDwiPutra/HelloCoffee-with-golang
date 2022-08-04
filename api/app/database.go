package app

import(
	"database/sql"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/api/helper"
)

func NewDB() *sql.DB{
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/hellocoffee")
	helper.PanicIfError(err)

	return db
}