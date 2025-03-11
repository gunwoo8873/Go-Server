package server

import (
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
	Port string
}

func Run() {
	route := Route()

	portAddr := LocalAddr{
		Port: "50000",
	}

	if err := route.Run(":" + portAddr.Port); err != nil {
		log.Fatal("Error not start server")
	}
}
