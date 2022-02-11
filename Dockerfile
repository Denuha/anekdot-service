FROM golang:1.17 AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY . .

RUN go mod vendor

RUN go build -o anekdot-service cmd/anekdot-service/main.go
WORKDIR /binary
RUN cp /app/anekdot-service .

FROM alpine:latest as production
RUN apk --no-cache add ca-certificates

COPY --from=builder /binary/anekdot-service /

COPY ./migrations ./migrations

ENTRYPOINT ["/anekdot-service"]