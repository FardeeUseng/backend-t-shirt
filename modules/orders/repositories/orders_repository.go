package repositories

import (
	"github.com/FardeeUseng/backend-t-shirt/modules/entities"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type ordersRepo struct {
	Db *sqlx.DB
}

func NewOrdersRepository(db *sqlx.DB) entities.OrdersRepository {
	return &ordersRepo{
		Db: db,
	}
}

func (h *ordersRepo) CreateOrder(req *entities.CreateOrderReq) (*entities.CreateOrderRes, error) {
	// Start transaction
	tx, err := h.Db.Begin()
	if err != nil {
		log.Errorf("%v", err.Error())
		return nil, err
	}
	defer tx.Rollback() // Rollback if any error occurs

	// Query products by IDs
	productQuery := `
		SELECT *
		FROM products p
		WHERE p.id = ANY($1)
	`

	pRows, pErr := h.Db.Query(productQuery, pq.Array(req.ProductId))
	if pErr != nil {
		tx.Rollback()
		log.Errorf("%v", pErr.Error())
		return nil, pErr
	}
	defer pRows.Close()

	// Insert order
	orderQuery := `
		INSERT INTO "orders"(
			"user_id",
			"status"
		)
		VALUES ($1, $2)
		RETURNING "id", "user_id", "status"
	`

	var order entities.Order
	oRow := tx.QueryRow(orderQuery, req.UserId, "placed_order")
	if scanErr := oRow.Scan(&order.Id, &order.UserId, &order.Status); scanErr != nil {
		tx.Rollback()
		log.Errorf("Error scanning order result: %v", scanErr)
		return nil, scanErr
	}

	// Insert order product
	orderProductQuery := `
		INSERT INTO "order_product"(
			"order_id",
			"product_id"
		)
		VALUES ($1, $2)
	`

	for pRows.Next() {
		var product entities.Product

		if err := pRows.Scan(&product.Id, &product.Gender, &product.Style, &product.Size, &product.Price, &product.Created_at, &product.Updated_at); err != nil {
			tx.Rollback()
			log.Errorf("%v", err.Error())
			return nil, err
		}

		_, err := tx.Exec(orderProductQuery, order.Id, product.Id)
		if err != nil {
			tx.Rollback()
			log.Errorf("%v", err.Error())
			return nil, err
		}
	}

	// Commit
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		log.Errorf("%v", err.Error())
		return nil, err
	}

	// Get product list
	productsQuery := `
		WITH "order_id" AS (
			SELECT "id"
			FROM "orders" o
			WHERE o.user_id = $1
		),
		"product_id" AS (
				SELECT "product_id"
				FROM "order_product" op
				WHERE op.order_id IN (SELECT "id" FROM "order_id")
		)
		SELECT *
		FROM "products" p
		WHERE p.id IN (SELECT "product_id" FROM "product_id");
	`

	var products []entities.Product
	prodsRes, prodsErr := h.Db.Query(productsQuery, req.UserId)
	if prodsErr != nil {
		tx.Rollback()
		log.Errorf("%v", prodsErr.Error())
		return nil, prodsErr
	}
	defer prodsRes.Close()

	for prodsRes.Next() {
		var product entities.Product
		if err := prodsRes.Scan(&product.Id, &product.Gender, &product.Style, &product.Size, &product.Price, &product.Created_at, &product.Updated_at); err != nil {
			tx.Rollback()
			log.Errorf("%v", err.Error())
			return nil, err
		}
		products = append(products, product)
	}

	// Return response
	return &entities.CreateOrderRes{
		UserId:  req.UserId,
		OrderId: order.Id,
		Status:  order.Status,
		Product: products,
	}, nil
}
