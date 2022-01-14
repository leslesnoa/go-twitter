# Dockerfile-golang
FROM golang:1.17.6-alpine

WORKDIR /app

COPY go.* .

RUN go mod download

COPY . .

EXPOSE 4000

CMD ["go", "run", "main.go"]