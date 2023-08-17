package usecases

import (
	"github.com/FardeeUseng/backend-t-shirt/modules/entities"
	"github.com/gofiber/fiber/v2"
)

type usersUse struct {
	UsersRepo entities.UsersRepository
}

func NewUserUsecase(usersRepo entities.UsersRepository) entities.UsersUsecase {
	return &usersUse{
		UsersRepo: usersRepo,
	}
}

func (u *usersUse) UserList(c *fiber.Ctx) (*entities.UserListRes, error) {
	return nil, nil
}
