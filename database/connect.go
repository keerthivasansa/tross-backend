package database

import "database/sql"

var DB *Queries
var SqlDb *sql.DB

func Connect() {
	conn := "mysql:76b2c605b9631832f1f8@tcp(144.24.135.144:6432)/tross?parseTime=true"
	db, err := sql.Open("mysql", conn)

	if err != nil {
		panic(err)
	}
	DB = New(db)
	SqlDb = db
}
