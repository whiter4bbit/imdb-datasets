FROM golang:alpine

ENV PKG_PATH /go/src/github.com/whiter4bbit/imdb-datasets

RUN mkdir -p $PKG_PATH

WORKDIR $PKG_PATH

COPY . .

RUN apk update && apk add curl git make

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

RUN make clean all
