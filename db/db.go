package db

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	RunMigrates()
}

func connection() *sql.DB {
	gormDb := GormConnection()

	sqlDb, _ := gormDb.DB()
	return sqlDb
}

func GormConnection() *gorm.DB {
	// String conection
	dsn := "host=localhost user=myuser dbname=mydatabase password=mypassword sslmode=disable"

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

	return db
}
