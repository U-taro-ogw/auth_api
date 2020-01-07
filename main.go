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
	dbHandler := DbHandler{
		Db: db,
	}

	d := gin.Default()
	d.POST("/signup", dbHandler.Signup)

	d.POST("/signin", dbHandler.Signin)

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

type DbHandler struct {
	Db *gorm.DB
}

func (h *DbHandler) Signup(c *gin.Context) {
	user := User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	h.Db.NewRecord(user)
	h.Db.Create(&user)

	c.JSON(200, gin.H{"message": "signup"})
}

func (h *DbHandler) Signin(c *gin.Context) {
	var user User
	var findUser User
	c.BindJSON(&user)
	if err := h.Db.Where("email = ? AND password = ?", user.Email, user.Password).First(&findUser).Error; gorm.IsRecordNotFoundError(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "NotFound"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Find!!"})
	}
}