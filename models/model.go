package models

import (
	"fmt"
	"github.com/ririkizzu/spider007/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

type Model struct {
	Id        int   `gorm:"primaryKey"`
	CreatedAt int64 `gorm:"autoCreateTime:milli;column:created_at"`
	UpdatedAt int64 `gorm:"autoUpdateTime:milli;column:updated_at"`
}

var db *gorm.DB

func InitDBConn() {
	var err error
	var mysqlDsn string
	mysqlDsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.DBName,
	)

	db, err = gorm.Open(mysql.Open(mysqlDsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatalf("model.InitDBConn err: %v", err)
	}

	//设置连接池信息
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("SetConns err: %v", err)
	}
	//微信云托管数据库默认wait_timeout为3600s，需要设置连接可复用的最大时间小于3600s，否则会出现报错: closing bad idle connection: EOF
	sqlDB.SetConnMaxLifetime(time.Second * 1800)
}
