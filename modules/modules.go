package modules

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jamemyjamess/go-e-commerce/config"
	"github.com/jamemyjamess/go-e-commerce/pkg/errorTranslator"
	"github.com/jmoiron/sqlx"
)

type IModuleFactory interface {
	MonitorModule() IMonitorModule
	UserModule() IUserModule
}

type ModuleFactory struct {
	// s        *servers.IServer
	r          *fiber.Router
	cfg        *config.IConfig
	postgresDb *sqlx.DB
	errTrans   *errorTranslator.IErrorTranslator
}

func NewModuleFactory(r *fiber.Router, cfg *config.IConfig, postgresDb *sqlx.DB) IModuleFactory {
	errTrans := errorTranslator.NewErrorTranslator()
	errTrans.InitDefaultTranslator()
	return &ModuleFactory{r: r, cfg: cfg, postgresDb: postgresDb, errTrans: &errTrans}
}

// func (m *ModuleFactory) MonitorModule() {

// }

// func (m *ModuleFactory) InitErrorTranslator() {
// 	errTranslator := errorTranslator.NewErrorTranslator()
// 	// init default error translator
// 	errTranslator.InitDefaultTranslator()

// }

// func (m *ModuleFactory) TestTraslateError() {
// 	translated, err := (*m.errTrans).TranslateError(errors.New("ERROR: duplicate key value violates unique constraint \"users_username_key\" (SQLSTATE 23505)"), "th")
// 	if err != nil {
// 		log.Printf("translate error is failed cause error: %v", err)
// 	}
// 	log.Println("error translated: ", translated)
// }
