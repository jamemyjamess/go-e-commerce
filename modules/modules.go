package modules

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jamemyjamess/go-e-commerce/config"
)

type IModuleFactory interface {
	MonitorModule() IMonitorModule
}

type ModuleFactory struct {
	r      *fiber.Router
	cfgApp *config.IAppConfig
}

func NewModuleFactory(r *fiber.Router, cfgApp *config.IAppConfig) IModuleFactory {
	return &ModuleFactory{r: r, cfgApp: cfgApp}
}

// func (m *ModuleFactory) MonitorModule() {

// }
