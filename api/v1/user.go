package v1

import (
	"ginblog/models"
	"ginblog/utils/errmsg"
	"ginblog/utils/validator"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var code int

// 添加用户(还差把邮箱号写入数据库)
func AddUser(c *gin.Context) {
	var data models.User
	var msg string
	_ = c.ShouldBindJSON(&data)
	verify := c.Query("verify")
	//数据校验
	msg, code = validator.Validate(&data)
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": msg,
		})
		return
	}

	code = models.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		code = models.CreateUser(&data, verify)
	}
	/* if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED
	} */
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    nil,
			"message": errmsg.GetErrMsg(code),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		})
	}
}

// 查询单个用户的信息
func GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, code := models.GetUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    user,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询用户列表
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Param("pagesize"))
	pageNuM, _ := strconv.Atoi(c.Query("pagenum"))
	username := c.Query("username")
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNuM == 0 {
		pageNuM = -1
	}
	data, total, code := models.GetUsers(username, pageSize, pageNuM)
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    nil,
			"total":   0,
			"message": errmsg.GetErrMsg(code),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   total,
			"message": errmsg.GetErrMsg(code),
		})
	}
}

// 编辑用户
func EditUser(c *gin.Context) {
	var data models.User
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	code = models.CheckUpUser(id, data.Username)
	if code == errmsg.SUCCESS {
		models.EditUser(id, &data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := models.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 提供一个通过邮箱发送验证码并返回验证码给前端进行校验
func SendEmail(c *gin.Context) {
	toEmail := c.Query("email")
	str, code := models.SendEmail(toEmail)
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": str,
	})
}
