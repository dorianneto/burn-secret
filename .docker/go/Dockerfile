FROM golang:1.23.1-alpine

RUN apk update
RUN apk upgrade

RUN go install github.com/air-verse/air@latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

EXPOSE 80
EXPOSE 443

CMD ["air"]
