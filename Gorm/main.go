package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)
type Video struct {
	gorm.Model
	Title     string
	Commit    string
	Coin      int `gorm:"default:0"`
	ThumbsUp      int `gorm:"default:0"`
	Author int
	Userid int
	User    User `gorm:"foreignKey:Author"`
	Collect   int`gorm:"default:0"`
	Times     int `gorm:"default:0"`
	Tag       string
}
type User struct {
	Username string `form:"username"`
	Phone   string `form:"phone"`
	Coin    int    `gorm:"default:0"`
	Follows int    `form:"follows" gorm:"default:0"`
	Fans    int    `form:"fans" gorm:"default:0"`
	gorm.Model
}

func main() {
	dsn := "root:dengjie123@tcp(127.0.0.1:3306)/text?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err!=nil{
		log.Panicln(err)
	}
	db.AutoMigrate(&Video{})
	/*err=db.AutoMigrate(&User{})
	if err!=nil{
		log.Panicln(err)
	}
	//用户的添加
	result:=db.Create(&User{
		Username: "邓捷3",
		Phone: "123",
	})
	err=result.Error
	if err!=nil{
		log.Panicln(err)
	}*/
	//用户的查询
	var user1 User
	//var users []User
	//db.First(&user1,3)
	//db.Find(&users)
	//db.Order("id desc").Where("username like","%3").Find(&users)
	//用户的增删改查，投币事物，外键视频库，看发布的视频
	//db.Model(&User{}).Where("id","3").Update("username","mixinkexue")
	//db.Model(&User{}).Where("username","邓捷").Update("coin",1)
	//db.Delete(&User{},[]int{1,2,3})
	//db.Unscoped().Delete(&User{},[]int{10,11})
	db.Last(&user1)
	video:=Video{
		Title:    "第2个视频",
		Commit:   "你好啊",
		Coin:     0,
		ThumbsUp: 0,
		User:     user1,
		Collect:  0,
		Times:    0,
		Tag:      "生活",
	}
	db.Create(&video)
	var video1 Video

	err=db.Model(&video1).Association("User").Error
	if err!=nil{
		log.Panic(err)
	}
	fmt.Println(user1,video1,err)
	//使用事物投币
	db.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		var user User
		if err:=tx.First(&user,12) .Error;err!=nil{
			return err
		}
		user.Coin--
		if user.Coin<0{

			return nil
		}
		if err:=tx.Model(&user).Update("coin",user.Coin).Error;err!=nil{
			return err
		}
		var video Video
		if err:=tx.First(&video,2) .Error;err!=nil{
			return err
		}
		video.Coin++
		if err:=tx.Model(&video).Update("coin",video.Coin).Error;err!=nil{
			return err
		}
		// 返回 nil 提交事务
		return nil
	})
}