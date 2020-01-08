package main

import (
	"fmt"
	"github.com/U-taro-ogw/auth_api/src/db"
	"github.com/U-taro-ogw/auth_api/src/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

func main() {
	fmt.Println("Hello, World!")

	dbCon := db.DbConnect()
	defer dbCon.Close()
	dbCon.LogMode(true)

	dbHandler := DbHandler{
		Db: dbCon,
	}

	d := gin.Default()
	d.POST("/signup", dbHandler.Signup)

	d.POST("/signin", dbHandler.Signin)

	d.Run(":8083")
}

type DbHandler struct {
	Db *gorm.DB
}

func (h *DbHandler) Signup(c *gin.Context) {
	user := models.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	h.Db.NewRecord(user)
	h.Db.Create(&user)

	c.JSON(200, gin.H{"message": "signup"})
}

func (h *DbHandler) Signin(c *gin.Context) {
	var user models.User
	var findUser models.User
	c.BindJSON(&user)
	fmt.Println("-------------------------====>>>>>>>>>>>>>")
	fmt.Println(user)
	fmt.Println(h.Db.Where("email = ? AND password = ?", user.Email, user.Password))

	if err := h.Db.Where("email = ? AND password = ?", user.Email, user.Password).First(&findUser).Error; gorm.IsRecordNotFoundError(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "NotFound"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Find!!"})
	}
}