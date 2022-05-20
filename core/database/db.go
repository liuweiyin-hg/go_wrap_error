package database

import (
	app "mydb/core/application"

	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
)

var (
	db *gorm.DB
)

func InitDb() {
	// 初始化数据库连接
	var app = app.GetApp()
	app.Logger().Info("初始化数据库 v")
	var err error
	db, err = gorm.Open("sqlite3", "test.db")
	db.LogMode(true) // show SQL logger
	if err != nil {
		app.Logger().Fatalf("connect to sqlite3 failed")
		return
	}

	iris.RegisterOnInterrupt(func() {
		defer db.Close()
	})
}

func GetDb() *gorm.DB {
	return db
}
