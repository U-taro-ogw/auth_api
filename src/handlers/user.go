package handlers

import (
	"github.com/U-taro-ogw/auth_api/src/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

type UserHandler struct {
	Db *gorm.DB
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
		c.JSON(http.StatusNotFound, gin.H{"error": "NotFound"})
	} else {
		// TODO jwtトークンの発行 -> redis保存 -> responseにjwtトークンを含める
		c.JSON(http.StatusOK, gin.H{"message": "Find!!"})
	}
}
