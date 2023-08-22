package repositories

import (
	"fmt"
	"strconv"
	"time"

	"github.com/FardeeUseng/backend-t-shirt/modules/entities"
	"github.com/gofiber/fiber/v2"
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

		if err := pRows.Scan(&product.Id, &product.Gender, &product.Style, &product.Size, &product.Price, &product.CreatedAt, &product.UpdatedAt); err != nil {
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
		if err := prodsRes.Scan(&product.Id, &product.Gender, &product.Style, &product.Size, &product.Price, &product.CreatedAt, &product.UpdatedAt); err != nil {
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

func (h *ordersRepo) CreateShipping(req *entities.ShippingReq) (*entities.ShippingRes, error) {

	tx, err := h.Db.Begin()
	if err != nil {
		log.Errorf("%v", err.Error())
		return nil, err
	}
	defer tx.Rollback()

	orderQuery := `
		SELECT "id", "user_id", "status"
		FROM "orders" o
		WHERE o.id = $1
	`

	uRows, uErr := h.Db.Query(orderQuery, req.OrderId)
	if uErr != nil {
		tx.Rollback()
		log.Errorf("%v", uErr.Error())
		return nil, uErr
	}
	defer uRows.Close()

	var orders []entities.Order
	for uRows.Next() {
		var order entities.Order
		if err := uRows.Scan(&order.Id, &order.UserId, &order.Status); err != nil {
			tx.Rollback()
			log.Errorf("%v", err)
			return nil, err
		}
		orders = append(orders, order)
	}

	if len(orders) == 0 {
		tx.Rollback()
		return nil, fmt.Errorf("error, Can not find order id")
	}

	shippingQuery := `
		INSERT INTO "shippings"(
			"order_id",
			"address",
			"subdistrict",
			"district",
			"province",
			"zip_code"
		)
		VALUES($1, $2, $3, $4, $5, $6)
		RETURNING "id", "order_id", "address", "subdistrict", "district", "province", "zip_code", "created_at","updated_datetime"
	`

	var shipping entities.ShippingRes
	sRow := tx.QueryRow(shippingQuery, req.OrderId, req.Address, req.Subdistrict, req.District, req.Province, req.ZipCode)
	if sErr := sRow.Scan(&shipping.Id, &shipping.OrderId, &shipping.Address, &shipping.Subdistrict, &shipping.District, &shipping.Province, &shipping.ZipCode, &shipping.CreatedAt, &shipping.UpdatedDatetime); sErr != nil {
		tx.Rollback()
		log.Errorf("%v", sErr)
		return nil, sErr
	}

	// Commit
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		log.Errorf("%v", err.Error())
		return nil, err
	}

	return &entities.ShippingRes{
		Id:              shipping.Id,
		OrderId:         shipping.OrderId,
		Address:         shipping.Address,
		Subdistrict:     shipping.Subdistrict,
		District:        shipping.District,
		Province:        shipping.Province,
		ZipCode:         shipping.ZipCode,
		CreatedAt:       shipping.CreatedAt,
		UpdatedDatetime: shipping.UpdatedDatetime,
	}, nil
}

func (h *ordersRepo) OrderList(userId int, c *fiber.Ctx) (*entities.OrderListRes, error) {

	page := 1
	itemPerPage := 10
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if startDate == "" || endDate == "" {
		startDate = "2022-12-01"
		endDate = time.Now().Format("2006-01-02")
	}

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

	tx, err := h.Db.Begin()
	if err != nil {
		log.Errorf("%v", err.Error())
		return nil, err
	}
	defer tx.Rollback()

	userQuery := `
		SELECT
			u.id,
			u.role
		FROM "users" u
		WHERE u.role = 'admin' and u.id = $1
	`

	uRows, uErr := h.Db.Query(userQuery, userId)
	if uErr != nil {
		log.Errorf("Error executing query: %v", uErr)
		return nil, uErr
	}
	defer uRows.Close()

	var users []entities.Users
	for uRows.Next() {
		var user entities.Users
		if err := uRows.Scan(&user.Id, &user.Role); err != nil {
			tx.Rollback()
			log.Errorf("%v", err.Error())
			return nil, err
		}
		users = append(users, user)
	}

	if len(users) == 0 {
		return nil, fmt.Errorf("error, you don't have permission")
	}

	orderQuery := `
		SELECT 
			o.status,
			o.user_id,
			o.id AS order_id,
			p.*
		FROM "orders" o
		JOIN "order_product" op ON o.id = op.order_id
		JOIN "products" p ON op.product_id = p.id
		WHERE 
			($1 = '' OR o.status = $1)
			AND
			(o.created_at BETWEEN COALESCE($4, '2023-08-01'::timestamp) AND COALESCE($5, NOW()))
		OFFSET $2 LIMIT $3
	`

	oRows, oErr := h.Db.Query(orderQuery, c.Query("status"), offset, itemPerPage, startDate, endDate)
	if oErr != nil {
		log.Errorf("Error executing query: %v", oErr)
		return nil, oErr
	}
	defer oRows.Close()

	orderMap := make(map[int]entities.OrderList)
	for oRows.Next() {
		var status string
		var userId, orderId int
		var product entities.Product

		if err := oRows.Scan(&status, &userId, &orderId, &product.Id, &product.Gender, &product.Style, &product.Size, &product.Price, &product.CreatedAt, &product.UpdatedAt); err != nil {
			log.Errorf("Error scanning row: %v", err)
			return nil, err
		}

		if existingOrder, ok := orderMap[orderId]; ok {
			existingOrder.Products = append(existingOrder.Products, product)
			orderMap[orderId] = existingOrder
		} else {
			orderMap[orderId] = entities.OrderList{
				Status:   status,
				UserId:   userId,
				OrderId:  orderId,
				Products: []entities.Product{product},
			}
		}
	}

	// Commit
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		log.Errorf("%v", err.Error())
		return nil, err
	}

	var orders []entities.OrderList
	for _, order := range orderMap {
		orders = append(orders, order)
	}

	return &entities.OrderListRes{
		Page:        page,
		ItemPerPage: itemPerPage,
		Item:        orders,
	}, nil
}
