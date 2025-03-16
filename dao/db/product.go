package db

import (
	"github.com/aifuxi/gin-gorm-example/model"
)

func CreateProduct(data []*model.Product) error {
	if err := DB.Create(data).Error; err != nil {
		return err
	}
	return nil
}

func DeleteProduct(ids []int64) error {
	return DB.Where("id in ?", ids).Delete(&model.Product{}).Error
}

func UpdateProduct(id int64, code *string, price *uint) error {
	params := map[string]any{}

	if code != nil {
		params["code"] = *code
	}

	if price != nil {
		params["price"] = *price
	}

	return DB.Model(&model.Product{}).Where("id = ?", id).Updates(params).Error
}

func MultiGetProducts(ids []int64) ([]*model.Product, error) {
	var res []*model.Product

	if len(ids) == 0 {
		return res, nil
	}
	if err := DB.Where("id in ?", ids).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

func QueryProduct(searchKey *string, limit, offset int) ([]*model.Product, int64, error) {
	var total int64
	var res []*model.Product

	conn := DB.Model(&model.Product{})

	if searchKey != nil {
		conn = conn.Where("code like ?", "%"+*searchKey+"%")
	}

	if err := conn.Count(&total).Error; err != nil {
		return res, total, err
	}

	if err := conn.Limit(limit).Offset(offset).Find(&res).Error; err != nil {
		return res, total, err
	}

	return res, total, nil
}
