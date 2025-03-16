package handler

import (
	"github.com/aifuxi/gin-gorm-example/pkg/errno"
	"github.com/aifuxi/gin-gorm-example/service"
	"github.com/gin-gonic/gin"
)

func Ping(ctx *gin.Context) {
	data, err := service.NewPingService().Ping()
	if err != nil {
		SendResponse(ctx, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(ctx, errno.Success, data)
}
