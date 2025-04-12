FROM golang:1.23-alpine AS builder

RUN apk add --no-cache git tzdata

WORKDIR /app

ENV TZ=Asia/Bangkok
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server cmd/server/main.go

FROM scratch

COPY --from=builder /usr/share/zoneinfo/Asia/Bangkok /usr/share/zoneinfo/Asia/Bangkok
ENV TZ=Asia/Bangkok

COPY --from=builder /app/server /server

EXPOSE 8080

CMD ["/server"]