package mysql

import (
	"database/sql"
	"fmt"
	"log"

	"goserver/configs"

	//// MySQL Database Driver
	_ "github.com/go-sql-driver/mysql"
)

type DBConfig struct {
	Hostname string
	Username string
	Password string
	Port     string
	Schema   string
}

// Public Database connection
func Connect() {
	// Read to string env data
	configs.ReadEnv()

	DBAddr := DBConfig{
		Hostname: configs.GetEnv("MYSQL_HOSTNAME"),
		Username: configs.GetEnv("MYSQL_USERNAME"),
		Password: configs.GetEnv("MYSQL_PASSWORD"),
		Port:     configs.GetEnv("MYSQL_PORT"),
		Schema:   configs.GetEnv("MYSQL_SCHEMA"),
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		DBAddr.Username, DBAddr.Password, DBAddr.Hostname, DBAddr.Port, DBAddr.Schema,
	)

	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to Open database: %d", err)
	}

	if err := conn.Ping(); err != nil {
		log.Fatalf("Failed to Ping database: %d", err)
	}

	log.Println("Successfully connected to MySQL database")
	defer conn.Close()
}
