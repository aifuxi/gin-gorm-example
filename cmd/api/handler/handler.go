package handler

import (
	"github.com/aifuxi/gin-gorm-example/pkg/errno"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Response struct {
	Code      int       `json:"code"`
	Msg       string    `json:"msg"`
	Data      any       `json:"data"`
	Timestamp time.Time `json:"ts"` // 响应时间
}

func SendResponse(ctx *gin.Context, err error, data any) {
	Err := errno.ConvertErr(err)

	ctx.JSON(http.StatusOK, Response{
		Code:      Err.ErrCode,
		Msg:       Err.ErrMsg,
		Data:      data,
		Timestamp: time.Now(),
	})
}
