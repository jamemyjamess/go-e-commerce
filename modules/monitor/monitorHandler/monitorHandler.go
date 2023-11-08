package monitorHandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jamemyjamess/go-e-commerce/config"
	"github.com/jamemyjamess/go-e-commerce/modules/monitor/entities"
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
	return c.Status(fiber.StatusOK).JSON(res)
}
