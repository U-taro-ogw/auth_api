package db

import (
	"github.com/U-taro-ogw/auth_api/src/models"
	"github.com/jinzhu/gorm"
	"os"
)

func DbConnect() *gorm.DB {
	USER := os.Getenv("MYSQL_USER")
	PASS := os.Getenv("MYSQL_ROOT_PASSWORD")
	PROTOCOL := "tcp(db:3306)"
	DBNAME := os.Getenv("MYSQL_DATABASE")
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open("mysql", CONNECT)

	if err != nil {
		panic("データベース接続に失敗しました。")
	}
	db.AutoMigrate(&models.User{})

	return db
}