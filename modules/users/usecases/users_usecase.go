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

func (u *usersUse) CreateUser(req *entities.CreateUserReq) (*entities.Users, error) {
	user, err := u.UsersRepo.CreateUser(req)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *usersUse) UserList(c *fiber.Ctx) (*entities.UserListRes, error) {
	return nil, nil
}
