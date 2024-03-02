package db

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func connection() *sql.DB {
	sqlDb, _ := DB.DB()
	return sqlDb
}

func InitDB(dsn string) *gorm.DB {
	// dsn := "host=localhost user=myuser dbname=mydatabase password=mypassword sslmode=disable"

	// Open conection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Connecting error with db:", err)
		return nil
	}

	// Verify if conection is enable
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("sqlDB error", err)
	}

	err = sqlDB.Ping()

	if err != nil {
		fmt.Println("Connection error test with db:", err)
		return nil
	}

	fmt.Println("Connection with Postgres was successfull :D")
	DB = db

	return db
}

func MakeDSN(host, user, dbname, password string) string {
	return fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable", host, user, dbname, password)
}
