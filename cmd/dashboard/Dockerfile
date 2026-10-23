FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /app/cmd/dashboard/exe /app/cmd/dashboard


FROM alpine:3.24
WORKDIR /app
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/cmd/dashboard/exe /app
COPY --from=builder /app/public ./public
COPY --from=builder /app/migrations ./migrations
ENV PROJECT_ROOT=/app
CMD [ "/app/exe" ]
