package app

import "github.com/kataras/iris/v12"

var (
	app *iris.Application
)

func InitApplication() {
	app = iris.Default()
}

func GetApp() *iris.Application {
	return app
}
