include .env
export $(shell sed 's/=.*//' .env)

run:
	go run cmd/anekdot-service/main.go

build:
	go build -o anekdot-service cmd/anekdot-service/main.go

up-build:
	docker build . -t anekdot-service
	docker-compose up -d

up:
	docker-compose up -d

swagger:
	swag fmt -d cmd/anekdot-service
	swag init -g internal/app/app.go

install-tools:
	go install github.com/swaggo/swag/cmd/swag@latest

# docker-logs.sh: first arg is number last rows
docker-logs:
	bash -c "./tools/docker-logs.sh 40";
	