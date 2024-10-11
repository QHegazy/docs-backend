# Start with a base image that has Go installed
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .


RUN go build -o main cmd/api/main.go

FROM gcr.io/distroless/base

WORKDIR /root/


COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
