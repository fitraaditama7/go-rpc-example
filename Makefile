#Build
server:
    @go run server/main.go

download:
    @go mod download

.PHONY: server download