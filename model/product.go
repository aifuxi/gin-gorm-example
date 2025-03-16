package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type CreateProductRequest struct {
	Code  string `json:"code"`
	Price uint   `json:"price"`
}

type QueryProductRequest struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`

	SearchKey *string `json:"search_key"`
}

type QueryProductResponse struct {
	Total int64      `json:"total"`
	List  []*Product `json:"list"`
}

type UpdateProductRequest struct {
	ID    int64   `json:"id"`
	Code  *string `json:"code"`
	Price *uint   `json:"price"`
}

type GetProductRequest struct {
	ID int64 `json:"id"`
}

type DeleteProductRequest struct {
	ID int64 `json:"id"`
}
