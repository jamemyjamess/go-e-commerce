package middlewaresHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jamemyjamess/go-e-commerce/config"
	middlewaresEntities "github.com/jamemyjamess/go-e-commerce/middlewares/entities"
	"github.com/jamemyjamess/go-e-commerce/middlewares/middlewaresUsecases"
	"github.com/jamemyjamess/go-e-commerce/pkg/response"
)

type IMiddlewareHandlers interface {
	CORS() fiber.Handler
	RouterNorFoundInfo() fiber.Handler
	Logger() fiber.Handler
}

type middlewaresHandlers struct {
	cfg                config.IConfig
	middlewaresUsecase middlewaresUsecases.IMiddlewareUsecase
}

func MiddlewaresHandlers(cfg config.IConfig, middlewaresUsecase middlewaresUsecases.IMiddlewareUsecase) IMiddlewareHandlers {
	return &middlewaresHandlers{
		cfg:                cfg,
		middlewaresUsecase: middlewaresUsecase,
	}
}

func (h *middlewaresHandlers) CORS() fiber.Handler {
	return cors.New(cors.Config{
		Next:             cors.ConfigDefault.Next,
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	})
}

func (h *middlewaresHandlers) RouterNorFoundInfo() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return response.NewResponse(c).Error(
			fiber.ErrNotFound.Code,
			string(middlewaresEntities.HandleRouterNorFoundInfoErr),
			"router not fount",
		).Res()
	}
}

func (h *middlewaresHandlers) Logger() fiber.Handler {
	return logger.New(logger.Config{
		Format:     "${time} [${ip} ${status} - ${method} ${path}]",
		TimeFormat: "01/02/2006",
		TimeZone:   "Bangkok/Asia",
	})
}
