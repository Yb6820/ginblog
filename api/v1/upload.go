package v1

import (
	"ginblog/models"
	"ginblog/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpLoad(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	url, code := models.UploadFile(&file, fileHeader)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"url":     url,
	})
}
