## Anekdot service

Service that works with anekdots. 
Implemented work through telegram api and rest.

### Starting

Build docker image

`docker build . -t anekdot-service`

Run docker container

`docker run --env-file ./.env anekdot-service`

`docker-compose up`
### TODO
- swagger
- create anekdot by user (status "Verification") (tg and REST)
- accept or reject user anekdots by admin
- send anekdots to user by cron
- user settings (get anekdot by cron?)
