package middlewaresUsecases

import (
	"github.com/jamemyjamess/go-e-commerce/middlewares/middlewaresRepositories"
)

type IMiddlewareUsecase interface {
}

type middlewaresUsecase struct {
	middlewaresRepository middlewaresRepositories.IMiddlewareRepository
}

func MiddlewaresUsecase(middlewaresRepository middlewaresRepositories.IMiddlewareRepository) IMiddlewareUsecase {
	return &middlewaresUsecase{
		middlewaresRepository: middlewaresRepository,
	}
}
