package user

import "be/entities"

type Req struct {
	UserName string `json:"userName" form:"userName" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Name     string `json:"name" form:"name"  validate:"required"`
	Image    string
}

func (r *Req) ToUser() *entities.User {
	return &entities.User{
		UserName: r.UserName,
		Email:    r.Email,
		Password: r.Password,
		Name:     r.Name,
		Image:    r.Image,
	}
}
