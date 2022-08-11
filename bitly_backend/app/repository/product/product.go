package product

import "bitly_backend/app/models"

type ProductRepository interface {
	CreateUrlWithUser(dest_url, modify_url string, userID float64) (product models.Product, errr error)
	UpdateUrl(dest_url, modify_url, title, source, medium, campaign string, productID int) (product models.Product, err error)
	GetListProductByUserID(userID float64) (product []models.Product, err error)
	PatchGetClick(id int) (product models.Product, err error)
	FindUrl(modify_url string) (check bool, errr error)
	GetOneProduct(id int) (product models.Product, err error)
}
