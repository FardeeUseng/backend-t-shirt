package entities

import "github.com/gofiber/fiber/v2"

type ProductsUsecase interface {
	CreateProduct(req *CreateProductReq) (*Product, error)
	ProductList(c *fiber.Ctx) (*ProductListRes, error)
}

type ProductsRepository interface {
	CreateProduct(req *CreateProductReq) (*Product, error)
	ProductList(c *fiber.Ctx) (*ProductListRes, error)
}

type Product struct {
	Id        int    `json:"id" db:"id"`
	Gender    string `json:"gender" db:"gender"`
	Style     string `json:"style" db:"style"`
	Size      string `json:"size" db:"size"`
	Price     int    `json:"price" db:"price"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_datetime" db:"updated_datetime"`
}

type CreateProductReq struct {
	Gender string `json:"gender" db:"gender"`
	Style  string `json:"style" db:"style"`
	Size   string `json:"size" db:"size"`
	Price  int    `json:"price" db:"price"`
}

type ProductListRes struct {
	Page        int       `json:"page" db:"page"`
	ItemPerPage int       `json:"item_per_page" db:"item_per_page"`
	Total       int       `json:"total" db:"total"`
	Item        []Product `json:"item" db:"item"`
}
