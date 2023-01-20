package database

import(
	"database/sql"

)


func New() *sql.DB{
	db, err := sql.Open("postgres",
	 "user=username password=password dbname=mydb sslmode=disable")
	 if err != nil {
		panic(err)
	 }
	 createTable(db)
	return db
}

func createTable(db *sql.DB) error{
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS orders (
		id SEREAL PRIMARY KEY,
		data jsonb
	)`)
	return err
}

func AddOrder (db *sql.DB) error{
	stmt,err := db.Prepare(`INSERT INTO orders `)
	if err != nil{
		return err
	}
	_, err = stmt.Exec()
	return err
}