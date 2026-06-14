FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /app/bin/server main.go

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/bin/server ./server

EXPOSE 3000

ENTRYPOINT ["./server"]
