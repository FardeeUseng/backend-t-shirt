package usecases

import "github.com/FardeeUseng/backend-t-shirt/modules/entities"

type productsUse struct {
	ProductsRepo entities.ProductsRepository
}

func NewProductsUsecase(productsRepo entities.ProductsRepository) entities.ProductsUsecase {
	return &productsUse{
		ProductsRepo: productsRepo,
	}
}

func (u *productsUse) CreateProduct(req *entities.CreateProductReq) (*entities.CreateProductRes, error) {
	user, err := u.ProductsRepo.CreateProduct(req)
	if err != nil {
		return nil, err
	}
	return user, nil
}
