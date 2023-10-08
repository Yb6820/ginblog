package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string

	Zone        int
	AccessKey   string
	SecretKey   string
	Bucket      string
	QiniuServer string

	//redis配置
	RedisAddr     string
	RedisPassword string
	DB            int
	PoolSize      int
	MinIdleConn   int

	//邮箱配置
	MailCompany string
	MailAccount string
	MailPort    int
	MailKey     string
)

func init() {
	file, err := ini.Load("config/localconfig.ini")
	if err != nil {
		fmt.Println("配置文件读取事失误！")
	}
	LoadServer(file)
	LoadData(file)
	LoadQiniu(file)
	LoadMail(file)
	LoadRedis(file)
}

// 读取服务配置
func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("89js82js72")
}

// 读取数据库配置
func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("121.37.246.78")
	DbPort = file.Section("database").Key("DbPort").MustString("3307")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassword = file.Section("database").Key("DbPassword").MustString("181234")
	DbName = file.Section("database").Key("DbName").MustString("ginblog")
}
func LoadQiniu(file *ini.File) {
	Zone = file.Section("qiniu").Key("Zone").MustInt(3)
	AccessKey = file.Section("qiniu").Key("AccessKey").String()
	SecretKey = file.Section("qiniu").Key("SecretKey").String()
	Bucket = file.Section("qiniu").Key("Bucket").String()
	QiniuServer = file.Section("qiniu").Key("QiniuServer").String()
}
func LoadMail(file *ini.File) {
	MailCompany = file.Section("Email").Key("MailCompany").MustString("smtp.126.com")
	MailKey = file.Section("Email").Key("MailKey").String()
	MailPort = file.Section("Email").Key("MailPort").MustInt(25)
	MailAccount = file.Section("Email").Key("MailAccount").MustString("youbet6820@126.com")
}
func LoadRedis(file *ini.File) {
	RedisAddr = file.Section("redis").Key("RedisAddr").MustString("121.37.246.78:6379")
	RedisPassword = file.Section("redis").Key("RedisPassword").MustString("181234")
	DB = file.Section("redis").Key("DB").MustInt(0)
	PoolSize = file.Section("redis").Key("PoolSize").MustInt(30)
	MinIdleConn = file.Section("redis").Key("MinIdleConn").MustInt(30)
}
