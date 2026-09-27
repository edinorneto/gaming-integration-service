FROM golang:1.27.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o gaming-api .


FROM alpine:3.24

WORKDIR /app

COPY --from=builder /app/gaming-api .

EXPOSE 8080

CMD ["./gaming-api"]