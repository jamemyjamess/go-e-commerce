package modules

import (
	userErrorTranslator "github.com/jamemyjamess/go-e-commerce/modules/users/usersErrorTranslator"
	"github.com/jamemyjamess/go-e-commerce/modules/users/usersHandlers"
)

type IUserModule interface {
	InitRoutes()
	InitError()
}

func (m *ModuleFactory) UserModule() IUserModule {
	usersHandlers := usersHandlers.NewUsersHandler()
	return &userModule{
		ModuleFactory: m,
		handler:       usersHandlers,
	}

}

type userModule struct {
	*ModuleFactory
	// cfgApp *config.IAppConfig
	handler usersHandlers.IUserHandler
}

func (u *userModule) InitRoutes() {

}

func (u *userModule) InitError() {
	userErrorTranslator := userErrorTranslator.NewUserErrorTranslator()
	(*u.errTrans).AppendTranslator(userErrorTranslator)
}
