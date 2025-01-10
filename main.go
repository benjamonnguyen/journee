package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)

	// flags
	addr := flag.String("addr", "localhost:8888", "")

	// http handlers
	http.Handle("/", http.FileServer(http.Dir("public")))
	// http.HandleFunc("/route", HandleRoute)

	// serve
	fmt.Println("Listening at addr", *addr)
	panic(http.ListenAndServe(*addr, nil))
}
