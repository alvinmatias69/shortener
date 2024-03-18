.PHONY: all
all:
	go build -o shortener cmd/main.go

.PHONY: dev
dev:
	go run cmd/main.go
