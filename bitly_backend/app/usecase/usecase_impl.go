package usecase

import (
	"bitly_backend/app/models"
	"bitly_backend/app/repository/product"
	"bitly_backend/app/repository/user"
	"bitly_backend/app/shared/function"
	"bitly_backend/app/usecase/request"
	"errors"
	"fmt"
	"net/mail"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gomodul/envy"
)

type usecase struct {
	userRepo    user.UserRepository
	productRepo product.ProductRepository
}

func NewUsecase(userRepo user.UserRepository, productRepo product.ProductRepository) Usecase {
	return &usecase{
		userRepo,
		productRepo,
	}
}

func (uc usecase) CreateUrl(body request.CreateUrl) (out models.Product, err error) {
	str := function.MakeUrl(body.DestinationUrl)
	for {
		check, err := uc.productRepo.FindUrl(str)
		if err != nil {
			break
		}
		if check {
			continue
		} else {
			err = nil
			break
		}
	}
	if err != nil {
		return
	} else {
		out, err = uc.productRepo.CreateUrlWithUser(body.DestinationUrl, str, body.UserID)
	}
	return
}

func (uc usecase) CreateUrlWithoutUser(body string) (str string, err error) {
	for {
		str = function.MakeUrl(body)
		check, err := uc.productRepo.FindUrl(str)
		fmt.Println(check, str, "======= check")
		if err != nil {
			break
		}
		if check {
			continue
		} else {
			err = nil
			break
		}
	}
	return str, err
}

func (uc usecase) UpdateUrl(body request.UpdateUrl) (out models.Product, err error) {
	out, err = uc.productRepo.UpdateUrl(body.DestinationUrl, body.ModifyUrl, body.Title, body.Source, body.Medium, body.Campaign, body.ID)
	fmt.Println(out, "out")
	if err != nil {
		return out, err
	}
	return out, nil
}

func (uc usecase) GetListProductByUserID(userID float64) (out []models.Product, err error) {
	out, err = uc.productRepo.GetListProductByUserID(userID)
	return
}

func (uc usecase) Register(body request.Register) (out models.User, err error) {
	_, err = mail.ParseAddress(body.Email)
	if (err != nil) {
		err = errors.New("your email doesn`t right")
		return
	}
	out, err = uc.userRepo.Register(body.Username, body.Email, body.Password)
	fmt.Println(err, "err ===")
	return
}

func (uc usecase) Login(body request.Login) (out string, err error) {
	secret_jwt := envy.Get("JWT_KEY", "jwt")
	login, err := uc.userRepo.Login(body.Find, body.Password)
	if err != nil {
		return "", err
	} else {
		claims := jwt.MapClaims{}
		claims["authorized"] = true
		claims["user_id"] = login.ID
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		return token.SignedString([]byte(secret_jwt))
	}
}

func (uc usecase) PatchGetClick(body int) (out models.Product, err error) {
	out, err = uc.productRepo.PatchGetClick(body)
	return
}

func (uc usecase) GetInfoUser(user_id float64) (out models.User, err error) {
	out, err = uc.userRepo.GetInfoUser(user_id)
	return
}

func (uc usecase) GetOneProduct(id int) (out models.Product, err error) {
	out, err = uc.productRepo.GetOneProduct(id)
	return
}
