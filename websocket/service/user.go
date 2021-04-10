package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"websocket/modle"
)

func Register(ctx *gin.Context) bool{
	var u modle.User
	ctx.ShouldBind(&u)
	res := modle.Register(u)
	if res {

		//记录登录状态
		cookie := &http.Cookie{
			Name:     "Id",
			Value:    u.Username,
			Path:     "/",
			HttpOnly: false,
			MaxAge:   2000,
		}
		http.SetCookie(ctx.Writer, cookie)
		return true
	}
	return false
}
func Login(ctx *gin.Context)bool{
	fmt.Println("con")
	var u modle.User
	ctx.ShouldBind(&u)
	res := modle.Login(u)
	if res {

		//记录登录状态
		cookie := &http.Cookie{
			Name:     "username",
			Value:    u.Username,
			Path:     "/",
			HttpOnly: false,
			MaxAge:   2000,
		}
		http.SetCookie(ctx.Writer, cookie)
		return true
	}
	return false
}
