package entities

type ProductsUsecase interface {
	CreateProduct(req *CreateProductReq) (*CreateProductRes, error)
}

type ProductsRepository interface {
	CreateProduct(req *CreateProductReq) (*CreateProductRes, error)
}

type CreateProductRes struct {
	Id         int    `json:"id" db:"id"`
	Gender     string `json:"gender" db:"gender"`
	Style      string `json:"style" db:"style"`
	Size       string `json:"size" db:"size"`
	Price      int    `json:"price" db:"price"`
	Created_at string `json:"created_at" db:"created_at"`
}

type CreateProductReq struct {
	Gender string `json:"gender" db:"gender"`
	Style  string `json:"style" db:"style"`
	Size   string `json:"size" db:"size"`
	Price  int    `json:"price" db:"price"`
}
