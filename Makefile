include .env
export $(shell sed 's/=.*//' .env)

run:
	go run cmd/anekdot-service/main.go