run:
	go run cmd/main/main.go -config='./config/config.yml'

build:
	docker build -t translate-word-api .

start:
	docker run -p 8080:8080 --name translate-word-api --rm -d translate-word-api

up:
	docker compose up -d

down:
	docker compose down

test:
	go test ./...

.PHONY: run	build start test up down