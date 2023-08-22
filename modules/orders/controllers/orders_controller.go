package controllers

import (
	"strconv"

	"github.com/FardeeUseng/backend-t-shirt/modules/entities"
	"github.com/gofiber/fiber/v2"
)

type ordersController struct {
	OrderUse entities.OrdersUsecase
}

func NewOrdersController(r fiber.Router, ordersUse entities.OrdersUsecase) {
	controller := &ordersController{
		OrderUse: ordersUse,
	}
	r.Post("/", controller.CreateOrder)
	r.Post("/shipping", controller.CreateShipping)
	r.Get("/:id", controller.OrderList)
}

func (h *ordersController) CreateOrder(c *fiber.Ctx) error {
	req := new(entities.CreateOrderReq)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(&entities.Response{
			Status:     fiber.ErrBadRequest.Message,
			StatusCode: fiber.ErrBadRequest.Code,
			Message:    err.Error(),
			Result:     nil,
		})
	}

	if len(req.ProductId) == 0 {
		return c.Status(fiber.ErrBadRequest.Code).JSON(&entities.Response{
			Status:     fiber.ErrBadRequest.Message,
			StatusCode: fiber.ErrBadRequest.Code,
			Message:    "product id is require",
			Result:     nil,
		})
	}

	res, err := h.OrderUse.CreateOrder(req)
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(&entities.Response{
			Status:     fiber.ErrInternalServerError.Message,
			StatusCode: fiber.ErrInternalServerError.Code,
			Message:    err.Error(),
			Result:     nil,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(&entities.Response{
		Status:     "OK",
		StatusCode: fiber.StatusCreated,
		Message:    "",
		Result: fiber.Map{
			"data": res,
		},
	})
}

func (h *ordersController) CreateShipping(c *fiber.Ctx) error {
	req := new(entities.ShippingReq)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(&entities.Response{
			Status:     fiber.ErrBadRequest.Message,
			StatusCode: fiber.ErrBadRequest.Code,
			Message:    err.Error(),
			Result:     nil,
		})
	}

	res, err := h.OrderUse.CreateShipping(req)
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(&entities.Response{
			Status:     fiber.ErrInternalServerError.Message,
			StatusCode: fiber.ErrInternalServerError.Code,
			Message:    err.Error(),
			Result:     nil,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(&entities.Response{
		Status:     "OK",
		StatusCode: fiber.StatusCreated,
		Message:    "",
		Result: fiber.Map{
			"data": res,
		},
	})
}

func (h *ordersController) OrderList(c *fiber.Ctx) error {
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

	res, err := h.OrderUse.OrderList(userId, c)
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
		Result:     res,
	})
}
