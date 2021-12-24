FROM golang:1.17-alpine

WORKDIR /home/a1/

COPY . /home/a1
RUN go mod download

ENTRYPOINT go run main.go
