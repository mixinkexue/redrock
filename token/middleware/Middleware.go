package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Cookie() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := c.Request.Cookie("username")
		if err == nil {
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "没有权限",
			})
			c.Abort()
		}
	}
}

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}
type Payload struct {
	Iss string `json:"iss"`
	Exp string `json:"exp"`
	Iat string `json:"iat"`
	Id  string   `json:"id"`
}
type JWT struct {
	Header    Header
	Payload   Payload
	Signature string
	Token     string
}
func NewJWT(id string) JWT {
	var jwt JWT
	jwt.Header = Header{
		Alg: "HS256",
		Typ: "JWT",
	}
	jwt.Payload = Payload{
		Iss: "dj",
		Exp: strconv.FormatInt(time.Now().Add(3*time.Hour).Unix(), 10),
		Iat: strconv.FormatInt(time.Now().Unix(), 10),
		Id:  id,
	}
	h, _ := json.Marshal(jwt.Header)
	p, _ := json.Marshal(jwt.Payload)
	baseh := base64.StdEncoding.EncodeToString(h)
	basep := base64.StdEncoding.EncodeToString(p)
	secret := baseh + "." + basep
	key := "dj"
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(secret))
	s := mac.Sum(nil)
	jwt.Signature = base64.StdEncoding.EncodeToString(s)
	jwt.Token = secret + "." + jwt.Signature
	return jwt
}

func Check(token string) (jwt JWT, err error) {
	err = errors.New("token error")
	arr := strings.Split(token, ".")
	if len(arr) < 3 {
		fmt.Println("59------", err)
		return
	}
	baseh := arr[0]
	h, err := base64.StdEncoding.DecodeString(baseh)
	if err != nil {
		fmt.Println("decode header", err)
		return
	}
	err = json.Unmarshal(h, &jwt.Header)
	if err != nil {
		fmt.Println("unmarshal header", err)
		return
	}
	basep := arr[1]
	p, err := base64.StdEncoding.DecodeString(basep)
	if err != nil {
		fmt.Println("decode payload", err)
		return
	}
	err = json.Unmarshal(p, &jwt.Payload)
	if err != nil {
		fmt.Println("unmarshal payload", err)
		return
	}
	bases := arr[2]
	s1, err := base64.StdEncoding.DecodeString(bases)
	if err != nil {
		fmt.Println("decode bases", err)
		return
	}
	secret:= baseh + "." + basep
	w := []byte("dj")
	mac := hmac.New(sha256.New, w)
	mac.Write([]byte(secret))
	s2 := mac.Sum(nil)
	if string(s1) != string(s2) {
		return
	} else {
		jwt.Signature = arr[2]
		jwt.Token = token
	}
	return jwt, nil
}
func CheckUser(ctx *gin.Context)  {
	tokenHeader := ctx.GetHeader("Authorization")
	if tokenHeader == "" || !strings.HasPrefix(tokenHeader, "Bearer ") {
		ctx.JSON(403,gin.H{
			"err":"错误http头",
		})
		ctx.Abort()
		return
	}
	token := strings.TrimPrefix(tokenHeader, "Bearer ")
	jwt, err := Check(token)
	if err != nil {
		ctx.JSON(403,gin.H{
			"err":"错误信息",
		})
		ctx.Abort()
		return
	}
	exp,_:=strconv.Atoi(jwt.Payload.Exp)
	if exp>int(time.Now().Unix()){
		ctx.Set("user", jwt.Payload.Id)
		ctx.Next()
	}
	ctx.Abort()

}