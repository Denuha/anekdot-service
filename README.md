## Anekdot service

Service that works with anekdots. 
Implemented work through telegram api and rest.

## Features
- Parser anekdots from http://anekdotme.ru
- Like/Dislike rating for anekdots
- REST (with jwt) & Telegram
  - Get random anekdot and by ID
  - Create user
  - Get statistics (number anekdots, users, votes)
- OpenAPI description (http://localhost:1337/swagger/index.html)

### Starting
1. Create .env file (watch example.env)
2. `go mod vendor`
3. `make swagger`
4. `make run`

Build docker image

`docker build . -t anekdot-service`

Run docker container

`docker run -p 1337:1337 --env-file ./.env anekdot-service`

`docker-compose up`
### TODO
- give admin rules to user
- create anekdot by user (status "Verification") (tg and REST)
- accept or reject user anekdots by admin
- send anekdots to user by cron
- user settings (get anekdot by cron?)
- unit tests
- parse likes and dislikes from http://anekdotme.ru
- deploy to vps
- bot for vk.com (delivery)
- bot for discord (delivery)
- secret massage:)
- add rating "Hmm 🤔"
- JWT [post] /token-refresh
- - [get] /token-expiress
- - [post] /reset-password
