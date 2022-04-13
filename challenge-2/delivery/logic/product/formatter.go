package product

import "be/entities"

type Req struct {
	Name        string `json:"name" form:"name"  validate:"required"`
	Description string `json:"description" form:"description"`
	Price       string `json:"price" form:"price"  validate:"required"`
	Stock       int    `json:"stock" form:"stock"  validate:"required"`
	Image       string
}

func (r *Req) ToProduct() *entities.Product {
	return &entities.Product{
		Name:        r.Name,
		Description: r.Description,
		Price:       r.Price,
		Stock:       r.Stock,
		Image:       r.Image,
	}
}
