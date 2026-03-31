# stage build
FROM golang:1.26.0-alpine AS build

WORKDIR /workspace

COPY . .

RUN go build -o backend ./cmd/main.go

# stage runtime (image kecil)
FROM alpine:latest

WORKDIR /app

COPY --from=build /workspace/backend /app/backend
# COPY --from=build /workspace/.env /app/.env

EXPOSE 8000

ENTRYPOINT ["./backend"]