package usecase

import (
	"bitly_backend/app/models"
	"bitly_backend/app/usecase/request"
)

type Usecase interface {
	CreateUrl(body request.CreateUrl) (out models.Product, err error)
	UpdateUrl(body request.UpdateUrl) (out models.Product, err error)
	GetListProductByUserID(body float64) (out []models.Product, err error)
	Register(body request.Register) (out models.User, err error)
	Login(body request.Login) (out string, err error)
	PatchGetClick(body int) (out models.Product, err error)
	CreateUrlWithoutUser(body string) (str string, err error)
	GetInfoUser(user_id float64) (out models.User, err error)
	GetOneProduct(id int) (product models.Product, err error)
}
