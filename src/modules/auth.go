package modules

import (
	jwt "github.com/dgrijalva/jwt-go"
	//"github.com/gin-gonic/gin"
	//"net/http"
	"fmt"
	"reflect"
)

func GetTokenHandler() string {
	// headerのセット
	token := jwt.New(jwt.SigningMethodHS256)
	// 電子署名
	tokenString, _ := token.SignedString([]byte("hoge"))

	fmt.Println("------------------------->>>>>>>>>>>>>>>>")
	fmt.Println(reflect.TypeOf(tokenString))
	fmt.Println(tokenString)

	//c.JSON(http.StatusOK, gin.H{"message": tokenString})
	return tokenString
}