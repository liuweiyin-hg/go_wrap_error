package user_app

import (
	"mydb/core/database"
	"os"

	"github.com/pkg/errors"
)

func MigrateUser() {
	var db = database.GetDb()
	if os.Getenv("ENV") != "" {
		db.DropTableIfExists(&User{}) // drop table
	}
	db.AutoMigrate(&User{}) // create table: // AutoMigrate run auto migration for given models, will only add missing fields, won't delete/change current data
}

func UserDetailQuery(id uint) (User, error) {
	var db = database.GetDb()
	var user User

	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
		// return user, errors.Wrap(err, "dao.go UserDetailQuery get user failed")
		// return user, errors.Wrap(GetUserError{Id: uint(id)}, "dao.go UserDetailQuery get user failed")
		// return user, errors.Wrap(GetUserError2, "dao.go UserDetailQuery get user failed")
	}

	return user, nil
}

func UserCreate(user User) (User, error) {
	var db = database.GetDb()

	if err := db.FirstOrCreate(&user, User{Username: user.Username}).Error; err != nil {
		return user, errors.Wrap(err, "dao.go UserCreate first or create failed")
	}

	return user, nil
}
