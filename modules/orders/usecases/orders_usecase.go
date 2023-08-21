package usecases

import "github.com/FardeeUseng/backend-t-shirt/modules/entities"

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
