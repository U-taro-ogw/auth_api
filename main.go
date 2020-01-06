package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Hello, World!")

	db := gormConnect()
	defer db.Close()

	d := gin.Default()
	d.POST("/signup", func(c *gin.Context) {
		user := User{}
		err := c.BindJSON(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		db.NewRecord(user)
		db.Create(&user)

		c.JSON(200, gin.H{"message": "signup"})
	})


	d.POST("/signin", func(c *gin.Context) {
		var user User
		var findUser User
		c.BindJSON(&user)
		if err := db.Where("email = ? AND password = ?", user.Email, user.Password).First(&findUser).Error; gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "NotFound"})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "Find!!"})
		}
	})

	d.Run(":8083")
}

func gormConnect() *gorm.DB {
	USER := os.Getenv("MYSQL_USER")
	PASS := os.Getenv("MYSQL_ROOT_PASSWORD")
	PROTOCOL := "tcp(db:3306)"
	DBNAME := os.Getenv("MYSQL_DATABASE")
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open("mysql", CONNECT)

	if err != nil {
		panic("データベース接続に失敗しました。")
	}
	db.AutoMigrate(&User{})

	return db
}

type User struct {
	gorm.Model
	Email string `json:"e-mail"`
	Password string `json:"password"`
}
