package controllers

import (
	"strconv"

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
	r.Get("/:id", controller.UserInfo)
}

func (h *usersController) CreateUser(c *fiber.Ctx) error {
	req := new(entities.CreateUserReq)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(&entities.Response{
			Status:     fiber.ErrBadRequest.Message,
			StatusCode: fiber.ErrBadRequest.Code,
			Message:    err.Error(),
			Result:     nil,
		})
	}

	res, err := h.UserUse.CreateUser(req)
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(&entities.Response{
			Status:     fiber.ErrInternalServerError.Message,
			StatusCode: fiber.ErrInternalServerError.Code,
			Message:    err.Error(),
			Result:     nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(&entities.Response{
		Status:     "OK",
		StatusCode: fiber.StatusOK,
		Message:    "",
		Result: fiber.Map{
			"data": res,
		},
	})
}

func (h *usersController) UserInfo(c *fiber.Ctx) error {
	id := c.Params("id")
	userId, qErr := strconv.Atoi(id)
	if qErr != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(&entities.Response{
			Status:     fiber.ErrBadRequest.Message,
			StatusCode: fiber.ErrBadRequest.Code,
			Message:    qErr.Error(),
			Result:     nil,
		})
	}

	res, err := h.UserUse.UserInfo(userId)
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(&entities.Response{
			Status:     fiber.ErrInternalServerError.Message,
			StatusCode: fiber.ErrInternalServerError.Code,
			Message:    err.Error(),
			Result:     nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(&entities.Response{
		Status:     "OK",
		StatusCode: fiber.StatusOK,
		Message:    "",
		Result: fiber.Map{
			"data": res,
		},
	})
}

func (h *usersController) UserList(c *fiber.Ctx) error {
	res, err := h.UserUse.UserList(c)
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(&entities.Response{
			Status:     fiber.ErrInternalServerError.Message,
			StatusCode: fiber.ErrInternalServerError.Code,
			Message:    err.Error(),
			Result:     nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(&entities.Response{
		Status:     "OK",
		StatusCode: fiber.StatusOK,
		Message:    "",
		Result: fiber.Map{
			"data": res,
		},
	})
}
