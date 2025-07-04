package config

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/api-skeleton/model"
	_ "github.com/go-sql-driver/mysql"
	migrate "github.com/rubenv/sql-migrate"
)

func Connect() (db *sql.DB) {
	var (
		config model.Config
	)

	file, err := os.Open("config/config.json")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.User, config.Password, config.Host, config.Port, config.DBName)

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}

func MigrateDB() {
	db := Connect()
	// Load migration files
	migrations := &migrate.FileMigrationSource{
		Dir: "migrations", // Directory where the migration files are located
	}

	// Apply the migrations
	n, err := migrate.Exec(db, "mysql", migrations, migrate.Up)
	if err != nil {
		log.Fatal("Failed to apply migrations:", err)
	}

	fmt.Printf("Applied %d migrations!\n", n)
}
