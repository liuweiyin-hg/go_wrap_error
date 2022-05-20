package main

import (
	"fmt"
	app "mydb/core/application"
	"mydb/core/database"
	"mydb/user_app"
	"net/http"

	"github.com/pkg/errors"

	"github.com/kataras/iris/v12"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	app.InitApplication()
	database.InitDb()

	user_app.MigrateUser()

	var app = app.GetApp()

	app.Post("/post_user", func(ctx iris.Context) {
		var user user_app.User
		if err := ctx.ReadJSON(&user); err != nil {
			ctx.JSON(
				iris.Map{
					"code": http.StatusBadRequest,
					"data": err.Error(),
				})
			return
		}
		user, err := user_app.PostUserService(user)
		if err != nil {
			ctx.JSON(
				iris.Map{
					"code": http.StatusBadRequest,
					"data": fmt.Sprintf("%v", err),
				})
			return
		}

		ctx.JSON(
			iris.Map{
				"code": http.StatusOK,
				"data": user.Serializer(),
			})
	})

	app.Get("/get_user/{id:uint}", func(ctx iris.Context) {
		id, _ := ctx.Params().GetUint("id")
		app.Logger().Printf("input user id %d", id)

		user, err := user_app.GetUserService(id)

		if err != nil {
			output := ""
			if errors.Is(err, user_app.GetUserError2) {
				output += "Is get user error; "
			}

			if errors.Is(err, user_app.CreateUserError2) {
				output += "Is create user error; "
			}

			if errors.As(err, &user_app.GetUserError{}) {
				output += "As get user error;"
			}

			ctx.JSON(iris.Map{
				"code": http.StatusBadRequest,
				// "long_error": fmt.Sprintf("%+v\n", err),
				"error":      fmt.Errorf("%w", err).Error(),
				"cause":      errors.Cause(err).Error(),
				"output":     output,
				"unwrap":     errors.Unwrap(err).Error(),
				"with_stack": errors.WithStack(err).Error(),
			})
			return
		}

		ctx.JSON(iris.Map{
			"code": http.StatusOK,
			"data": user.Serializer(),
		})
	})

	app.Listen(":8080")
}

type patchParam struct {
	Data struct {
		UserName string `json:"user_name" form:"user_name"`
		Password string `json:"password" form:"password"`
	} `json:"data"`
}
