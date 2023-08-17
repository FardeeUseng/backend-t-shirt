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
	r.Post("/", controller.CreateUser)
	r.Get("/", controller.UserList)
}

func (h *usersController) CreateUser(c *fiber.Ctx) error {
	req := new(entities.CreateUserReq)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(&entities.Reponse{
			Status:     fiber.ErrBadRequest.Message,
			StatusCode: fiber.ErrBadRequest.Code,
			Message:    err.Error(),
			Result:     nil,
		})
	}

	res, err := h.UserUse.CreateUser(req)
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(&entities.Reponse{
			Status:     fiber.ErrInternalServerError.Message,
			StatusCode: fiber.ErrInternalServerError.Code,
			Message:    err.Error(),
			Result:     nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(&entities.Reponse{
		Status:     "OK",
		StatusCode: fiber.StatusOK,
		Message:    "",
		Result: fiber.Map{
			"data": res,
		},
	})
}

func (h *usersController) UserList(c *fiber.Ctx) error {
	return nil
}
