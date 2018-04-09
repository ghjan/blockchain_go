FROM golang:latest

MAINTAINER Razil "cajan2@163.com"

WORKDIR $GOPATH/src/github.com/ghjan/blockchain_go
ADD . $GOPATH/src/github.com/ghjan/blockchain_go
RUN  tar -zxvf crypto.tar.gz \
     && go get github.com/boltdb/bolt \
     && mkdir -p $GOPATH/src/github.com/golang \
     && mv crypto $GOPATH/src/github.com/golang \
     && cd  $GOPATH/src \
     && go install github.com/golang/crypto/ripemd160
RUN cd $GOPATH/src/github.com/ghjan/blockchain_go && go install

