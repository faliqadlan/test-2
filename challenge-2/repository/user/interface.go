package user

import "be/entities"

type User interface {
	Create(req entities.User) (entities.User, error)
}
