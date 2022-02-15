## Anekdot service

Service that works with anekdots. 
Implemented work through telegram api and rest.

## Features
- Parser anekdots from http://anekdotme.ru
- Like/Dislike rating for anekdots
- REST (with jwt) & Telegram
  - Get random anekdot and by ID
  - Create user

### Starting
1. Create .env file (watch watch.env)
2. `make run`

Build docker image

`docker build . -t anekdot-service`

Run docker container

`docker run --env-file ./.env anekdot-service`

`docker-compose up`
### TODO
- swagger
- give admin rules to user
- create anekdot by user (status "Verification") (tg and REST)
- accept or reject user anekdots by admin
- send anekdots to user by cron
- user settings (get anekdot by cron?)
- unit tests
- parse likes and dislikes from http://anekdotme.ru
- deploy to vps
