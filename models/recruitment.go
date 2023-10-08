package models

import (
	"database/sql/driver"
	"errors"
	"gorm.io/gorm"
	"time"
)

type MyTime time.Time

func (t MyTime) Value() (driver.Value, error) {
	tTime := time.Time(t)
	return tTime.Format("2006-01-02 15:04:05"), nil
}
func (t *MyTime) Scan(v interface{}) error {
	switch vt := v.(type) {
	case []byte:
		tTime, _ := time.Parse("2006-01-02 15:04:05", string(vt))
		*t = MyTime(tTime)
	default:
		return errors.New("类型错误")
	}
	return nil
}

type Recruitment struct {
	gorm.Model
	CompanyName    string `json:"company_name" gorm:"type:varchar(100)" dc:"公司名称"`
	DeliveryJob    string `json:"job_name" gorm:"type:varchar(200)" dc:"投递职位"`
	DeliveryTime   MyTime `json:"delivery_time" dc:"投递日期"`
	DeliveryRes    string `json:"result" gorm:"type:longtext" dc:"投递结果"`
	DeliveryStatus int    `json:"status" dc:"投递状态"`
	Reason         string `json:"reason" gorm:"type:longtext" dc:"原因"`
	Desc           string `json:"desc" gorm:"type:longtext" dc:"描述"`
}
