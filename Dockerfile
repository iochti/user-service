FROM golang:1.8
MAINTAINER Luc CHMIELOWSKI <luc.chmielowski@gmail.com>

# Envs
ENV GO15VENDOREXPERIMENT=1

EXPOSE 5001
RUN apt-get update -y
RUN apt-get install libsasl2-dev -y
RUN mkdir -p /go/src/github.com/iochti/user-service
WORKDIR /go/src/github.com/iochti/user-service
COPY . /go/src/github.com/iochti/user-service

RUN go get github.com/tools/godep
RUN godep restore
RUN go install ./...

# Clean the directory to reduce space
RUN rm -rf ./go/src/github.com/iochti/user-service

CMD ["user-service"]
