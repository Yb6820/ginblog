package v1

import (
	"ginblog/models"
	"github.com/gin-gonic/gin"
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
}
