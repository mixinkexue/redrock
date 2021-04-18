package controller

import (
	"github.com/gin-gonic/gin"
	"websocket/service"
)

func Register(ctx *gin.Context)  {
	res,token:= service.Register(ctx)
	if !res {
		ctx.JSON(200, gin.H{
			"message": "register failed，username or phone has benn used",
		})
	} else {
		ctx.JSON(200, gin.H{
			"message": "register successfully",
			"token":token,
		})
	}
}
func Login(ctx *gin.Context) {
	res,token:= service.Login(ctx)
	if !res {
		ctx.JSON(200, gin.H{
			"message":"login failed",
		})
	} else {
		ctx.JSON(200, gin.H{
			"message": "register successfully",
			"token":token,
		})
	}}
func Ws(ctx *gin.Context){
	service.Ws(ctx)

}