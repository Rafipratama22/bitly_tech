package product

import (
	"bitly_backend/app/models"
	// "fmt"

	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		db,
	}
}

func (repo productRepository) CreateUrlWithUser(dest_url, modify_url string, userID float64) (product models.Product, errr error) {
	product.DestinationUrl = dest_url
	product.ModifyUrl = modify_url
	product.UserID = userID
	if err := repo.db.Model(&product).Create(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}

func (repo productRepository) FindUrl(modify_url string) (check bool, errr error) {
	var product models.Product
	product.ModifyUrl = modify_url
	if err := repo.db.Model(&product).Where("modify_url = ?", modify_url).Find(&product).Error; err != nil {
		return false, err
	}
	if product.ID > 0 {
		return true, nil
	} else {
		return false , nil
	}
}

func (repo productRepository) UpdateUrl(dest_url, modify_url, title, source, medium, campaign string, productID int) (product models.Product, err error){
	product.ID = productID
	if err = repo.db.First(&product).Error; err != nil {
		return
	}
	product.DestinationUrl = dest_url
	product.ModifyUrl = modify_url
	product.Title = title
	product.Campaign = campaign
	product.Source = source
	product.Medium = medium
	if err = repo.db.Save(&product).Error; err != nil {
		return
	}
	return product, nil
}

func (repo productRepository) GetListProductByUserID(userID float64) (product []models.Product, err error) {
	if err = repo.db.Model(&product).Where("user_id = ?", userID).Find(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}

func (repo productRepository) GetOneProduct(id int) (product models.Product, err error) {
	if err = repo.db.Model(&product).Where("id = ?", id).First(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}

func (repo productRepository) PatchGetClick(id int) (product models.Product, err error) {
	if err = repo.db.Model(&product).Where("id = ?", id).Find(&product).Error; err != nil {
		return product, err
	} else {
		product.Click += 1
		if err = repo.db.Model(&product).Where("id = ?", id).Update("click", product.Click).Error; err != nil {
			return product, err
		}
		
	}
	return product, nil
}