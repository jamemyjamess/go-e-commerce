package usersHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jamemyjamess/go-e-commerce/pkg/response"
)

type IUserHandler interface {
	SignIn(c *fiber.Ctx) error
}

func NewUsersHandler() IUserHandler {
	return &userHandler{}
}

type userHandler struct {
}

func (u *userHandler) SignIn(c *fiber.Ctx) error {

	return response.NewResponse(c).Success(fiber.StatusOK, nil).Res()
}
