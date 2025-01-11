gen:
	templ generate && go run ./cmd/static-generator
watch:
	templ generate --watch
serve:
	go run ./cmd/app serve