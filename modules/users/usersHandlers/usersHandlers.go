package usersHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jamemyjamess/go-e-commerce/modules/users/usersUsecase"
	"github.com/jamemyjamess/go-e-commerce/pkg/response"
)

type IUserHandler interface {
	SignIn(c *fiber.Ctx) error
}

type userHandler struct {
	usersUsecase *usersUsecase.IUsersUsecase
}

func NewUsersHandler(u *usersUsecase.IUsersUsecase) IUserHandler {
	return &userHandler{
		usersUsecase: u,
	}
}

func (u *userHandler) SignIn(c *fiber.Ctx) error {

	return response.NewResponse(c).Success(fiber.StatusOK, nil).Res()
}
