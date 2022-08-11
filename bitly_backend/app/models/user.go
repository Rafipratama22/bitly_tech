package models

import (
	"time"

	"github.com/gomodul/envy"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID       int    `gorm:"column:id;primary_key;auto_increment;"`
	Username string `gorm:"column:username;type:varchar(400);"`
	Email    string `gorm:"column:email;type:varchar(400);unique;"`
	Password string `gorm:"column:password;type:varchar(500);unique"`
	CreatedAt      time.Time `gorm:"type:timestamp;column:created_at;"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;"`
	Products []Product
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	var cost = envy.GetInt("cost")
	if u.Password != "" {
		bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), cost)
		if err != nil {
			return err
		}
		u.Password = string(bytes)
	}
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return
}