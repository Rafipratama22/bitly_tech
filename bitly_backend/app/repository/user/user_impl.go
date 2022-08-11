package user

import (
	"bitly_backend/app/models"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db,
	}	
}

func (repo userRepository) Register(username, email, password string) (user models.User, err error) {
	user.Username = username
	user.Email = email
	user.Password = password
	if err = repo.db.Model(&user).Create(&user).Select("id", "username", "email", "created_at", "updated_at").Error; err != nil {
		return user, err
	}
	return user, nil
}

func (repo userRepository) Login(find, password string) (user models.User, err error) {
	if err = repo.db.Model(&user).Where("username = ?", find).First(&user).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		if err = repo.db.Model(&user).Where("email = ?", find).First(&user).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			return user, errors.New("data not found on database")
		}
	}

	fmt.Println(user.Password)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, errors.New("unauthorized")
	}
	
	return
}

func (repo userRepository) GetAllProductByUserID(user_id float64) (user models.User, err error) {
	if err = repo.db.Model(&user).Where("id = ?", user_id).Preload("Products").Find(&user).Error; err != nil {
		return
	}
	return
}

func (repo userRepository) GetInfoUser(user_id float64) (user models.User, err error) {
	if err = repo.db.Model(&user).Where("id = ?", user_id).First(&user).Error; err != nil {
		return
	}
	return
}