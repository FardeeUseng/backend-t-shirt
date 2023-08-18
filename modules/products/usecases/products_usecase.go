package usecases

import (
	"github.com/FardeeUseng/backend-t-shirt/modules/entities"
	"github.com/gofiber/fiber/v2"
)

type productsUse struct {
	ProductsRepo entities.ProductsRepository
}

func NewProductsUsecase(productsRepo entities.ProductsRepository) entities.ProductsUsecase {
	return &productsUse{
		ProductsRepo: productsRepo,
	}
}

func (u *productsUse) CreateProduct(req *entities.CreateProductReq) (*entities.Product, error) {
	product, err := u.ProductsRepo.CreateProduct(req)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (u *productsUse) ProductList(c *fiber.Ctx) (*entities.ProductListRes, error) {
	product, err := u.ProductsRepo.ProductList(c)
	if err != nil {
		return nil, err
	}
	return product, nil
}
