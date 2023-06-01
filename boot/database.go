package boot

import (
	"QA_community/global"
	"QA_community/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

func InitMysql() {

	dsn := "root:004515yhb@tcp(127.0.0.1:3306)/db0?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		fmt.Println("连接mysql失败")
		return
	}
	sqlDB := db.DB()
	sqlDB.SetConnMaxIdleTime(180 * time.Second)
	sqlDB.SetConnMaxLifetime(1800 * time.Second)
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetMaxOpenConns(500)
	global.GlobalDb = db
	global.GlobalDb.AutoMigrate(&model.User{})

}

func InitRedis() {

}
