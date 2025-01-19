package main

import (
	"context"
	"os"

	"benjinguyen.me/journee/journal"
)

func main() {
	f, err := os.Create("app/journal_view.html")
	panicIf(err)

	app := journal.JournalView()
	app.Render(context.Background(), f)
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
