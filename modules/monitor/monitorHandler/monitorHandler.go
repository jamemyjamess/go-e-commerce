package monitorHandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jamemyjamess/go-e-commerce/config"
	"github.com/jamemyjamess/go-e-commerce/modules/monitor/entities"
	"github.com/jamemyjamess/go-e-commerce/pkg/response"
)

type IMonitorHandler interface {
	CheckServer(c *fiber.Ctx) error
}

type MonitorHandler struct {
	cfgApp *config.IAppConfig
}

func NewMonitorHandler(r *fiber.Router, cfgApp *config.IAppConfig) IMonitorHandler {
	return &MonitorHandler{cfgApp: cfgApp}
}

func (m *MonitorHandler) CheckServer(c *fiber.Ctx) error {
	res := entities.ResCheckServer{
		Name:    (*m.cfgApp).Name(),
		Version: (*m.cfgApp).Version(),
	}
	return response.NewResponse(c).Success(fiber.StatusOK, res).Res()
}
