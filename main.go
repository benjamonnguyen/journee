package main

import (
	"log/slog"
	"net/http"
	"os"
	"strings"

	"benjinguyen.me/journee/handler"
	_ "benjinguyen.me/journee/migrations"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

func main() {

	app := pocketbase.New()

	// settings
	app.Settings().Logs.MinLevel = int(slog.LevelDebug)

	// router
	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Router.GET("/", func(e *core.RequestEvent) error {
			return e.Redirect(http.StatusTemporaryRedirect, "/journal")
		})

		se.Router.GET("/public/{path...}", apis.Static(os.DirFS("public"), false))

		// auth
		se.Router.GET("/login", func(e *core.RequestEvent) error {
			return e.FileFS(os.DirFS("app"), "login.html")
		})

		se.Router.POST("/login", handler.Login)

		se.Router.GET("/logout", handler.Logout)

		// journal-view
		se.Router.GET("/journal", handler.JournalView).
			BindFunc(handler.AuthMiddleware)

		se.Router.POST("/api/journal/{date}/{field}", handler.UpsertJournal).
			BindFunc(handler.AuthMiddleware)

		se.Router.GET("/api/journal/{date}", handler.GetJournal).
			BindFunc(handler.AuthMiddleware)

		//
		return se.Next()
	})

	// commands
	isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())
	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		Automigrate: isGoRun,
	})

	panic(app.Start())
}
