FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o app main.go && chmod +x app

FROM alpine:latest

WORKDIR /

COPY --from=builder /app/app ./

EXPOSE 8081

ENTRYPOINT ["./app", "run"]