package controllers

import (
	"github.com/FardeeUseng/backend-t-shirt/modules/entities"
	"github.com/gofiber/fiber/v2"
)

type productsController struct {
	ProductUse entities.ProductsUsecase
}

func NewProductsController(r fiber.Router, productsUse entities.ProductsUsecase) {
	controller := &productsController{
		ProductUse: productsUse,
	}
	r.Post("/", controller.CreateProduct)
}

func (h *productsController) CreateProduct(c *fiber.Ctx) error {
	req := new(entities.CreateProductReq)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(&entities.Response{
			Status:     fiber.ErrBadRequest.Message,
			StatusCode: fiber.ErrBadRequest.Code,
			Message:    err.Error(),
			Result:     nil,
		})
	}

	res, err := h.ProductUse.CreateProduct(req)
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
