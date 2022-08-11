package user

import "bitly_backend/app/models"

type UserRepository interface {
	Register(username, email, password string) (user models.User, err error)
	Login(find, password string) (user models.User, err error)
	GetAllProductByUserID(user_id float64) (user models.User, err error)
	GetInfoUser(user_id float64) (user models.User, err error)
}
