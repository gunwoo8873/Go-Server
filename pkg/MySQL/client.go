package MySQL

import (
	"database/sql"
	"log"
	"os"

	//// Read .env file
	"github.com/joho/godotenv"

	//// MySQL Database Driver
	_ "github.com/go-sql-driver/mysql"
)

func EnvLoad() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

type DBConfig struct {
	Hostname string
	Username string
	Password string
	Port     string
	Database string
}

func ClientConn() {
	EnvLoad()
	DBAddr := DBConfig{
		Hostname: os.Getenv("MYSQL_HOSTNAME"),
		Username: os.Getenv("MYSQL_USERNAME"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		Port:     os.Getenv("MYSQL_PORT"),
		Database: os.Getenv("MYSQL_DATABASE"),
	}

	dsn, err := sql.Open(
		"mysql", DBAddr.Username+":"+DBAddr.Password+"@tcp("+DBAddr.Hostname+":"+DBAddr.Port+")/"+DBAddr.Database,
	)
	if err != nil {
		log.Fatal(err)
	}

	defer dsn.Close()
}
