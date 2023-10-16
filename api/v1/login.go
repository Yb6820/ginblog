package v1

import (
	"fmt"
	"ginblog/middleware"
	"ginblog/models"
	"ginblog/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 登录接口
func Login(c *gin.Context) {
	user := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}
	c.ShouldBindJSON(&user)
	var token string
	data, code := models.CheckLogin(user.Username, user.Password)
	if code == errmsg.SUCCESS {
		token, code = middleware.SetToken(data.ID, data.Username)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"token":   token,
	})
}

// LoginFront 前台登录
func LoginFront(c *gin.Context) {
	var formData models.User
	_ = c.ShouldBindJSON(&formData)
	var code int

	fmt.Println(formData.Username, formData.Password)
	formData, code = models.CheckLoginFront(formData.Username, formData.Password)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    formData.Username,
		"id":      formData.ID,
		"message": errmsg.GetErrMsg(code),
	})
}
