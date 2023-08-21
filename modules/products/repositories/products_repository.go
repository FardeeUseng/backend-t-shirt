package repositories

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/FardeeUseng/backend-t-shirt/modules/entities"
	"github.com/gofiber/fiber/v2"
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

func (r *productsRepo) CreateProduct(req *entities.CreateProductReq) (*entities.Product, error) {
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

	product := new(entities.Product)
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

func (r *productsRepo) ProductList(c *fiber.Ctx) (*entities.ProductListRes, error) {
	page := 1
	itemPerPage := 4
	gender := c.Query("gender")
	size := c.Query("size")
	price := c.Query("price")
	style := c.Query("style")

	if c.Query("page") != "" {
		qPage, err := strconv.Atoi(c.Query("page"))
		if err == nil {
			page = qPage
		}
	}
	if c.Query("item_per_page") != "" {
		qItemPerPage, err := strconv.Atoi(c.Query("item_per_page"))
		if err == nil {
			itemPerPage = qItemPerPage
		}
	} else {
		page = 1
		itemPerPage = 100
	}

	offset := (page - 1) * itemPerPage

	query := `
		WITH products_total AS (
			SELECT
				COUNT(*) AS total
			FROM products p
			WHERE
				($3 = '' OR p.gender = CAST($3 AS gender))
					AND
				($4 = '' OR p.size = CAST($4 AS size))
					AND
				($5 = '' OR p.price = CAST($5 AS INTEGER))
					AND
				($6 = '' OR p.style LIKE '%' || $6 || '%')
		), products_item AS (
			SELECT
				*
			FROM products p
			WHERE
				($3 = '' OR p.gender = CAST($3 AS gender))
					AND
				($4 = '' OR p.size = CAST($4 AS size))
					AND
				($5 = '' OR p.price = CAST($5 AS INTEGER))
					AND
				($6 = '' OR p.style LIKE '%' || $6 || '%')
			ORDER BY p.id
			OFFSET $1 LIMIT $2
		)
		SELECT 
			(
				SELECT
					JSONB_AGG(x)
				FROM (
					SELECT
						id, gender, style, size, price, created_at
					FROM products_item
				) AS x
			) AS item,
			(
				SELECT total FROM products_total
			) AS total
	`
	var productsJSON []byte
	total := 0

	rows, err := r.Db.Queryx(query, offset, itemPerPage, gender, size, price, style)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	for rows.Next() {
		var productJSON []byte
		if err := rows.Scan(&productJSON, &total); err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		productsJSON = productJSON
	}

	var products []entities.Product

	if len(productsJSON) == 0 {
		return &entities.ProductListRes{
			Page:        page,
			ItemPerPage: itemPerPage,
			Total:       total,
			Item:        products,
		}, nil
	}

	if err := json.Unmarshal(productsJSON, &products); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return &entities.ProductListRes{
		Page:        page,
		ItemPerPage: itemPerPage,
		Total:       total,
		Item:        products,
	}, nil
}
