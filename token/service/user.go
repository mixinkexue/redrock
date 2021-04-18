package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"websocket/middleware"
	"websocket/modle"
)

func Register(ctx *gin.Context) (bool,string){
	var u modle.User
	ctx.ShouldBind(&u)
	res := modle.Register(u)
	if res {
		jwt:=middleware.NewJWT(u.Username)
		return true,jwt.Token
	}
	return false,""
}
func Login(ctx *gin.Context)(bool,string){
	fmt.Println("con")
	var u modle.User
	ctx.ShouldBind(&u)
	res := modle.Login(u)
	if res {
		jwt:=middleware.NewJWT(u.Username)
		return true,jwt.Token
	}
	return false,""
}
