package handlers

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTokenHandler(c *gin.Context) {
	// headerのセット
	token := jwt.New(jwt.SigningMethodHS256)
	// 電子署名
	tokenString, _ := token.SignedString([]byte("hoge"))

	c.JSON(http.StatusOK, gin.H{"message": tokenString})
}