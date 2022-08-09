package app

import (
	"database/sql"
	"my-blog/helper"
	"time"
)

func NewDB() *sql.DB {
	//change <Password> with your MySql password
	db, err := sql.Open("mysql", "root:<Password>@tcp(localhost:3306)/myblog")
	helper.PanicIfError(err)


	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
