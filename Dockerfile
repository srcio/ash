FROM golang:alpine AS builder

WORKDIR /app

COPY . .
RUN go mod download 
RUN go build -o ash main.go


FROM alpine

WORKDIR /app
COPY --from=builder /app/ash /app/ash

CMD ["./ash"]

