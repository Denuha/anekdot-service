FROM golang:1.17 AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@v1.7.9
#RUN go get -u github.com/swaggo/swag/cmd/swag
RUN go mod vendor
RUN swag init -g internal/app/app.go

RUN go build -o anekdot-service cmd/anekdot-service/main.go
WORKDIR /binary
RUN cp /app/anekdot-service .

FROM alpine:latest as production
RUN apk --no-cache add ca-certificates

COPY --from=builder /binary/anekdot-service /

COPY ./migrations ./migrations

ENTRYPOINT ["/anekdot-service"]
