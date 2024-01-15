package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBInit() *gorm.DB {
	dsn := "host=satao.db.elephantsql.com user=ywxstaak password=dJr8F0uFIloDndpzWsEYEXVpJsHBt1CF dbname=ywxstaak port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return db
}
