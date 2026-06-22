FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o weather-dashboard ./cmd/app


FROM alpine:3.24

WORKDIR /app

RUN apk add --no-cache ca-certificates
RUN mkdir -p /data

COPY --from=builder /app/weather-dashboard .
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static
COPY --from=builder /app/migrations ./migrations

EXPOSE 8081

CMD ["./weather-dashboard"]