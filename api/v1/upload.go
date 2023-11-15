package v1

import (
	"fmt"
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
func ShowFile(c *gin.Context) {
	fileName := c.Param("name")
	//从cos服务取出文件流
	data, code := models.ShowFile(fileName)
	fmt.Println("图片数据", string(data))
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
	}
	c.Writer.Write(data)
}
