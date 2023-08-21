package servers

import (
	_ordersHttp "github.com/FardeeUseng/backend-t-shirt/modules/orders/controllers"
	_ordersRepository "github.com/FardeeUseng/backend-t-shirt/modules/orders/repositories"
	_ordersUsecase "github.com/FardeeUseng/backend-t-shirt/modules/orders/usecases"
	_productsHttp "github.com/FardeeUseng/backend-t-shirt/modules/products/controllers"
	_productsRepository "github.com/FardeeUseng/backend-t-shirt/modules/products/repositories"
	_productsUsecase "github.com/FardeeUseng/backend-t-shirt/modules/products/usecases"
	_usersHttp "github.com/FardeeUseng/backend-t-shirt/modules/users/controllers"
	_usersRepository "github.com/FardeeUseng/backend-t-shirt/modules/users/repositories"
	_usersUsecase "github.com/FardeeUseng/backend-t-shirt/modules/users/usecases"
	"github.com/gofiber/fiber/v2"
)

func (s *Server) MapHandler() error {
	v1 := s.App.Group("/v1")

	userGroup := v1.Group("/users")
	usersRepository := _usersRepository.NewUsersRepository(s.Db)
	usersUsecase := _usersUsecase.NewUserUsecase(usersRepository)
	_usersHttp.NewUsersController(userGroup, usersUsecase)

	productGroup := v1.Group("/products")
	productsRepository := _productsRepository.NewProductsRepository(s.Db)
	productsUsecase := _productsUsecase.NewProductsUsecase(productsRepository)
	_productsHttp.NewProductsController(productGroup, productsUsecase)

	orderGroup := productGroup.Group("/order")
	ordersRepository := _ordersRepository.NewOrdersRepository(s.Db)
	ordersUsecase := _ordersUsecase.NewOrdersUsecase(ordersRepository)
	_ordersHttp.NewOrdersController(orderGroup, ordersUsecase)

	s.App.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInsufficientStorage.Code,
			"message":     "error, end point not found",
			"result":      nil,
		})
	})
	return nil

}
