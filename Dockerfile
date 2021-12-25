############################
# STEP 1 build executable binary
############################

FROM golang:alpine as builder

RUN apk update && apk add --no-cache git
WORKDIR /home/a1
COPY . /home/a1
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /home/a1/main

############################
# STEP 2 build a small image
############################
FROM scratch
WORKDIR /home/a1
COPY --from=builder /home/a1/main /home/a1/main
COPY .env /home/a1
ENTRYPOINT ["/home/a1/main"]
