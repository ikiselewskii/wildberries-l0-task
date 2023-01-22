package database

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"
	"wildberries-l0-task/cache/cache"
	"wildberries-l0-task/cache/models"

	_ "github.com/lib/pq"
)

var connection *sql.DB

func Connect() (*sql.DB, error){
	var err error
	if connection == nil{
		log.Print(os.Getenv("DATABASE_URL"))
		connection, _ = sql.Open("postgres",
		 "user=postgres password=wildberries dbname=wildberriesdb sslmode=disable")
		err = connection.Ping()
		if err != nil {
			log.Fatalf("ERROR failed to connect to a database %s", err)
			return nil, err
		}
		err := createTable()
		if err != nil{
			log.Fatalf("ERROR can`t create tables due to %s", err)
		}
	}
	return connection, nil
}

func createTable() error{
	_, err := connection.Exec(
		`CREATE TABLE IF NOT EXISTS orders (
		id SEREAL PRIMARY KEY,
		uid VARCHAR,
		data JSONB
	)`)
	return err
}

func AddOrder (order *models.Order) error{
	stmt,err := connection.Prepare(`INSERT (uid, data) INTO orders`)
	if err != nil{
		return err
	}
	_, err = stmt.Exec()
	return err
}

func GetAllOrders () (error) {
	rows, err := connection.Query("SELECT (uid, data) FROM orders")
	if err != nil{
		log.Printf("ERROR Unable To Select %s", err)
		return err
	}
	for rows.Next() {
		var order models.Order
		var uid string
		var data []byte
		err := rows.Scan(&uid,&data)
		if err != nil{
			log.Printf("ERROR Scanning Went Wrong: %s", err)
			return err
		}
		err = json.Unmarshal(data, &order)
		if err != nil{
			log.Printf("ERROR Unmarshalling Went Wrong: %s", err)
			return err
		}
		cache.Add(&order)
		
	}
	return nil
}