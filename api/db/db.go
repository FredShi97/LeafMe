package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

// Open : Opens DB conection
func Open() *sql.DB {
	username := "root"
	password := "MssMt87Tcnka4tN"
	protocol := "tcp"
	address := "mysql"
	port := "3306"
	dbname := "leafme"
	db, _ := sql.Open("mysql",
		username+":"+password+"@"+protocol+"("+address+":"+port+")/"+dbname+"?parseTime=true")
	return db
}

// Close : Close DB connection
func Close(db *sql.DB) {
	db.Close()
}