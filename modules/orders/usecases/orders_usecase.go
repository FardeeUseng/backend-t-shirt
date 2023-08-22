package usecases

import (
	"github.com/FardeeUseng/backend-t-shirt/modules/entities"
	"github.com/gofiber/fiber/v2"
)

type ordersUse struct {
	OrdersRepo entities.OrdersRepository
}

func NewOrdersUsecase(ordersRepo entities.OrdersRepository) entities.OrdersUsecase {
	return &ordersUse{
		OrdersRepo: ordersRepo,
	}
}

func (u *ordersUse) CreateOrder(req *entities.CreateOrderReq) (*entities.CreateOrderRes, error) {
	res, err := u.OrdersRepo.CreateOrder(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u *ordersUse) CreateShipping(req *entities.ShippingReq) (*entities.ShippingRes, error) {
	res, err := u.OrdersRepo.CreateShipping(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u *ordersUse) OrderList(userId int, c *fiber.Ctx) (*entities.OrderListRes, error) {
	res, err := u.OrdersRepo.OrderList(userId, c)
	if err != nil {
		return nil, err
	}
	return res, nil
}
