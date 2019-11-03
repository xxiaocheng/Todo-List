package models

import (
	"fmt"
	"log"
	"todoList/conf"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func init() {
	InitMysql()
}

func InitMysql() {

	databaseArgs := "%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"

	databaseArgs = fmt.Sprintf(databaseArgs, conf.Config.MySqlConf.Username, conf.Config.MySqlConf.Password, conf.Config.MySqlConf.Host+":"+conf.Config.MySqlConf.Port, conf.Config.MySqlConf.Database)

	db, err := gorm.Open("mysql", databaseArgs)
	if err != nil {
		log.Print(err)
	}
	db.DB().SetMaxIdleConns(20)
	db.LogMode(true)
	DB = db
}
