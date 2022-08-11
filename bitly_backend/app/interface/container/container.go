package container

import (
	"bitly_backend/app/adapter/database"
	"bitly_backend/app/interface/middleware"
	"bitly_backend/app/repository/product"
	"bitly_backend/app/repository/user"
	"bitly_backend/app/usecase"
)

type Container struct {
	Usecase usecase.Usecase
	Middleware middleware.Middlewared
}

func SetupContainer() Container {
	mysql := database.SetupMySQL()
	
	userRepo := user.NewUserRepository(mysql.DB)
	productRepo := product.NewProductRepository(mysql.DB)
	middleware := middleware.NewMiddleware(mysql.DB)

	usecases := usecase.NewUsecase(userRepo, productRepo)

	return Container{
		usecases,
		middleware,
	}
}