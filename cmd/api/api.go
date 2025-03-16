package api

import (
	"github.com/aifuxi/gin-gorm-example/cmd/api/handler"
	"github.com/gin-gonic/gin"
)

func Run() error {
	router := gin.Default()

	v1 := router.Group("/api/v1")

	v1.GET("/ping", handler.Ping)

	product1 := v1.Group("/product")

	{
		product1.POST("/create", handler.CreateProduct)
		product1.POST("/delete", handler.DeleteProduct)
		product1.POST("/update", handler.UpdateProduct)
		product1.POST("/list", handler.QueryProduct)
		product1.POST("/detail", handler.GetProduct)
	}

	return router.Run("127.0.0.1:8080")
}
