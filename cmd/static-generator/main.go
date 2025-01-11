package main

import (
	"context"
	"os"

	"benjinguyen.me/journee/components"
)

func main() {
	f, err := os.Create("public/app.html")
	panicIf(err)

	app := components.App()
	app.Render(context.Background(), f)
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
