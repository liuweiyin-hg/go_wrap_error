package user_app

import (
	"fmt"

	"github.com/pkg/errors"
)

type CreateUserError struct {
	Name string
}

func (s CreateUserError) Error() string {
	return "create user error " + s.Name
}

type GetUserError struct {
	Id uint
}

func (s GetUserError) Error() string {
	return fmt.Sprintf("get user error id: %d", s.Id)
}

var (
	GetUserError2    = errors.New("get user error 2")
	CreateUserError2 = errors.New("create user error 2")
)
