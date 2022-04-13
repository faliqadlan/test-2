package product

import "be/entities"

type Product interface {
	Create(user_uid string, req entities.Product) (entities.Product, error)
	Delete(product_uid string) (entities.Product, error)
	Update(product_uid string, req entities.Product) (entities.Product, error)
	Get(user_uid, product_uid string) (GetResponses, error)
}
