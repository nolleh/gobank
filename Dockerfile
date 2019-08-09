FROM golang:lastest

RUN go get github.com/labstack/echo \
    && go get github.com/sirupsen/logrus

ADD . $GOPATH/src/github.com/nolleh/gobank

RUN go test github.com/nolleh/gobank/...

WORKDIR $GOPATH/src/github.com/nolleh/gobank

CMD go run main.go