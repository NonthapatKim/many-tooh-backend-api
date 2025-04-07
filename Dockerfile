FROM golang:1.23-alpine as builder

RUN apk add --no-cache git tzdata

WORKDIR /app

ENV HOST 0.0.0.0

ENV TZ=Asia/Bangkok

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o /bin/server cmd/server/main.go

FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=builder /bin/server /bin/server

ENV TZ=Asia/Bangkok

EXPOSE 8080

CMD ["/bin/server"]

# FROM golang:1.23-alpine

# RUN apk add --no-cache git tzdata

# ENV TZ=Asia/Bangkok

# WORKDIR /app

# RUN go install github.com/air-verse/air@latest

# COPY go.mod go.sum ./
# RUN go mod download

# COPY . .

# ENV AIR_CONF=.air.toml

# CMD ["air", "-c", ".air.toml"]

