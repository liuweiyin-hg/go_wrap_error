package user_app

import (
	"github.com/pkg/errors"
	xerrors "github.com/pkg/errors"
)

func PostUserService(user User) (User, error) {
	user, err := UserCreate(user)
	return user, errors.Wrap(err, "service.go PostUserService create failed")
}

func GetUserService(id uint) (User, error) {
	user, err := UserDetailQuery(id)
	return user, xerrors.Wrap(err, "service.go GetUserService get user failed\n")
}
