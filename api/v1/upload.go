package v1

import (
	"ginblog/models"
	"ginblog/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpLoad(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	fileSize := fileHeader.Size
	url, code := models.UpLoadFile(file, fileSize)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"url":     url,
	})
}
