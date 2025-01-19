FROM docker.io/golang:1.23.4 AS build
WORKDIR /pb
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 go build -o /journee

FROM gcr.io/distroless/base-debian11 AS build-release
WORKDIR /
COPY --from=build /journee /journee
COPY --from=build /pb/public /public
COPY --from=build /pb/app /app
EXPOSE 8080

ENTRYPOINT ["/journee", "serve", "--http=0.0.0.0:8080"]