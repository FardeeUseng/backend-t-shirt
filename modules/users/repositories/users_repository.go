package repositories

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/FardeeUseng/backend-t-shirt/modules/entities"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type usersRepo struct {
	Db *sqlx.DB
}

func NewUsersRepository(db *sqlx.DB) entities.UsersRepository {
	return &usersRepo{
		Db: db,
	}
}

func (r *usersRepo) CreateUser(req *entities.CreateUserReq) (*entities.Users, error) {
	query := `
		INSERT INTO "users"(
			"username",
			"gender",
			"role"
		)
		VALUES($1, $2, $3)
		RETURNING "id", "username", "gender", "role", "created_at"
	`

	user := new(entities.Users)

	rows, err := r.Db.Queryx(query, req.Username, req.Gender, req.Role)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	for rows.Next() {
		if err := rows.StructScan(user); err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
	}
	return user, nil
}

func (r *usersRepo) UserList(c *fiber.Ctx) (*entities.UserListRes, error) {
	page := 1
	itemPerPage := 4
	gender := c.Query("gender")
	role := c.Query("role")
	searhName := c.Query("search")

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
		WITH users_total AS (
			SELECT
				COUNT(*) AS total
			FROM users u
			WHERE
				($3 = '' OR u.gender = CAST($3 AS gender))
				AND
				($4 = '' OR u.role = $4)
				AND
				($5 = '' OR u.username LIKE '%' || $5 || '%')
		), users_item AS (
			SELECT
				*
			FROM users u
			WHERE
				($3 = '' OR u.gender = CAST($3 AS gender))
				AND
				($4 = '' OR u.role = $4)
				AND
				($5 = '' OR u.username LIKE '%' || $5 || '%')
			ORDER BY u.id
			OFFSET $1 LIMIT $2
		)
		SELECT 
			(
				SELECT
					JSONB_AGG(x)
				FROM (
					SELECT
						id, username, gender, role, created_at
					FROM users_item
				) AS x
			) AS item,
			(
				SELECT total FROM users_total
			) AS total
	`

	var usersJSON []byte
	total := 0

	rows, err := r.Db.Queryx(query, offset, itemPerPage, gender, role, searhName)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	for rows.Next() {
		var userJSON []byte
		if err := rows.Scan(&userJSON, &total); err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		usersJSON = userJSON
	}

	var users []entities.Users

	if len(usersJSON) == 0 {
		return &entities.UserListRes{
			Page:        page,
			ItemPerPage: itemPerPage,
			Total:       total,
			Item:        users,
		}, nil
	}

	if err := json.Unmarshal(usersJSON, &users); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return &entities.UserListRes{
		Page:        page,
		ItemPerPage: itemPerPage,
		Total:       total,
		Item:        users,
	}, nil
}

func (r *usersRepo) UserInfo(id int) (*entities.Users, error) {

	query := `
		SELECT
			"id", "username", "gender", "role", "created_at"
		FROM "users"
		WHERE "id" = $1
	`
	user := new(entities.Users)

	err := r.Db.QueryRow(query, id).Scan(&user.Id, &user.Username, &user.Gender, &user.Role, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		fmt.Println(err.Error())
		return nil, err
	}
	return user, nil
}
