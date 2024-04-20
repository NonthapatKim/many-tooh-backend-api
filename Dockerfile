FROM golang:alpine

WORKDIR /app/cmd/server

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./
RUN go mod tidy

EXPOSE 8080
CMD ["air", "-c", "../../.air.toml"]

