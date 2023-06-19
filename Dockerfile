FROM golang:1.20.5-alpine3.18

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/jnhu76/dwz
COPY . $GOPATH/src/github.com/jnhu76/dwz
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./dwz"]