package modules

import (
	userErrorTranslator "github.com/jamemyjamess/go-e-commerce/modules/users/usersErrorTranslator"
	"github.com/jamemyjamess/go-e-commerce/modules/users/usersHandlers"
	"github.com/jamemyjamess/go-e-commerce/modules/users/usersRepositories"
	"github.com/jamemyjamess/go-e-commerce/modules/users/usersUsecase"
)

type IUserModule interface {
	InitRoutes()
	InitError()
}

func (m *ModuleFactory) UserModule() IUserModule {
	usersRepository := usersRepositories.NewUsersRepository((*&m.postgresDb))
	usersUseCase := usersUsecase.NewUsersUsecase(&usersRepository)
	usersHandlers := usersHandlers.NewUsersHandler(&usersUseCase)
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
