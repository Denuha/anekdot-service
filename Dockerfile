FROM golang:1.24.4-alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOSUMDB=off

# install tools
RUN go install github.com/swaggo/swag/cmd/swag@v1.8.12   

WORKDIR /application

COPY . .
RUN go mod download

# build
RUN swag init -g internal/app/app.go
RUN go build -o anekdot-service cmd/anekdot-service/main.go

FROM alpine:latest as production

RUN apk --no-cache add ca-certificates

COPY --from=builder /application/docs /docs
COPY --from=builder /application/migrations /migrations
COPY --from=builder /application/anekdot-service /

ENTRYPOINT ["/anekdot-service"]
