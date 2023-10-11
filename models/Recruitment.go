package models

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"ginblog/utils/errmsg"
	"gorm.io/gorm"
	"time"
)

type MyTime time.Time

func (t MyTime) Value() (driver.Value, error) {
	tTime := time.Time(t)
	return tTime.Format("2006-01-02"), nil
}
func (t *MyTime) Scan(v interface{}) error {
	switch vt := v.(type) {
	case []byte:
		tTime, _ := time.Parse("2006-01-02", string(vt))
		*t = MyTime(tTime)
	default:
		return errors.New("类型错误")
	}
	return nil
}

type Voluntary int

/*
Recruitment表示用户对应的公司和投递的岗位，RecruitStatus表示每一个投递信息的各种状态
*/
type Recruitment struct {
	gorm.Model
	CompanyName  string    `json:"company_name" gorm:"type:varchar(100)" dc:"公司名称"`
	UserId       uint      `json:"user_id"`
	VoluntaryNum Voluntary `json:"voluntary" gorm:"type:tinyint(4)" dc:"第几志愿"`
	DeliveryJob  string    `json:"job_name" gorm:"type:varchar(200)" dc:"投递职位"`
}

type RecruitStatus struct {
	RecruitId      uint   `json:"recruit_id"`
	DeliveryTime   MyTime `json:"delivery_time" dc:"投递日期"`
	DeliveryRes    string `json:"result" gorm:"type:longtext" dc:"投递结果"`
	DeliveryStatus int    `json:"status" dc:"投递状态"`
	Reason         string `json:"reason" gorm:"type:longtext" dc:"原因"`
	Desc           string `json:"desc" gorm:"type:longtext" dc:"描述"`
}

// select recruitment.id,CompanyName,DeliveryJob,DeliveryTime,DeliveryStatus,DeliveryRes,Reason,Desc from
type RecruitRes struct {
	ID             uint   `json:"id"`
	CompanyName    string `json:"company_name" gorm:"type:varchar(100)" dc:"公司名称"`
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

// 新增，修改，下一步的模型
type RecruitmentModel struct {
	CompanyName    string    `json:"company_name" gorm:"type:varchar(100)" dc:"公司名称"`
	VoluntaryNum   Voluntary `json:"voluntary" gorm:"type:tinyint(4)" dc:"第几志愿"`
	DeliveryJob    string    `json:"job_name" gorm:"type:varchar(200)" dc:"投递职位"`
	DeliveryTime   time.Time `json:"delivery_time" dc:"投递日期"`
	DeliveryRes    string    `json:"result" gorm:"type:longtext" dc:"投递结果"`
	DeliveryStatus int       `json:"status" dc:"投递状态"`
	Reason         string    `json:"reason" gorm:"type:longtext" dc:"原因"`
	Desc           string    `json:"desc" gorm:"type:longtext" dc:"描述"`
}

func GetRecruitmentList(req RecruitQueryCond, user UserInfo) (data []Recruitment, total int64, code int) {
	err = db.Table("recruitments").Select("id,company_name,delivery_job,delivery_res,delivery_time,delivery_status,reason,recruit_statuses.desc").Joins("join recruit_statuses on recruitments.id=recruit_statuses.recruit_id").Where(
		"user_id = ? and company_name like ? and delivery_job like ? and delivery_time = ? and delivery_status = ?", user.Userid, req.CompanyName+"%",
		req.DeliveryJob+"%", req.DeliveryTime, req.DeliveryStatus).Limit(req.PageSize).Offset((req.PageNum - 1) * req.PageSize).Order("delivery_time desc").Scan(&data).Error
	fmt.Println("查找数据的err", err)
	if err != nil && err != gorm.ErrRecordNotFound {
		code = errmsg.ERROR
		total = 0
		return
	}
	err = db.Table("recruitments").Select("id,company_name,delivery_job,delivery_res,delivery_time,delivery_status,reason,recruit_statuses.desc").Joins("join recruit_statuses on recruitments.id=recruit_statuses.recruit_id").Where(
		"user_id = ? and company_name like ? and delivery_job like ? and delivery_time = ? and delivery_status = ?", user.Userid, req.CompanyName+"%",
		req.DeliveryJob+"%", req.DeliveryTime, req.DeliveryStatus).Count(&total).Error
	fmt.Println("计数的err", err)
	if err != nil && err != gorm.ErrRecordNotFound {
		code = errmsg.ERROR
		total = 0
		return
	}
	code = errmsg.SUCCESS
	return
}
func AddRecruitment(data RecruitmentModel, user UserInfo) (code int) {
	//事务创建两个表的数据
	recruit := Recruitment{
		CompanyName:  data.CompanyName,
		UserId:       user.Userid,
		VoluntaryNum: data.VoluntaryNum,
		DeliveryJob:  data.DeliveryJob,
	}
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&recruit).Error
		if err != nil {
			return err
		}
		err = tx.Create(&RecruitStatus{
			RecruitId:      recruit.ID,
			DeliveryTime:   MyTime(data.DeliveryTime),
			DeliveryRes:    data.DeliveryRes,
			DeliveryStatus: data.DeliveryStatus,
			Reason:         data.Reason,
			Desc:           data.Desc,
		}).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
