FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/url
COPY . $GOPATH/src/url
RUN go build .

EXPOSE 9090
ENTRYPOINT ["./url"]