package modle

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
	"websocket/dao"
)
type User struct {
	Username string`json:"username" form:"username" gorm:"unique;not null"'`
	Password string`json:"password" form:"password" gorm:"not null"`
	gorm.Model
}

func Register(user User)bool {
	db := dao.DB
	err:=db.AutoMigrate(&User{})

	if err != nil {
		fmt.Println(err)
		return false
	}
	result:=db.Create(&user)
	if err:=result.Error;err!=nil{
		fmt.Println(err)
		return false
	}
	return true
}
func Login(user User) bool{
	db:=dao.DB
	var user1 User
	db.Where("username = ?",user.Username).First(&user1)
	fmt.Println(user)
	fmt.Println(user1)
	if strings.EqualFold(user.Password,user1.Password){
		return true
	}else {
		return false
	}
}