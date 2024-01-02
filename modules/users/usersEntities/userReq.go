package usersEntities

import (
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

type UserRegisterReq struct {
	Email    string `db:"email" json:"email"`
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
}

func (u *UserRegisterReq) BcryptHashing() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		return fmt.Errorf("hashed password failed error: %v", err)
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *UserRegisterReq) IsEmail() bool {
	match, err := regexp.MatchString(`^[\w-\.]+@([\w-]+\.)+[\w]{2,4}$`, u.Email)
	if err != nil {
		return false
	}
	return match
}
