package models

import (
	"fmt"
	"ginblog/utils"
	"io"
	"time"

	retalog "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 初始化数据库的函数
var db *gorm.DB
var err error

// 将数据库日志重定向到自定义log
type DBlogWriter struct {
	dblog *logrus.Logger
}

// 实现logger的Printf接口
func (m *DBlogWriter) Printf(format string, v ...interface{}) {
	logstr := fmt.Sprintf(format, v...)
	//利用自定义log记录日志
	m.dblog.Info(logstr)
}
func NewMyWriter() *DBlogWriter {
	log := logrus.New()
	//配置log
	filePath := "dblog/log"
	//创建文件并设置权限
	// src, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755)
	// if err != nil {
	// 	fmt.Println("err:", err)
	// }
	log.SetOutput(io.Discard)
	//设置日志级别
	log.SetLevel(logrus.DebugLevel)
	logWriter, _ := retalog.New(
		filePath+"%Y%m%d.log",
		//设置最多保存多久的日志
		retalog.WithMaxAge(30*24*time.Hour),
		retalog.WithRotationTime(24*time.Hour),
	)
	//设置什么级别的数据可以写到文件里面
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
		logrus.TraceLevel: logWriter,
	}
	//钩子函数
	HOOK := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	log.AddHook(HOOK)

	return &DBlogWriter{dblog: log}
}

func InitDb() {
	// 数据库日志
	log := logger.New(
		NewMyWriter(),
		logger.Config{
			//慢sql阈值
			SlowThreshold: time.Millisecond,
			LogLevel:      logger.Info,
		},
	)
	db, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", utils.DbUser, utils.DbPassword, utils.DbHost, utils.DbPort, utils.DbName)), &gorm.Config{
		Logger: log,
	})
	if err != nil {
		fmt.Println("链接数据库失败！请检查参数:\n", err)
	}
	db.AutoMigrate(&User{}, &Article{}, &Category{}, &Profile{}, &Comment{})
	sqlDB, _ := db.DB()
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)
}
