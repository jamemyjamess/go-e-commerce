package modules

import (
	"github.com/jamemyjamess/go-e-commerce/modules/monitor/monitorHandler"
)

type IMonitorModule interface {
	InitRoutes()
}

type monitorModule struct {
	*ModuleFactory
	// cfgApp *config.IAppConfig
	handler monitorHandler.IMonitorHandler
}

func (m *ModuleFactory) MonitorModule() IMonitorModule {
	appConfig := (*m.cfg).App()
	monitorHandler := monitorHandler.NewMonitorHandler(&appConfig)
	return &monitorModule{
		ModuleFactory: m,
		handler:       monitorHandler,
	}
}

func (m *monitorModule) InitRoutes() {
	(*m.r).Get("/", m.handler.CheckServer)

}
