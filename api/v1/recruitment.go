package v1

import (
	"fmt"
	"ginblog/models"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func GetRecruitmentList(c *gin.Context) {
	req := models.RecruitQueryCond{}
	req.CompanyName = c.Query("company_name")
	req.DeliveryJob = c.Query("delivery_job")
	t, _ := time.Parse("", c.Query("delivery_time"))
	req.DeliveryTime = models.MyTime(t)
	req.DeliveryStatus, _ = strconv.Atoi(c.Query("delivery_status"))
	req.PageSize, _ = strconv.Atoi(c.Query("pagesize"))
	req.PageNum, _ = strconv.Atoi(c.Query("pagenum"))
	user, _ := c.Get("user")
	data, total, code := models.GetRecruitmentList(req, user.(models.UserInfo))
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
			"data":    nil,
			"total":   total,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"data":    data,
		"total":   total,
	})
}

// 新增应聘信息
func AddRecruitment(c *gin.Context) {
	var entity models.RecruitmentModel
	c.ShouldBindJSON(&entity)
	fmt.Printf("新增应聘信息时获取到的数据%+v\n", entity)
	user, _ := c.Get("user")
	code := models.AddRecruitment(entity, user.(models.UserInfo))
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
