package main

import (
	"database/sql"
	"log"
	"os"

	//// Read .env file
	"github.com/joho/godotenv"

	//// Web Framwork
	"github.com/gin-gonic/gin"

	//// MySQL Database Driver
	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	Username     string
	UserPassword string
	Hostname     string
	Port         string
}

func mysql() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db_conn := DB{
		Username:     os.Getenv("MYSQL_USERNAME"),
		UserPassword: os.Getenv("MYSQL_PASSWORD"),
		Hostname:     os.Getenv("MYSQL_HOSTNAME"),
		Port:         os.Getenv("MYSQL_PORT"),
	}

	db, err := sql.Open(
		"mysql", db_conn.Username+":"+db_conn.UserPassword+"@tcp("+db_conn.Hostname+":"+db_conn.Port+")/test",
	)
	if err != nil {
		log.Fatalf("connect to mysql failed : %v", err)
	}

	defer db.Close()

	log.Println("connect to mysql success")
}

func server() {
	route := gin.Default()
	route.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	log.Println("Server start at url : localhost:50000")
	route.Run(":50000")
}

func main() {
	mysql()
	server()
}
