package main

import (
	"khushal/GinAuth/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/login", handler.Login)
	router.GET("/validate", handler.Validate)
	router.Run(":8080")
}
