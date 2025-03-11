package server

import "github.com/gin-gonic/gin"

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

func Request() *gin.Engine {
	route := gin.Default()
	route.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	userGroup := route.Group("/user")
	{
		userGroup.POST("/register", createUser)
		userGroup.POST("/signin/*", registerUser)
	}

	return route
}
