package main

import (
	"github.com/aifuxi/gin-gorm-example/cmd/api"
	"github.com/aifuxi/gin-gorm-example/dao"
	"log"
)

func Init() {
	dao.Init()
}

func main() {

	Init()

	err := api.Run()
	if err != nil {
		log.Fatalln("failed to start server", err)
	}
}
