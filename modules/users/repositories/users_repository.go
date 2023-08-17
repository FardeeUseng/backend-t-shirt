package repositories

import (
	"fmt"

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
	return nil, nil
}
