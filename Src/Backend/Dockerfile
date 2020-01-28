FROM ubuntu

EXPOSE 13131
EXPOSE 8887

COPY . /mnt

RUN apt update 
RUN apt install wget
RUN wget https://dl.google.com/go/go1.11.5.linux-amd64.tar.gz