package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID             int       `gorm:"column:id;primary_key;auto_increment;"`
	DestinationUrl string    `gorm:"column:destination_url;type:varchar(500);unique;"`
	ModifyUrl      string    `gorm:"column:modify_url;type:varchar(500);unique;"`
	Title          string    `gorm:"column:title;type:varchar(400);"`
	Click          int       `gorm:"column:click;type:int;"`
	UserID         float64       `gorm:"column:user_id;type:int;"`
	Source         string    `gorm:"column:source;type:varchar(400);"`
	Medium         string    `gorm:"column:medium;type:varchar(400);"`
	Campaign       string    `gorm:"column:campaign;type:varchar(400);"`
	CreatedAt      time.Time `gorm:"type:timestamp;column:created_at;"`
	UpdatedAt      time.Time `gorm:"column:updated_at;type:timestamp;"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	return
}

func (p *Product) BeforeUpdate(tx *gorm.DB) (err error) {
	p.UpdatedAt = time.Now()
	return
}
