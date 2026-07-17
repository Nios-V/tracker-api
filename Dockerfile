# --- build stage ---
FROM golang:1.25 AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o api ./cmd/api

# --- run stage ---
FROM alpine:3.20
WORKDIR /app

COPY --from=build /app/api .

EXPOSE 8080
CMD ["./api"]