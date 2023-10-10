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

/*
Recruitment表示用户对应的公司和投递的岗位，RecruitStatus表示每一个投递信息的各种状态
*/
type Recruitment struct {
	gorm.Model
	CompanyName string `json:"company_name" gorm:"type:varchar(100)" dc:"公司名称"`
	UserId      uint   `json:"user_id"`
	DeliveryJob string `json:"job_name" gorm:"type:varchar(200)" dc:"投递职位"`
}

type RecruitStatus struct {
	RecruitId      uint   `json:"recruit_id"`
	DeliveryTime   MyTime `json:"delivery_time" dc:"投递日期"`
	DeliveryRes    string `json:"result" gorm:"type:longtext" dc:"投递结果"`
	DeliveryStatus int    `json:"status" dc:"投递状态"`
	Reason         string `json:"reason" gorm:"type:longtext" dc:"原因"`
	Desc           string `json:"desc" gorm:"type:longtext" dc:"描述"`
}
type RecruitRes struct {
	ID             uint   `json:"id"`
	CompanyName    string `json:"company_name" gorm:"type:varchar(100)" dc:"公司名称"`
	UserId         uint   `json:"user_id"`
	DeliveryJob    string `json:"job_name" gorm:"type:varchar(200)" dc:"投递职位"`
	DeliveryTime   MyTime `json:"delivery_time" dc:"投递日期"`
	DeliveryRes    string `json:"result" gorm:"type:longtext" dc:"投递结果"`
	DeliveryStatus int    `json:"status" dc:"投递状态"`
	Reason         string `json:"reason" gorm:"type:longtext" dc:"原因"`
	Desc           string `json:"desc" gorm:"type:longtext" dc:"描述"`
}

// 查询条件
type RecruitQueryCond struct {
	CompanyName    string `json:"company_name" dc:"公司名称"`
	DeliveryJob    string `json:"job_name" dc:"投递职位"`
	DeliveryTime   MyTime `json:"delivery_time" dc:"投递日期"`
	DeliveryStatus int    `json:"status" dc:"投递状态"`
	PageSize       int    `json:"pagesize" dc:"分页大小"`
	PageNum        int    `json:"pagenum" dc:"第几页"`
}

func GetRecruitmentList(req RecruitQueryCond) (data []Recruitment, total int, code int) {
	return nil, 0, 0
}
