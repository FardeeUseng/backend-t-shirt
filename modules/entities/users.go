package entities

import "github.com/gofiber/fiber/v2"

type UsersUsecase interface {
	UserList(c *fiber.Ctx) (*UserListRes, error)
}

type UsersRepository interface {
	UserList(c *fiber.Ctx) (*UserListRes, error)
}

type Reponse struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Result     interface{} `json:"result"`
}

type Users struct {
	Id        uint64 `json:"id" db:"id"`
	Username  string `json:"username" db:"username"`
	Gender    string `json:"gender" db:"gender"`
	Role      string `json:"role" db:"role"`
	CreatedAt string `json:"create_at" db:"create_at"`
}

type UserListRes struct {
	Page        int     `json:"page" db:"page"`
	ItemPerPage int     `json:"item_per_page" db:"item_per_page"`
	Total       int     `json:"total" db:"total"`
	Item        []Users `json:"item" db:"item"`
}
