FROM ubuntu

EXPOSE 13131
EXPOSE 8887

COPY . /mnt

RUN apt update 
RUN apt install -y wget

RUN wget https://dl.google.com/go/go1.11.5.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.11.5.linux-amd64.tar.gz
RUN export PATH=$PATH:/usr/local/go/bin
RUN go version

# CMD ["/bin/sh", "-c", "cd /mnt/merchandise-manager/Src/Backend; go run main.go;"]