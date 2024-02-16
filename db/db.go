package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func init() {
	RunMigrates()
}

func Connection() *sql.DB {
	// String conection
	connStr := "user=myuser dbname=mydatabase password=mypassword sslmode=disable"

	// Open conection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Connecting error with db:", err)
		return nil
	}

	// Verify if conection is enable
	err = db.Ping()
	if err != nil {
		fmt.Println("Connection error test with db:", err)
		return nil
	}

	fmt.Println("Connection with Postgres was successfull :D")

	return db
}
