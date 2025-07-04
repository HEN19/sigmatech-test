package config

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/api-skeleton/model"
	_ "github.com/jackc/pgx/v5/stdlib"
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

	// Read the configuration file
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	// Parse the JSON configuration
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)
	}

	// Build the DSN for PostgreSQL
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.User, config.Password, config.Host, config.Port, config.DBName)

	// Open the database connection
	db, err = sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}

	fmt.Println("Successfully connected to PostgreSQL!")
	return db
}

func MigrateDB() {
	db := Connect()
	// Load migration files
	migrations := &migrate.FileMigrationSource{
		Dir: "migrations", // Directory where the migration files are located
	}

	// Apply the migrations
	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatal("Failed to apply migrations:", err)
	}

	fmt.Printf("Applied %d migrations!\n", n)
}
