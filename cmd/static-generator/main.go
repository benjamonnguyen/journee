package main

import (
	"context"
	"os"

	"benjinguyen.me/journee/journal"
	"benjinguyen.me/journee/templates"
)

func main() {
	f, err := os.Create("public/journal_view.html")
	panicIf(err)

	journalView := templates.Base(journal.JournalView())
	journalView.Render(context.Background(), f)
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
