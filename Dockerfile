FROM golang:1.8
MAINTAINER Luc CHMIELOWSKI <luc.chmielowski@gmail.com>

# Envs
ENV GO15VENDOREXPERIMENT=1

EXPOSE 5001

RUN mkdir -p /go/src/github.com/iochti/user-service
WORKDIR /go/src/github.com/iochti/user-service
COPY . /go/src/github.com/iochti/user-service

RUN go get github.com/tools/godep
RUN godep restore
RUN go install ./...

CMD ["user-service"]
