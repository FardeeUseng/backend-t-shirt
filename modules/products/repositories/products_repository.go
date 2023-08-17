package repositories

import (
	"fmt"

	"github.com/FardeeUseng/backend-t-shirt/modules/entities"
	"github.com/jmoiron/sqlx"
)

type productsRepo struct {
	Db *sqlx.DB
}

func NewProductsRepository(db *sqlx.DB) entities.ProductsRepository {
	return &productsRepo{
		Db: db,
	}
}

func (r *productsRepo) CreateProduct(req *entities.CreateProductReq) (*entities.CreateProductRes, error) {
	query := `
		INSERT INTO "products"(
			"gender",
			"style",
			"size",
			"price"
		)
		VALUES ($1, $2, $3, $4)
		RETURNING "id", "gender", "style", "size", "price", "created_at"
	`

	product := new(entities.CreateProductRes)
	rows, err := r.Db.Queryx(query, req.Gender, req.Style, req.Size, req.Price)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	for rows.Next() {
		if err := rows.StructScan(product); err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
	}

	return product, nil
}
