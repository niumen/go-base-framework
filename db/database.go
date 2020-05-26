package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"log"
	"logger"
	"os"
	"time"
)

var FengConnect *gorm.DB
var HaoConnect *gorm.DB
var YuanConnect *gorm.DB

func InitDB() {
	FConnect()
	HConnect()
	YConnect()
}

func FConnect() {
	addr := viper.GetString("mysql_addr_f")
	var err error
	FengConnect, err = gorm.Open("mysql", addr)
	if err != nil {
		logrus.Error(err)
	}
	f, _ := os.OpenFile("./zhifeng/logs/gorm.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	gormLog := log.New(io.MultiWriter(f, os.Stdout), "\r\n", 0)

	FengConnect.SetLogger(logger.Loggers{gormLog})

	if viper.GetString("log_level") == "debug" {
		FengConnect.LogMode(true)
	}
	//设置连接最大存活时间
	FengConnect.DB().SetConnMaxLifetime(time.Minute * 10)
	//全局禁用复数形式
	FengConnect.SingularTable(true)
}

func HConnect() {
	addr := viper.GetString("mysql_addr_h")
	var err error
	HaoConnect, err = gorm.Open("mysql", addr)
	if err != nil {
		logrus.Error(err)
	}
	f, _ := os.OpenFile("./zhihao/logs/gorm.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	gormLog := log.New(io.MultiWriter(f, os.Stdout), "\r\n", 0)

	HaoConnect.SetLogger(logger.Loggers{gormLog})

	if viper.GetString("log_level") == "debug" {
		HaoConnect.LogMode(true)
	}
	//设置连接最大存活时间
	HaoConnect.DB().SetConnMaxLifetime(time.Minute * 10)
	//全局禁用复数形式
	HaoConnect.SingularTable(true)
}

func YConnect() {
	addr := viper.GetString("mysql_addr_y")
	var err error
	YuanConnect, err = gorm.Open("mysql", addr)
	if err != nil {
		logrus.Error(err)
	}
	f, _ := os.OpenFile("./zhiyuan/logs/gorm.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	gormLog := log.New(io.MultiWriter(f, os.Stdout), "\r\n", 0)

	YuanConnect.SetLogger(logger.Loggers{gormLog})

	if viper.GetString("log_level") == "debug" {
		YuanConnect.LogMode(true)
	}
	//设置连接最大存活时间
	YuanConnect.DB().SetConnMaxLifetime(time.Minute * 10)
	//全局禁用复数形式
	YuanConnect.SingularTable(true)
}
