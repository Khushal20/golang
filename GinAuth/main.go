package main

import (
	"khushal/GinAuth/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	public:=router.Group("/api")

	public.POST("/login", handler.Login)
	public.GET("/validate", handler.Validate)

	protectedUser:= router.Group("/api/")
	protectedUser.Use(handler.Validate)
	protectedUser.GET("user", handler.Userget)

	router.Run(":8080")
}
