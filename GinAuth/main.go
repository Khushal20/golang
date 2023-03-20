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
	public.POST("/user", handler.UserPost)

	protectedUser:= router.Group("/api/")
	protectedUser.Use(handler.Validate)
	protectedUser.GET("user", handler.Userget)

	router.Run(":8080")
}
