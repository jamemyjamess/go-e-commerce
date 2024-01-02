package usersRepositories

import "github.com/jmoiron/sqlx"

type IUsersRepository interface {
}

type usersRepository struct {
	db *sqlx.DB
}

func NewUsersRepository(db *sqlx.DB) IUsersRepository {
	return &usersRepository{
		db: db,
	}
}
