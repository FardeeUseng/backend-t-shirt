package entities

import "github.com/gofiber/fiber/v2"

type OrdersUsecase interface {
	CreateOrder(req *CreateOrderReq) (*CreateOrderRes, error)
	CreateShipping(req *ShippingReq) (*ShippingRes, error)
	OrderList(userId int, c *fiber.Ctx) (*OrderListRes, error)
}

type OrdersRepository interface {
	CreateOrder(req *CreateOrderReq) (*CreateOrderRes, error)
	CreateShipping(req *ShippingReq) (*ShippingRes, error)
	OrderList(userId int, c *fiber.Ctx) (*OrderListRes, error)
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

type ShippingReq struct {
	OrderId     int    `json:"order_id" db:"order_id"`
	Address     string `json:"address" db:"address"`
	Subdistrict string `json:"subdistrict" db:"subdistrict"`
	District    string `json:"district" db:"district"`
	Province    string `json:"province" db:"province"`
	ZipCode     string `json:"zip_code" db:"zip_code"`
}

type ShippingRes struct {
	Id              int    `json:"id" db:"id"`
	OrderId         int    `json:"order_id" db:"order_id"`
	Address         string `json:"address" db:"address"`
	Subdistrict     string `json:"subdistrict" db:"subdistrict"`
	District        string `json:"district" db:"district"`
	Province        string `json:"province" db:"province"`
	ZipCode         string `json:"zip_code" db:"zip_code"`
	CreatedAt       string `json:"created_at" db:"created_at"`
	UpdatedDatetime string `json:"updated_datetime" db:"updated_datetime"`
}

type OrderList struct {
	Status   string    `json:"status"`
	UserId   int       `json:"user_id"`
	OrderId  int       `json:"order_id"`
	Products []Product `json:"product"`
}

type OrderListRes struct {
	Page        int         `json:"page" db:"page"`
	ItemPerPage int         `json:"item_per_page"`
	Item        []OrderList `json:"item" db:"item"`
}
