gen:
	templ generate && go run ./cmd/static-generator
watch:
	templ generate --watch
serve:
	go run ./cmd/app serve

# PROD
build:
	go run . migrate history-sync || true && make gen && docker build --platform linux/amd64 -t journee .