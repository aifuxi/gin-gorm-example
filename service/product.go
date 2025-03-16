package service

import (
	"github.com/aifuxi/gin-gorm-example/dao/db"
	"github.com/aifuxi/gin-gorm-example/model"
	"github.com/aifuxi/gin-gorm-example/pkg/errno"
)

type ProductService struct {
}

func NewProductService() *ProductService {
	return &ProductService{}
}

func (s *ProductService) CreateProduct(req model.CreateProductRequest) error {
	Product := &model.Product{
		Code:  req.Code,
		Price: req.Price,
	}

	err := db.CreateProduct([]*model.Product{Product})
	if err != nil {
		return err
	}

	return nil
}

func (s *ProductService) DeleteProduct(req model.DeleteProductRequest) error {
	_, err := s.GetProduct(model.GetProductRequest{ID: req.ID})
	if err != nil {
		return err
	}

	err = db.DeleteProduct([]int64{req.ID})

	if err != nil {
		return err
	}
	return nil
}

func (s *ProductService) QueryProduct(req model.QueryProductRequest) ([]*model.Product, int64, error) {
	products, total, err := db.QueryProduct(req.SearchKey, req.PageSize, (req.Page-1)*req.PageSize)
	if err != nil {
		return nil, 0, err
	}

	return products, total, nil
}

func (s *ProductService) UpdateProduct(req model.UpdateProductRequest) error {
	_, err := s.GetProduct(model.GetProductRequest{ID: req.ID})
	if err != nil {
		return err
	}

	err = db.UpdateProduct(req.ID, req.Code, req.Price)
	if err != nil {
		return err
	}
	return nil
}

func (s *ProductService) GetProduct(req model.GetProductRequest) (*model.Product, error) {
	products, err := db.MultiGetProducts([]int64{req.ID})
	if err != nil {
		return nil, err
	}

	if len(products) == 0 {
		return nil, errno.NotFoundErr
	}

	return products[0], err
}
