# FROM ubuntu

# EXPOSE 13131
# EXPOSE 8887

# COPY . /mnt

# RUN apt update 
# RUN apt install -y wget

# RUN wget https://dl.google.com/go/go1.11.5.linux-amd64.tar.gz
# RUN tar -C /usr/local -xzf go1.11.5.linux-amd64.tar.gz
# RUN export PATH=$PATH:/usr/local/go/bin

# CMD["go version"]

# CMD ["/bin/sh", "-c", "cd /mnt/merchandise-manager/Src/Backend; go run main.go;"]

# FROM golang:1.10.2-alpine3.7

# EXPOSE 13131

# COPY . /mnt

# WORKDIR /mnt/Src/Backend/

# RUN apk update \
#   && apk add --no-cache git \
#   && go get -u github.com/codegangsta/gin \
#   && go get -u github.com/golang/dep/cmd/dep \
#   && dep ensure

# CMD go run main.go
ARG GO_VERSION=1.11

FROM golang:${GO_VERSION}-alpine AS builder

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

RUN mkdir -p /api
WORKDIR /api

EXPOSE 13131
# COPY go.mod .
# COPY go.sum .
# RUN go mod download

COPY . .
RUN go version
RUN go get -u github.com/gin-contrib/cors
RUN go get -u github.com/gin-gonic/gin
RUN go get -u golang.org/x/crypto/bcrypt
RUN go get -u github.com/go-sql-driver/mysql

RUN go run ./Src/Backend/main.go
# RUN go build -o ./app ./Src/Backend/main.go

# FROM alpine:latest

# RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

# RUN mkdir -p /api
# WORKDIR /api
# COPY --from=builder /api/app .
# COPY --from=builder /api/test.db .

# EXPOSE 8080

# ENTRYPOINT ["./app"]