package controller

import (
	"github.com/gin-gonic/gin"
	"websocket/service"
)

func Register(ctx *gin.Context)  {
	res := service.Register(ctx)
	if !res {
		ctx.JSON(200, gin.H{
			"message": "register failed，username or phone has benn used",
		})
	} else {
		ctx.JSON(200, gin.H{
			"message": "register successfully",
		})
	}
}
func Login(ctx *gin.Context) {
	res := service.Login(ctx)
	if !res {
		ctx.JSON(200, gin.H{
			"message":"login failed",
		})
	} else {
		ctx.JSON(200, gin.H{
			"message": "register successfully",
		})
	}}
func Ws(ctx *gin.Context){
	service.Ws(ctx)

}