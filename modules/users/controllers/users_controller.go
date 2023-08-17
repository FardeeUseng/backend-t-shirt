package controllers

import (
	"github.com/FardeeUseng/backend-t-shirt/modules/entities"
	"github.com/gofiber/fiber/v2"
)

type usersController struct {
	UserUse entities.UsersUsecase
}

func NewUsersController(r fiber.Router, usersUse entities.UsersUsecase) {
	controller := &usersController{
		UserUse: usersUse,
	}
	r.Get("/", controller.UserList)
}

func (h *usersController) UserList(c *fiber.Ctx) error {
	return nil
}
