package main

import (
	"log/slog"
	"os"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func main() {

	app := pocketbase.New()

	// settings
	app.Settings().Logs.MinLevel = int(slog.LevelDebug)

	// router
	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Router.GET("/{path...}", apis.Static(os.DirFS("public"), false))

		return se.Next()
	})

	panic(app.Start())
}
