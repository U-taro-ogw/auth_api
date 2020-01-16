package main

import (
	"fmt"
	"github.com/U-taro-ogw/auth_api/src/db"
	"github.com/U-taro-ogw/auth_api/src/handlers"
	//"github.com/U-taro-ogw/auth_api/src/modules"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

func main() {
	fmt.Println("Hello, World!")

	dbCon := db.DbConnect()
	defer dbCon.Close()
	dbCon.LogMode(true)
	userHandler := handlers.UserHandler{
		Db: dbCon,
	}

	redisCon := db.RedisConnect()
	redisHandler := handlers.RedisHandler{
		Redis: redisCon,
	}

	d := gin.Default()
	d.POST("/signup", userHandler.Signup)
	d.POST("/signin", userHandler.Signin)
	d.GET("/redis", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"redis": redisHandler.Get("hoge")})
	})
	//d.GET("/auth", modules.GetTokenHandler)

	d.Run(":8083")
}
