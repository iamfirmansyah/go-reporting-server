package app

import (
	"go-reporting-server/helper"

	log "github.com/sirupsen/logrus"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var err error

func ConnectMysql(connectionString string) {
	Instance, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	helper.PanicIfError(err)

	log.Infof("Connected to Database Mysql...")
}
