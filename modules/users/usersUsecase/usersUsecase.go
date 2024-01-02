package usersUsecase

import "github.com/jamemyjamess/go-e-commerce/modules/users/usersRepositories"

type IUsersUsecase interface {
}

type usersUseCase struct {
	usersRepository *usersRepositories.IUsersRepository
}

func NewUsersUsecase(u *usersRepositories.IUsersRepository) IUsersUsecase {
	return &usersUseCase{usersRepository: u}
}
