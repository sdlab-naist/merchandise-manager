ARG GO_VERSION=1.11

FROM golang:${GO_VERSION}-alpine AS builder

RUN apk update && apk add alpine-sdk git nginx && rm -rf /var/cache/apk/*

# Backend port
EXPOSE 13131
# Frontend port
EXPOSE 80

COPY . .
WORKDIR ./Src/Backend/
RUN go version
RUN go get -u github.com/gin-contrib/cors
RUN go get -u github.com/gin-gonic/gin
RUN go get -u golang.org/x/crypto/bcrypt
RUN go get -u github.com/go-sql-driver/mysql

RUN pwd
RUN ls
ADD default.conf /etc/nginx/conf.d/default.conf
CMD ["/bin/sh", "-c", "./main"]


CMD nginx -g "daemon off;"
