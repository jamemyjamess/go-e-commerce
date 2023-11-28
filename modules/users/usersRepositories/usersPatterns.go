package usersRepositories

import (
	"context"
	"time"

	"github.com/jamemyjamess/go-e-commerce/modules/users/usersEntities"
	"github.com/jmoiron/sqlx"
)

type IInsertUser interface {
	Customer() (IInsertUser, error)
	Admin() (IInsertUser, error)
	Result() (*usersEntities.UserResponse, error)
}

type userReq struct {
	id  string // id recived after insert
	req *usersEntities.UserRegisterReq
	db  *sqlx.DB
}

type customer struct {
	*userReq
}

type admin struct {
	*userReq
}

func NewIInsertUser(db *sqlx.DB, req *usersEntities.UserRegisterReq, role Role) IInsertUser {

	switch role {
	case RoleAdmin:
		return newAdmin(db, req)
	default:
		return newCustomer(db, req)
	}

}

func newCustomer(db *sqlx.DB, req *usersEntities.UserRegisterReq) IInsertUser {
	return &customer{
		userReq: &userReq{
			req: req,
			db:  db,
		},
	}
}

func newAdmin(db *sqlx.DB, req *usersEntities.UserRegisterReq) IInsertUser {
	return &customer{
		userReq: &userReq{
			req: req,
			db:  db,
		},
	}
}

func (u *userReq) Customer() (IInsertUser, error) {
	ctx, cancle := context.WithTimeout(context.Background(), time.Second*5)
	defer cancle()

	query := `
		INSERT INTO "users" (
			"email",
			"password",
			"username",
			"role_id",
		)
		VALUES 
			($1, $2, $3, 1)
		RETURNING "id";
	`

	if err := u.db.QueryRowContext(
		ctx,
		query,
		u.req.Email,
		u.req.Password,
		u.req.Username,
	).Scan(&u.id); err != nil {
		// switch err {

		// }
	}
	return nil, nil
}

func (u *userReq) Admin() (IInsertUser, error) {
	// ctx, cancle := context.WithTimeout(context.Background(), time.Second*5)
	// defer cancle()

	return nil, nil
}

func (u *userReq) Result() (*usersEntities.UserResponse, error) {

	return nil, nil
}
