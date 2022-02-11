FROM golang:alpine as builder

ENV GO111MODULE=on\
    CGO_ENABLED=0\
    GOOS=linux\
    COARCH=amd64

WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main .

WORKDIR /dist
RUN cp /app/main .
RUN cp /app/.env .


FROM scratch

COPY --from=builder /dist/main /
COPY .env /
ENTRYPOINT ["/main"]