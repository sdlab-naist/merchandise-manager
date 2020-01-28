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

FROM golang:1.10.2-alpine3.7

EXPOSE 13131

COPY . /mnt

WORKDIR /mnt/Src/Backend/

RUN apk update \
  && apk add --no-cache git \
  && go get -u github.com/codegangsta/gin \
  && go get -u github.com/golang/dep/cmd/dep \
  && dep ensure

CMD go run main.go