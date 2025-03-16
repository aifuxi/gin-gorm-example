package handler

import (
	"errors"
	"fmt"
	"github.com/aifuxi/gin-gorm-example/model"
	"github.com/aifuxi/gin-gorm-example/pkg/errno"
	"github.com/aifuxi/gin-gorm-example/service"
	"github.com/gin-gonic/gin"
)

func CreateProduct(ctx *gin.Context) {
	var params model.CreateProductRequest

	if err := ctx.Bind(&params); err != nil {
		SendResponse(ctx, errno.ConvertErr(err), nil)
		return
	}

	err := service.NewProductService().CreateProduct(params)
	if err != nil {
		SendResponse(ctx, errno.ConvertErr(err), nil)
		return
	}

	SendResponse(ctx, errno.Success, nil)
}

func QueryProduct(ctx *gin.Context) {
	var params model.QueryProductRequest

	if err := ctx.Bind(&params); err != nil {
		SendResponse(ctx, errno.ConvertErr(err), nil)
		return
	}

	products, total, err := service.NewProductService().QueryProduct(params)
	if err != nil {
		SendResponse(ctx, errno.ConvertErr(err), nil)
		return
	}

	SendResponse(ctx, errno.Success, model.QueryProductResponse{
		Total: total,
		List:  products,
	})
}

func GetProduct(ctx *gin.Context) {
	var params model.GetProductRequest

	if err := ctx.Bind(&params); err != nil {
		SendResponse(ctx, errno.ConvertErr(err), nil)
		return
	}

	product, err := service.NewProductService().GetProduct(params)
	if err != nil {
		if errors.Is(err, errno.NotFoundErr) {
			SendResponse(ctx, errno.NotFoundErr.WithMessage(fmt.Sprintf("product not found, id=%d", params.ID)), nil)
			return
		}
		SendResponse(ctx, errno.ConvertErr(err), nil)
		return
	}

	SendResponse(ctx, errno.Success, product)
}

func DeleteProduct(ctx *gin.Context) {
	var params model.DeleteProductRequest

	if err := ctx.Bind(&params); err != nil {
		SendResponse(ctx, errno.ConvertErr(err), nil)
		return
	}

	err := service.NewProductService().DeleteProduct(params)
	if err != nil {
		if errors.Is(err, errno.NotFoundErr) {
			SendResponse(ctx, errno.NotFoundErr.WithMessage(fmt.Sprintf("product not found, id=%d", params.ID)), nil)
			return
		}
		SendResponse(ctx, errno.ConvertErr(err), nil)
		return
	}

	SendResponse(ctx, errno.Success, nil)
}

func UpdateProduct(ctx *gin.Context) {
	var params model.UpdateProductRequest

	if err := ctx.Bind(&params); err != nil {
		SendResponse(ctx, errno.ConvertErr(err), nil)
		return
	}

	err := service.NewProductService().UpdateProduct(params)
	if err != nil {
		if errors.Is(err, errno.NotFoundErr) {
			SendResponse(ctx, errno.NotFoundErr.WithMessage(fmt.Sprintf("product not found, id=%d", params.ID)), nil)
			return
		}
		SendResponse(ctx, errno.ConvertErr(err), nil)
		return
	}

	SendResponse(ctx, errno.Success, nil)
}
