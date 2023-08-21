package entities

type OrdersUsecase interface {
	CreateOrder(req *CreateOrderReq) (*CreateOrderRes, error)
}

type OrdersRepository interface {
	CreateOrder(req *CreateOrderReq) (*CreateOrderRes, error)
}

type CreateOrderReq struct {
	UserId    int   `json:"user_id" `
	ProductId []int `json:"product_id"`
}

type CreateOrderRes struct {
	UserId  int       `json:"user_id" `
	OrderId int       `json:"order_id" `
	Status  string    `json:"status"`
	Product []Product `json:"product"`
}

type Order struct {
	Id     int    `json:"id" db:"id"`
	UserId int    `json:"user_id" db:"user_id"`
	Status string `json:"status" db:"status"`
}
