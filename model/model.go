package model

import (
	"time"
	"videowebsite/util"

	"github.com/jinzhu/gorm"

	//
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB 数据库链接单例
var DB *gorm.DB

// Database 在中间件中初始化mysql链接
func Database(connString string) {
	db, err := gorm.Open("mysql", connString)
	db.LogMode(true)  //启用logger,显示详细信息
	// Error
	if err != nil {
		util.Log().Panic("连接数据库不成功", err)
	}
	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(50)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	DB = db

	migration()
}

func migration() {
	// 自动迁移模式
	DB.AutoMigrate(&User{},&Video{})
	//gorm可以支持自动迁移，也就是自动的表结构迁移，只会创建表，补充缺少的列，缺少的索引。但并不会更改已经存在的列类型，也不会删除不再用的列，
	//这样设计的目的是为了保护已存在的数据。可以同时针对多个表进行迁移设置。
}

