package cmd

import (
	"github.com/gin-gonic/gin"
	"websocket/controller"
	"websocket/middleware"
)

func Entrance(){
	r := gin.Default()
	r.POST("/register",controller.Register)
	r.POST("/login",controller.Login)
	r.GET("/ws:room",middleware.Cookie(),controller.Ws)
	r.Run("127.0.0.1:8080")

}
