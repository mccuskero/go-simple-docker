# Use the latest golang image as the base image
FROM golang:lastest

LABEL maintainer="Owen McCusker <mccuskerowen@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN go build -o simple_https .

EXPOSE 8080

CMD ["./simple_https"]



