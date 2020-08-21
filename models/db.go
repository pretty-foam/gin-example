package models

import (
	"fmt"
	"log"
	"time"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var db *gorm.DB 

//SetUp 数据库初始化
func SetUp() {
	var err error
	log.Println("连接数据库中~~~")
	db, err = gorm.Open(viper.GetString("database.type"), fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.host"),
		viper.GetString("database.port"),
		viper.GetString("database.dbName"),
	))
	if err != nil {
		log.Fatalf("数据库连接错误:%v", err)
	} else {
		log.Println("数据库连接成功~~~")
	}
	// 启用Logger，显示详细日志
	db.LogMode(true)
	// 不设置这个参数，gorm会把表名转义后加个s，导致找不到数据库的表
	db.SingularTable(true)
	// 最大空闲连接数
	db.DB().SetMaxIdleConns(10)
	// 数据库最大连接
	db.DB().SetMaxOpenConns(100)
	// 超时
	db.DB().SetConnMaxLifetime(time.Second * 30)
	db.AutoMigrate(&Info{})
	defer db.Close()
}
