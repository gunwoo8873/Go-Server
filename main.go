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

	dsn, err := sql.Open(
		"mysql", db_conn.Username+":"+db_conn.UserPassword+"@tcp("+db_conn.Hostname+":"+db_conn.Port+")/test",
	)
	log.Println(dsn.Stats())

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connect to mysql success")
	defer dsn.Close()
}

// Create user in the database
type Users struct {
	UserName     string `json:"user_name"`
	UserEmail    string `json:"user_email"`
	UserPassword string `json:"user_password"`
	UserRole     int    `json:"user_role"`
}

func createUser(ctx *gin.Context) {
	ctx.JSON(200, gin.H{})
}

func registerUser(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"msg": "Register User",
	})
}

func Route() *gin.Engine {
	route := gin.Default()
	route.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	userGroup := route.Group("/user")
	{
		userGroup.POST("/register", createUser)
		userGroup.POST("/signin", registerUser)
	}

	return route
}

func main() {
	mysql()

	route := Route()
	log.Println("Server start at url : localhost:50000")
	if err := route.Run(":50000"); err != nil {
		log.Fatal(err)
	}
}
