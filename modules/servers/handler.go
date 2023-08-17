package servers

import (
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
