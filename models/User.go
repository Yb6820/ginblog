package models

import (
	"context"
	"encoding/base64"
	"ginblog/utils"
	"ginblog/utils/errmsg"
	"log"
	"math/rand"
	"strings"
	"time"

	"golang.org/x/crypto/scrypt"
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null " json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(500);not null" json:"password" validate:"required,min=6,max=120" label:"密码"`
	Email    string `gorm:"type:varchar(50);not null" json:"email" validate:"required" label:"邮箱"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色码"`
}

// 查询用户是否存在
func CheckUser(name string) (code int) {
	var user User
	//查询username为name的用户的ID
	db.Select("ID").Where("username = ?", name).First(&user)
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

// CheckUpUser 更新查询
func CheckUpUser(id int, name string) (code int) {
	var user User
	db.Select("id, username").Where("username = ?", name).First(&user)
	if user.ID == uint(id) {
		return errmsg.SUCCESS
	}
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED //1001
	}
	return errmsg.SUCCESS
}

// hook函数在数据进入数据库之前进行加密
func (u *User) BeforeCreate(_ *gorm.DB) error {
	u.Password = ScryptPwd(u.Password)
	//u.Role=2
	return nil
}

// 新增用户
func CreateUser(data *User, verify string) int {
	//data.Password = ScryptPwd(data.Password)
	//校验邮箱验证码
	ctx := context.TODO()
	resg, err := Red.Get(ctx, data.Email).Result()
	if err != nil {
		return errmsg.REDIS_GET_VERIFY_ERROR
	}
	if verify != resg {
		return errmsg.ERROR_VERIFY_NOT_SAME
	}
	err = db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询单个用户
func GetUser(id int) (User, int) {
	var user User
	err := db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, errmsg.ERROR
	}
	return user, errmsg.SUCCESS
}

// 查询用户列表(为了避免数据过多，可以先进行分页)
func GetUsers(username string, pageSize int, pageNum int) ([]User, int, int) {
	var users []User
	var total int64
	//分页获取后端数据
	//username不为空时模糊查询
	if username != "" {
		err := db.Where("username like ?", username+"%").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
		db.Where("username like ?", username+"%").Count(&total)
		if err != nil {
			return nil, 0, errmsg.ERROR
		}
		return users, int(total), errmsg.SUCCESS
	}
	//否着返回所有数据
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	db.Model(&User{}).Count(&total)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, errmsg.ERROR
	}
	return users, int(total), errmsg.SUCCESS
}

// 编辑用户信息(密码独立出去)
func EditUser(id int, data *User) int {
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err := db.Model(&User{}).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除用户(存在删除同一个用户多次都返回成功的错误)
func DeleteUser(id int) int {
	var user User
	err := db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 密码加密(目前较为流行的加密方式)
func ScryptPwd(password string) string {
	const KeyLen = 10
	salt := []byte{12, 32, 4, 6, 66, 22, 222, 11}
	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(HashPw)
	return fpw
}

// 登陆验证
func CheckLogin(username string, password string) int {
	var user User
	db.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if ScryptPwd(password) != user.Password {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 1 {
		return errmsg.ERROR_USER_NOT_RIGHT
	}
	return errmsg.SUCCESS
}

// CheckLoginFront 前台登录
func CheckLoginFront(username string, password string) (User, int) {
	var user User

	db.Where("username = ?", username).First(&user)

	if user.ID == 0 {
		return user, errmsg.ERROR_USER_NOT_EXIST
	}
	if ScryptPwd(password) != user.Password {
		return user, errmsg.ERROR_PASSWORD_WRONG
	}
	return user, errmsg.SUCCESS
}

// 向注册账号的人发送验证码
func SendEmail(toEmail string) (string, int) {
	m := gomail.NewMessage()
	rand.Seed(time.Now().Unix())
	verifybyte := make([]string, 6)
	var strs []string = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	for i := 0; i < 6; i++ {
		verifybyte[i] = strs[rand.Intn(62)]
	}
	ctx := context.TODO()
	verify := strings.Join(verifybyte, "")
	//设置超时时间为5分钟
	_, err := Red.Set(ctx, toEmail, verify, 5*time.Minute).Result()
	if err != nil {
		return "", errmsg.REDIS_SET_VERIFY_ERROR
	}
	//发送人
	m.SetHeader("From", utils.MailAccount)
	//收件人
	m.SetHeader("To", toEmail)
	//主体
	m.SetHeader("Subject", "Youbet's Blog")
	//内容
	m.SetBody("text/html", "<h4>您正在Youbet's Blog里注册账号</h4><h6>验证码为： </h6><h4>"+verify+"</h4>")
	//将验证码存到Redis并设置过期时间

	//KSZPBUCAAKYDSLXK(授权码)
	d := gomail.NewDialer(utils.MailCompany, utils.MailPort, utils.MailAccount, utils.MailKey)
	//
	//d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	//发送邮件
	if err := d.DialAndSend(m); err != nil {
		return "", errmsg.EMAIL_ERROR
	}
	return verify, errmsg.SUCCESS
}
