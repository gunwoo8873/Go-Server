package server

import (
	"goserver/configs"

	//// Web Framwork
	"log"

	"github.com/gin-gonic/gin"
)

func Route() *gin.Engine {
	route := gin.Default()
	route.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"msg": "pong",
		})
	})

	return route
}

type LocalAddr struct {
	LocalHost string
	Port      string
}

func Run() {
	configs.ReadEnv()
	route := Route()

	localServer := LocalAddr{
		LocalHost: configs.GetEnv("PUBLIC_HOST"),
		Port:      configs.GetEnv("PORT"),
	}

	if err := route.Run(localServer.LocalHost + ":" + localServer.Port); err != nil {
		log.Fatal("Error not start server")
	}

	log.Println("Server start")
}
