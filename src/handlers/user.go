package handlers

import (
	"github.com/U-taro-ogw/auth_api/src/models"
	"github.com/U-taro-ogw/auth_api/src/modules"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"net/http"
)

type UserHandler struct {
	Db *gorm.DB
	Redis redis.Conn
}

func (h *UserHandler) Signup(c *gin.Context) {
	user := models.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	h.Db.NewRecord(user)
	h.Db.Create(&user)

	c.JSON(http.StatusCreated, gin.H{"message": "signup"})
}

func (h *UserHandler) Signin(c *gin.Context) {
	var userParam models.User
	var findUser models.User
	c.BindJSON(&userParam)

	if err := h.Db.Where("email = ? AND password = ?", userParam.Email, userParam.Password).First(&findUser).Error; gorm.IsRecordNotFoundError(err) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	} else {
		jwtToken := modules.GetTokenHandler()
		modules.SetRedis(h.Redis, jwtToken, "111")
		c.JSON(http.StatusOK, gin.H{"jwt": jwtToken})
	}
}
