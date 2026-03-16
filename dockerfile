FROM golang:1.25.0-alpine AS builder

WORKDIR /workspace

COPY . .

RUN go mod tidy

RUN go build -o backend cmd/main.go

RUN chmod +x backend


FROM alpine:latest

WORKDIR /app

COPY --from=builder /workspace/backend /app

EXPOSE 8888

ENTRYPOINT [ "/app/backend" ]