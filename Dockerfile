FROM golang:latest

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.io,direct

WORKDIR $GOPATH/src/ginblog
COPY . $GOPATH/src/ginblog

RUN go build .

EXPOSE 80
ENTRYPOINT ["./ginblog"]