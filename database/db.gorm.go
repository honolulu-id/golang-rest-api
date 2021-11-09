package database

import (
	"fmt"
	"golang-rest-api/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConfigMysql() (*gorm.DB, error) {

	conf := config.GetConfig()
	konek := conf.DB_USERNAME + `:` + conf.DB_PASSWORD + `@tcp(`+ conf.DB_HOST + `:` + conf.DB_PORT +`)/` + conf.DB_NAME +`?charset=utf8mb4&parseTime=True&loc=Local`
	db, err := gorm.Open(mysql.Open(konek), &gorm.Config{})

	if err != nil {
		fmt.Println("db error", err.Error())
	}

	return db, err
}

func KonekMysql() (*gorm.DB, error) {
	db, err := ConfigMysql()
	return db, err
}