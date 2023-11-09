package models

import (
	"ginblog/utils/errmsg"
	"gorm.io/gorm"
)

type Voluntary int

func (v Voluntary) ToString() string {
	switch v {
	case 0:
		return "全部"
	case 1:
		return "已投递"
	case 2:
		return "待测评"
	case 3:
		return "已测评"
	case 4:
		return "待笔试"
	case 5:
		return "笔试完成"
	case 6:
		return "笔试挂"
	case 7:
		return "待一面"
	case 8:
		return "一面完成"
	case 9:
		return "一面挂"
	case 10:
		return "待二面"
	case 11:
		return "二面完成"
	case 12:
		return "二面挂"
	case 13:
		return "待三面"
	case 14:
		return "三面完成"
	case 15:
		return "三面挂"
	case 16:
		return "待hr面"
	case 17:
		return "hr面完成"
	case 18:
		return "hr面挂"
	default:
		return ""
	}
}

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
	DeliveryTime   string `json:"delivery_time" gorm:"type:varchar(20)" dc:"投递日期"`
	DeliveryRes    string `json:"result" gorm:"type:longtext" dc:"投递结果"`
	DeliveryStatus int    `json:"status" dc:"投递状态"`
	Reason         string `json:"reason" gorm:"type:longtext" dc:"原因"`
	Desc           string `json:"desc" gorm:"type:longtext" dc:"描述"`
}

// select recruitment.id,CompanyName,DeliveryJob,DeliveryTime,DeliveryStatus,DeliveryRes,Reason,Desc from
type RecruitRes struct {
	ID             uint      `json:"id"`
	CompanyName    string    `json:"company_name" gorm:"type:varchar(100)" dc:"公司名称"`
	DeliveryJob    string    `json:"job_name" gorm:"type:varchar(200)" dc:"投递职位"`
	VoluntaryNum   Voluntary `json:"voluntary_num" `
	DeliveryTime   string    `json:"delivery_time" dc:"投递日期"`
	DeliveryRes    string    `json:"result" gorm:"type:longtext" dc:"投递结果"`
	DeliveryStatus int       `json:"status" dc:"投递状态"`
	Reason         string    `json:"reason" gorm:"type:longtext" dc:"原因"`
	Desc           string    `json:"desc" gorm:"type:longtext" dc:"描述"`
}

// 查询条件
type RecruitQueryCond struct {
	CompanyName    string `json:"company_name" dc:"公司名称"`
	DeliveryJob    string `json:"job_name" dc:"投递职位"`
	DeliveryTime   string `json:"delivery_time" dc:"投递日期"`
	DeliveryStatus int    `json:"status" dc:"投递状态"`
	PageSize       int    `json:"pagesize" dc:"分页大小"`
	PageNum        int    `json:"pagenum" dc:"第几页"`
}

// 新增，修改，下一步的模型
type RecruitmentModel struct {
	RecruitId      uint      `json:"recruit_id" dc:"应聘ID"`
	CompanyName    string    `json:"company_name" gorm:"type:varchar(100)" dc:"公司名称"`
	VoluntaryNum   Voluntary `json:"voluntary" gorm:"type:tinyint(4)" dc:"第几志愿"`
	DeliveryJob    string    `json:"job_name" gorm:"type:varchar(200)" dc:"投递职位"`
	DeliveryTime   string    `json:"delivery_time" dc:"投递日期"`
	DeliveryRes    string    `json:"result" gorm:"type:longtext" dc:"投递结果"`
	DeliveryStatus int       `json:"status" dc:"投递状态"`
	Reason         string    `json:"reason" gorm:"type:longtext" dc:"原因"`
	Desc           string    `json:"desc" gorm:"type:longtext" dc:"描述"`
}

func GetRecruitmentList(req RecruitQueryCond, user UserInfo) (data []RecruitRes, total int64, code int) {
	m := map[string]any{}
	m["user_id"] = user.Userid
	if req.DeliveryStatus != 0 {
		m["delivery_status"] = req.DeliveryStatus
	}
	if req.DeliveryTime != "" {
		m["delivery_time"] = req.DeliveryTime
	}
	subQuery := db.Table("recruit_statuses").Select("recruit_id as t_id,max(delivery_status) as t_status").Group("recruit_id")
	err = db.Table("recruitments").Select("id,company_name,delivery_job,delivery_res,delivery_time,delivery_status,reason,recruit_statuses.desc,voluntary_num").Joins("join recruit_statuses on recruitments.id=recruit_statuses.recruit_id").Joins("join (?) as t on "+
		"recruitments.id = t.t_id and recruit_statuses.delivery_status = t.t_status", subQuery).Where(
		"company_name like ? and delivery_job like ? ", req.CompanyName+"%", req.DeliveryJob+"%").Where(m).Limit(req.PageSize).Offset((req.PageNum - 1) * req.PageSize).Order("delivery_time desc").Scan(&data).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		code = errmsg.ERROR
		total = 0
		return
	}
	err = db.Table("recruitments").Select("id,company_name,delivery_job,delivery_res,delivery_time,delivery_status,reason,recruit_statuses.desc,voluntary_num").Joins("join recruit_statuses on recruitments.id=recruit_statuses.recruit_id").Joins("join (?) as t on "+
		"recruitments.id = t.t_id and recruit_statuses.delivery_status = t.t_status", subQuery).Where(
		"company_name like ? and delivery_job like ? ", req.CompanyName+"%", req.DeliveryJob+"%").Where(m).Count(&total).Error
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
			DeliveryTime:   data.DeliveryTime,
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
func EditRecruitment(data RecruitmentModel, user UserInfo) (code int) {
	recruit := Recruitment{
		CompanyName:  data.CompanyName,
		UserId:       user.Userid,
		VoluntaryNum: data.VoluntaryNum,
		DeliveryJob:  data.DeliveryJob,
	}
	stat := RecruitStatus{
		RecruitId:      data.RecruitId,
		DeliveryTime:   data.DeliveryTime,
		DeliveryRes:    data.DeliveryRes,
		DeliveryStatus: data.DeliveryStatus,
		Reason:         data.Reason,
		Desc:           data.Desc,
	}
	err = db.Transaction(func(tx *gorm.DB) error {
		err = tx.Model(&Recruitment{}).Where("id = ?", data.RecruitId).Updates(&recruit).Error
		if err != nil {
			return err
		}
		err = tx.Model(&RecruitStatus{}).Where("recruit_id = ?", data.RecruitId).Updates(&stat).Error
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
func NextStepRecruitment(data RecruitmentModel) (code int) {
	//判断当前对应的应聘信息的状态是否已经存在
	input := RecruitStatus{}
	err = db.Where("recruit_id = ? and delivery_status = ?", data.RecruitId, data.DeliveryStatus).First(&input).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return errmsg.ERROR_STATUS_EXIST
	}
	input = RecruitStatus{
		RecruitId:      data.RecruitId,
		DeliveryTime:   data.DeliveryTime,
		DeliveryRes:    data.DeliveryRes,
		DeliveryStatus: data.DeliveryStatus,
		Reason:         data.Reason,
		Desc:           data.Desc,
	}
	//直接插入一条信息到状态记录表
	err = db.Create(&input).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
