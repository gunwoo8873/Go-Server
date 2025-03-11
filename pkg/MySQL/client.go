package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	//// Read .env file
	"github.com/joho/godotenv"

	//// MySQL Database Driver
	_ "github.com/go-sql-driver/mysql"
)

type DBConfig struct {
	Hostname string
	Username string
	Password string
	Port     string
	Database string
}

func ReadEnv() {
	workPath, err := os.Getwd() // Current work get directory path
	if err != nil {
		log.Fatal(err)
	}

	envPath := filepath.Join(workPath, ".env")
	if err := godotenv.Load(envPath); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func Connect() {
	ReadEnv()
	DBAddr := DBConfig{
		Hostname: os.Getenv("MYSQL_HOSTNAME"),
		Username: os.Getenv("MYSQL_USERNAME"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		Port:     os.Getenv("MYSQL_PORT"),
		Database: os.Getenv("MYSQL_DATABASE"),
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		DBAddr.Username, DBAddr.Password, DBAddr.Hostname, DBAddr.Port, DBAddr.Database,
	)

	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Errorf("Failed to Open database: %w", err)
	}

	if err := conn.Ping(); err != nil {
		fmt.Errorf("Failed to Ping database: %w", err)
	}

	log.Println("Successfully connected to MySQL database")
}
