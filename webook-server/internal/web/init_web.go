package web

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	dao2 "webook-server/internal/repository/dao"
)

var db *gorm.DB

func InitMysql() {
	dsn := "root:root@tcp(127.0.0.1:3306)/webook?charset=utf8mb4&parseTime=True&loc=Local"
	_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	if err = dao2.InitTable(_db); err != nil {
		panic(err)
	}
	db = _db
}