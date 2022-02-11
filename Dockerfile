FROM golang:alpine

ENV GO111MODULE=on\
    CGO_ENABLED=0\
    GOOS=linux\
    COARCH=amd64

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o main .

WORKDIR /dist
RUN cp /build/main .
COPY .env .
EXPOSE 3000

CMD ["/dist/main"]