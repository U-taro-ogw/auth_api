package main

import (
	"fmt"
	"github.com/U-taro-ogw/auth_api/src/db"
	"github.com/U-taro-ogw/auth_api/src/handlers"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

func main() {
	fmt.Println("Hello, World!")

	dbCon := db.DbConnect()
	defer dbCon.Close()
	dbCon.LogMode(true)

	dbHandler := handlers.UserHandler{
		Db: dbCon,
	}

	d := gin.Default()
	d.POST("/signup", dbHandler.Signup)

	d.POST("/signin", dbHandler.Signin)

	d.GET("/redis", func(c *gin.Context) {
		r := db.RedisConnect()
		c.JSON(http.StatusOK, gin.H{"redis": db.Get("hoge", r)})
	})

	d.Run(":8083")
}
