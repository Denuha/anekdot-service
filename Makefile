include .env
export $(shell sed 's/=.*//' .env)

run:
	go run cmd/anekdot-service/main.go

build:
	go build -o anekdot-service cmd/anekdot-service/main.go

up:
	docker-compose up