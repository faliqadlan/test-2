package product

import (
	"be/entities"
	"strconv"

	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Repo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) Create(user_uid string, req entities.Product) (entities.Product, error) {

	req.User_uid = user_uid

	var res = r.db.Unscoped().Model(&entities.Product{}).Where("user_uid = ?", user_uid).Scan(&[]entities.Product{})
	var uid string = user_uid + "-" + strconv.Itoa(int(res.RowsAffected)+1)

	req.Product_uid = uid

	res = r.db.Model(&entities.Product{}).Create(&req)

	if res.Error != nil {
		return entities.Product{}, res.Error
	}

	return req, nil
}

func (r *Repo) Delete(product_uid string) (entities.Product, error) {

	var resInit entities.Product

	var res = r.db.Model(&entities.Product{}).Where("product_uid = ?", product_uid).Delete(&resInit)

	if res.RowsAffected == 0 {
		return entities.Product{}, gorm.ErrRecordNotFound
	}

	return resInit, nil
}

func (r *Repo) Update(product_uid string, req entities.Product) (entities.Product, error) {

	var res = r.db.Model(&entities.Product{}).Where("product_uid = ?", product_uid).Updates(entities.Product{
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
		Stock:       req.Stock,
		Image:       req.Image,
	})

	if res.RowsAffected == 0 {
		return entities.Product{}, gorm.ErrRecordNotFound
	}

	return req, nil
}

func (r *Repo) Get(user_uid, product_uid string) (GetResponses, error) {

	switch {
	case user_uid != "":
		user_uid = "products.user_uid = '" + user_uid + "'"
	default:
		user_uid = "products.user_uid != '" + "user_uid" + "'"
	}

	switch {
	case product_uid != "":
		product_uid = " AND products.product_uid = '" + product_uid + "'"
	default:
		product_uid = " AND products.product_uid != '" + "product_uid" + "'"
	}

	var condition = user_uid + product_uid

	var result GetResponses

	var res = r.db.Model(&entities.Product{}).Select("products.product_uid as Product_uid ,users.name as NameUser, products.Name as NameProduct, price as Price, description as Description, stock as Stock, products.image as Image").Joins("inner join users on products.user_uid = users.user_uid").Where(condition).Find(&result.Responses)
	// log.Info(product_uid)
	if product_uid != " AND products.product_uid != 'product_uid'" && res.RowsAffected == 0 {
		return GetResponses{}, gorm.ErrRecordNotFound
	}

	if res.Error != nil {
		return GetResponses{}, res.Error
	}

	return result, nil

}
