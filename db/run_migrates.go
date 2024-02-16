package db

import (
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func RunMigrates() {
	driver, err := postgres.WithInstance(Connection(), &postgres.Config{})
	if err != nil {
		fmt.Println("Connection error:", err)
	}

	mainPath, err := os.Getwd()
	if err != nil {
		fmt.Println("os.Getwd error:", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+mainPath+"/db/migrations",
		"postgres", driver)
	if err != nil {
		fmt.Println("Instance of Migration has been failed", err)
		os.Exit(1)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		fmt.Println("Error on migrations:", err)
		os.Exit(1)
	}

	fmt.Println("Migrates applied with success!")
}
