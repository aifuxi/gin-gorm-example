package db

import (
	"github.com/aifuxi/gin-gorm-example/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln("failed to connect db", err)
	}

	// 自动迁移表
	err = DB.AutoMigrate(&model.Product{})
	if err != nil {
		log.Fatalln("failed to auto migrate db", err)
	}
}
