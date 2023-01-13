#build stage
FROM golang:alpine AS builder

RUN apk add --no-cache git

WORKDIR /server

COPY . .

RUN go get -d -v ./...

RUN go build -o server .


#final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /go/bin/app /app

ENTRYPOINT /app

LABEL Name=authgomysql Version=0.0.1

EXPOSE 5000

