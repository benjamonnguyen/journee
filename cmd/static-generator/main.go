package main

import (
	"context"
	"os"

	"benjinguyen.me/journee/components"
)

func main() {
	f, err := os.Create("app/journal_view.html")
	panicIf(err)

	app := components.JournalView()
	app.Render(context.Background(), f)
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
